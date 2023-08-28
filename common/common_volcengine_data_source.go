package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type ExtraData func([]interface{}) ([]interface{}, error)

type EachResource func([]interface{}, *schema.ResourceData) ([]interface{}, error)

type DataSourceInfo struct {
	RequestConverts  map[string]RequestConvert
	ResponseConverts map[string]ResponseConvert
	NameField        string
	IdField          string
	CollectField     string
	ContentType      RequestContentType
	ExtraData        ExtraData
	ServiceCategory  ServiceCategory
	EachResource     EachResource
}

func DataSourceToRequest(d *schema.ResourceData, r *schema.Resource, info DataSourceInfo) (req map[string]interface{}, err error) {
	if info.RequestConverts == nil {
		info.RequestConverts = make(map[string]RequestConvert)
	}
	if _, ok := info.RequestConverts["name_regex"]; !ok {
		info.RequestConverts["name_regex"] = RequestConvert{
			Ignore: true,
		}
	}
	if _, ok := info.RequestConverts["output_file"]; !ok {
		info.RequestConverts["output_file"] = RequestConvert{
			Ignore: true,
		}
	}
	return ResourceDateToRequest(d, r, false, info.RequestConverts, RequestConvertAll, info.ContentType)
}

func ResponseToDataSource(d *schema.ResourceData, r *schema.Resource, info DataSourceInfo, collection []interface{}) (err error) {
	var (
		result []map[string]interface{}
	)

	for _, item := range collection {
		var (
			temp map[string]interface{}
			flag bool
		)
		temp, flag, err = mergeNameRegex(d, item.(map[string]interface{}), info.NameField)
		if err != nil {
			return err
		}
		if flag {
			if temp != nil {
				result = append(result, temp)
			}
		} else {
			result = append(result, item.(map[string]interface{}))
		}
	}
	_, _, err = datasourceMapping(d, result, dataSource{
		idField: info.IdField,
		idValue: func(idField string, item map[string]interface{}) string {
			return item[idField].(string)
		},
		sliceValue: func(item map[string]interface{}) map[string]interface{} {
			return mergeDatasource(r, info.CollectField, item, info.ResponseConverts)
		},
		targetName: info.CollectField,
	})
	return err
}

type sliceValueFunc func(map[string]interface{}) map[string]interface{}

type idValueFunc func(string, map[string]interface{}) string

type dataSource struct {
	idField    string
	idValue    idValueFunc
	sliceValue sliceValueFunc
	targetName string
}

func mapMapping(result interface{}, ds dataSource) (map[string]interface{}, error) {
	var data map[string]interface{}
	if reflect.TypeOf(result).Kind() == reflect.Map {
		if v, ok := result.(map[string]interface{}); ok {
			if ds.sliceValue != nil {
				data = ds.sliceValue(v)
			}
		}
	}
	return data, nil
}

func mergeDatasource(resource *schema.Resource, collectField string, item map[string]interface{}, extraMapping map[string]ResponseConvert) map[string]interface{} {
	result := make(map[string]interface{})
	keys := strings.Split(collectField, ".")
	extra := make(map[string]interface{})
	if len(keys) == 0 {
		return result
	}

	if _, ok := resource.Schema[keys[0]]; ok {
		elem := getSchemeElem(resource, keys)
		for k, v := range item {
			target := HumpToDownLine(k)
			m := ResponseConvert{}
			if extraMapping != nil {
				if _, ok := extraMapping[k]; ok {
					m = extraMapping[k]
					if m.Ignore {
						continue
					}
					if m.Chain == "" {
						//if no chain to set auto check convert in elem.schema.if in and set new target
						if _, ok1 := elem.Schema[m.TargetField]; ok1 {
							target = m.TargetField
						} else {
							m = ResponseConvert{}
						}
					} else if strings.HasSuffix(collectField, m.Chain) {
						// if set a chain,auto check is in collectField chain
						target = m.TargetField
					} else {
						//do nothing
						m = ResponseConvert{}
					}
				}
			}
			if targetValue, ok := elem.Schema[target]; ok {
				// response value change
				if m.Convert != nil {
					v = m.Convert(v)
				}

				if targetValue.Type == schema.TypeList || targetValue.Type == schema.TypeSet {
					if _, ok := targetValue.Elem.(*schema.Schema); ok {
						extra[target] = v
					} else {
						if _, ok := extra[target]; !ok {
							if l, ok := v.([]interface{}); ok {
								_, result, _ := datasourceMapping(nil, l, dataSource{
									sliceValue: func(m1 map[string]interface{}) map[string]interface{} {
										return mergeDatasource(resource, collectField+"."+target, m1, extraMapping)
									},
								})
								extra[target] = result
							} else if m, ok := v.(map[string]interface{}); ok {
								result, _ := mapMapping(m, dataSource{
									sliceValue: func(m1 map[string]interface{}) map[string]interface{} {
										return mergeDatasource(resource, collectField+"."+target, m1, extraMapping)
									},
								})
								extra[target] = []map[string]interface{}{
									result,
								}

							}
						}
					}
				}
				if _, ok := extra[target]; !ok {
					if m.TargetField == "" {
						result[HumpToDownLine(k)] = v
					} else {
						result[m.TargetField] = v
						if m.KeepDefault {
							result[HumpToDownLine(k)] = result[m.TargetField]
						}
					}
				} else {
					result[target] = extra[target]
				}
			} else {
				continue
			}
		}
	}

	for k, convert := range extraMapping {
		if strings.Contains(k, ".") {
			v, _ := ObtainSdkValue(k, item)

			if v == nil {
				continue
			}

			if convert.TargetField == "" {
				continue
			}

			if convert.Convert == nil {
				result[convert.TargetField] = v
			} else {
				result[convert.TargetField] = convert.Convert(v)
			}
		}
	}
	return result
}

func datasourceMapping(d *schema.ResourceData, result interface{}, datasource dataSource) ([]string, []map[string]interface{}, error) {
	var err error
	var ids []string
	ids = []string{}
	var data []map[string]interface{}
	data = []map[string]interface{}{}

	if reflect.TypeOf(result).Kind() == reflect.Slice {
		var length int
		if v, ok := result.([]map[string]interface{}); ok {
			length = len(v)
			for _, v1 := range v {
				ids, data = datasourceSliceMapping(ids, data, datasource, v1)
			}
		} else {
			root := result.([]interface{})
			length = len(root)
			for _, v2 := range root {
				ids, data = datasourceSliceMapping(ids, data, datasource, v2)
			}
		}

		if d != nil && datasource.targetName != "" {
			d.SetId(hashStringArray(ids))
			_ = d.Set("total_count", length)
			err = d.Set(datasource.targetName, data)
			if err != nil {
				return nil, nil, err
			}
			if outputFile, ok := d.GetOk("output_file"); ok && outputFile.(string) != "" {
				err = writeToFile(outputFile.(string), data)
				if err != nil {
					return nil, nil, err
				}
			}
		}

	}
	return ids, data, nil
}

func datasourceSliceMapping(ids []string, data []map[string]interface{}, datasource dataSource, item interface{}) ([]string, []map[string]interface{}) {
	if mm, ok := item.(map[string]interface{}); ok {
		if datasource.idValue != nil && datasource.idField != "" {
			ids = append(ids, datasource.idValue(datasource.idField, mm))
		}
		if datasource.sliceValue != nil {
			data = append(data, datasource.sliceValue(mm))
		}
	}
	return ids, data
}

func writeToFile(filePath string, data interface{}) error {
	absPath, err := absolutePath(filePath)
	if err != nil {
		return err
	}
	_ = os.Remove(absPath)
	var bs []byte
	switch data := data.(type) {
	case string:
		bs = []byte(data)
	default:
		bs, err = json.MarshalIndent(data, "", "\t")
		if err != nil {
			return fmt.Errorf("MarshalIndent data %#v and got an error: %#v", data, err)
		}
	}

	return ioutil.WriteFile(absPath, bs, 0422)
}

func absolutePath(filePath string) (string, error) {
	if strings.HasPrefix(filePath, "~") {
		usr, err := user.Current()
		if err != nil {
			return "", fmt.Errorf("get current user got an error: %#v", err)
		}

		if usr.HomeDir != "" {
			filePath = strings.Replace(filePath, "~", usr.HomeDir, 1)
		}
	}
	return filepath.Abs(filePath)
}

func hashStringArray(arr []string) string {
	var buf bytes.Buffer
	for _, s := range arr {
		buf.WriteString(fmt.Sprintf("%s-", s))
	}
	return fmt.Sprintf("%d", hashcode.String(buf.String()))
}

func mergeNameRegex(d *schema.ResourceData, data map[string]interface{}, nameField string) (result map[string]interface{}, flag bool, err error) {
	if nameRegex, ok := d.GetOk("name_regex"); ok {
		match := regexp.MustCompile(nameRegex.(string))
		if match.MatchString(data[nameField].(string)) {
			return data, true, err
		}
		return nil, true, err
	}
	return nil, false, err
}

func getSchemeElem(resource *schema.Resource, keys []string) *schema.Resource {
	r := resource
	if r == nil {
		return nil
	}
	for _, v := range keys {
		if elem, o := r.Schema[v].Elem.(*schema.Resource); o {
			r = elem
		}
	}
	return r
}
