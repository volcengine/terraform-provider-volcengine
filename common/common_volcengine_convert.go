package common

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type FieldResponseConvert func(interface{}) interface{}

type ResponseConvert struct {
	TargetField string
	KeepDefault bool
	Convert     FieldResponseConvert
	Ignore      bool
	Chain       string
}

type FieldRequestConvert func(*schema.ResourceData, interface{}) interface{}

type RequestConvert struct {
	ConvertType      RequestConvertType
	Convert          FieldRequestConvert
	Ignore           bool
	ForceGet         bool
	TargetField      string
	NextLevelConvert map[string]RequestConvert
	StartIndex       int
	SpecialParam     *SpecialParam
}
type SpecialParam struct {
	Type  SpecialParamType
	Index int
}

var supportRequestConvertType = map[RequestContentType]map[RequestConvertType]bool{
	ContentTypeDefault: {
		ConvertDefault:    true,
		ConvertWithN:      true,
		ConvertListN:      true,
		ConvertListUnique: true,
		ConvertSingleN:    true,
	},
	ContentTypeJson: {
		ConvertDefault:         true,
		ConvertJsonObject:      true,
		ConvertJsonArray:       true,
		ConvertJsonObjectArray: true,
	},
}

func checkRequestConvertTypeSupport(contentType RequestContentType, convertType RequestConvertType) bool {
	if v, ok := supportRequestConvertType[contentType]; ok {
		if _, ok1 := v[convertType]; ok1 {
			return true
		}
	}
	return false
}

func ResponseToResourceData(d *schema.ResourceData, resource *schema.Resource, data interface{}, extra map[string]ResponseConvert, start ...bool) (rd interface{}, err error) {
	setFlag := false
	if start == nil || (len(start) > 0 && start[0]) {
		setFlag = true
	}
	kind := reflect.ValueOf(data).Kind()
	switch kind {
	case reflect.Map:
		tempRd := make(map[string]interface{})
		dataMap := data.(map[string]interface{})
		for k, v := range dataMap {
			var (
				targetValue interface{}
				targetField string
				convert     ResponseConvert
			)
			// k -> targetField convert
			targetField = HumpToDownLine(k)
			if _, ok := extra[k]; ok {
				convert = extra[k]
				if convert.Ignore {
					continue
				}
				if convert.TargetField != "" {
					targetField = convert.TargetField
				}
			}
			if r, ok := resource.Schema[targetField]; ok {
				//v -> targetValue convert
				if r.Elem != nil {
					if elem, ok1 := r.Elem.(*schema.Resource); ok1 {
						if convert.Convert != nil {
							targetValue, err = ResponseToResourceData(d, elem, convert.Convert(v), extra, false)
						} else {
							targetValue, err = ResponseToResourceData(d, elem, v, extra, false)
						}
						if err != nil {
							return rd, err
						}
					} else if _, ok2 := r.Elem.(*schema.Schema); ok2 {
						if convert.Convert != nil {
							targetValue = convert.Convert(v)
						} else {
							targetValue = v
						}
					} else if _, ok3 := r.Elem.(schema.ValueType); ok3 {
						if convert.Convert != nil {
							targetValue = convert.Convert(v)
						} else {
							targetValue = v
						}
					}
				} else {
					if convert.Convert != nil {
						targetValue = convert.Convert(v)
					} else {
						targetValue = v
					}
				}
				// set targetValue to terraform local storage
				if setFlag {
					if (resource.Schema[targetField].Type == schema.TypeList ||
						resource.Schema[targetField].Type == schema.TypeSet) &&
						reflect.ValueOf(targetValue).Kind() == reflect.Map {
						err = d.Set(targetField, []interface{}{targetValue})
					} else {
						err = d.Set(targetField, targetValue)
					}
					if err != nil {
						return rd, err
					}
				} else {
					tempRd[targetField] = targetValue
				}
			} else {
				continue
			}
		}

		if setFlag {
			for k, convert := range extra {
				if strings.Contains(k, ".") {
					v, _ := ObtainSdkValue(k, data)
					v1 := v

					if v == nil || convert.TargetField == "" {
						continue
					}

					if convert.Convert != nil {
						v1 = convert.Convert(v)
					}

					tempRd[convert.TargetField] = v1

					if (resource.Schema[convert.TargetField].Type == schema.TypeList ||
						resource.Schema[convert.TargetField].Type == schema.TypeSet) &&
						reflect.ValueOf(convert.TargetField).Kind() == reflect.Map {
						_ = d.Set(convert.TargetField, []interface{}{v1})

						continue
					}

					_ = d.Set(convert.TargetField, v1)
				}
			}
		}

		if len(tempRd) > 0 {
			return tempRd, err
		}
	case reflect.Slice:
		var (
			tempRd []interface{}
		)
		root := data.([]interface{})
		for _, v := range root {
			var (
				targetValue interface{}
			)
			targetValue, err = ResponseToResourceData(d, resource, v, extra, false)
			tempRd = append(tempRd, targetValue)
		}
		if len(tempRd) > 0 {
			return tempRd, err
		}
	default:
		return rd, err
	}
	return rd, err
}

func ResourceDateToRequest(d *schema.ResourceData, resource *schema.Resource, isUpdate bool, convert map[string]RequestConvert, mode RequestConvertMode, contentType RequestContentType) (map[string]interface{}, error) {
	var req map[string]interface{}
	var err error
	req = make(map[string]interface{})
	count := 1
	var onlyMode bool

	switch mode {
	case RequestConvertAll:
		onlyMode = false
	case RequestConvertInConvert:
		onlyMode = true
	case RequestConvertIgnore:
		return req, err
	default:
		onlyMode = true
	}

	if convert != nil && onlyMode {
		for k, v := range convert {
			if isUpdate {
				count, err = RequestUpdateConvert(d, k, v, count, &req, contentType)
			} else {
				count, err = RequestCreateConvert(d, k, v, count, &req, v.ForceGet, contentType)
			}
		}
	} else {
		for k := range resource.Schema {
			if v, ok := convert[k]; ok {
				if isUpdate {
					count, err = RequestUpdateConvert(d, k, v, count, &req, contentType)
				} else {
					count, err = RequestCreateConvert(d, k, v, count, &req, v.ForceGet, contentType)
				}
			} else {
				if isUpdate {
					count, err = RequestUpdateConvert(d, k, RequestConvert{}, count, &req, contentType)
				} else {
					count, err = RequestCreateConvert(d, k, RequestConvert{}, count, &req, false, contentType)
				}
			}
		}
	}

	return req, err
}

func Convert(d *schema.ResourceData, k string, v interface{}, t RequestConvert, index int, req *map[string]interface{}, chain string, forceGet bool, contentType RequestContentType, schemaChain string, setIndex []int) (err error) {
	if !checkRequestConvertTypeSupport(contentType, t.ConvertType) {
		return fmt.Errorf("Can not support the RequestContentType [%v] when RequestContentType is [%v] ", t.ConvertType, contentType)
	}

	switch t.ConvertType {
	case ConvertDefault:
		err = RequestConvertDefault(v, k, t, req, chain)
	case ConvertSingleN:
		err = RequestConvertSingleN(v, k, t, req, chain)
	case ConvertWithN:
		err = RequestConvertWithN(v, k, t, req, chain)
	case ConvertListN:
		err = RequestConvertListN(v, k, t, req, chain, d, forceGet, false, contentType, schemaChain, setIndex)
	case ConvertListUnique:
		err = RequestConvertListN(v, k, t, req, chain, d, forceGet, true, contentType, schemaChain, setIndex)
	case ConvertJsonObject: //equal ConvertListUnique
		err = RequestConvertListN(v, k, t, req, chain, d, forceGet, true, contentType, schemaChain, setIndex)
	case ConvertJsonArray: //equal ConvertWithN
		err = RequestConvertWithN(v, k, t, req, chain)
	case ConvertJsonObjectArray: //equal ConvertListN
		err = RequestConvertListN(v, k, t, req, chain, d, forceGet, false, contentType, schemaChain, setIndex)
		//case ConvertWithFilter:
		//	index, err = RequestConvertWithFilter(v, k, t, index, req)
		//	break
		//case ConvertListFilter:
		//	index, err = RequestConvertListFilter(v, k, t, index, req)
		//	break
	}
	return err
}

func RequestCreateConvert(d *schema.ResourceData, k string, t RequestConvert, index int, req *map[string]interface{}, forceGet bool, contentType RequestContentType) (int, error) {
	var err error
	var ok bool
	var v interface{}
	if t.Ignore {
		return index, err
	}
	if t.Convert != nil {
		v = t.Convert(d, d.Get(k))
		ok = true
	} else {
		if forceGet {
			v = d.Get(k)
			ok = true
			if str, ok1 := v.(string); ok1 {
				if str == "" {
					ok = false
				}
			}

		} else {
			v, ok = d.GetOk(k)
		}
	}
	if ok {
		err = Convert(d, k, v, t, index, req, "", forceGet, contentType, "", nil)
	}
	return index, err
}

func RequestUpdateConvert(d *schema.ResourceData, k string, t RequestConvert, index int, req *map[string]interface{}, contentType RequestContentType) (int, error) {
	var err error

	if t.ForceGet || (d.HasChange(k) && !d.IsNewResource()) {
		index, err = RequestCreateConvert(d, k, t, index, req, true, contentType)
	}
	return index, err
}

func RequestConvertDefault(v interface{}, k string, t RequestConvert, req *map[string]interface{}, chain string) error {
	if strings.TrimSpace(t.TargetField) == "" {
		(*req)[chain+DownLineToHump(k)] = v
	} else {
		(*req)[chain+t.TargetField] = v
	}
	return nil
}

func RequestConvertSingleN(v interface{}, k string, t RequestConvert, req *map[string]interface{}, chain string) error {
	if strings.TrimSpace(t.TargetField) == "" {
		(*req)[chain+DownLineToHump(k)+".1"] = v
	} else {
		(*req)[chain+t.TargetField+".1"] = v
	}
	return nil
}

func RequestConvertWithN(v interface{}, k string, t RequestConvert, req *map[string]interface{}, chain string) error {
	if x, ok := v.(*schema.Set); ok {
		for i, value := range (*x).List() {
			if strings.TrimSpace(t.TargetField) == "" {
				(*req)[chain+DownLineToHump(k)+"."+strconv.Itoa(i+t.StartIndex+1)] = value
			} else {
				(*req)[chain+t.TargetField+"."+strconv.Itoa(i+t.StartIndex+1)] = value
			}
		}
	}
	if x, ok := v.([]interface{}); ok {
		for i, value := range x {
			if strings.TrimSpace(t.TargetField) == "" {
				(*req)[chain+DownLineToHump(k)+"."+strconv.Itoa(i+t.StartIndex+1)] = value
			} else {
				(*req)[chain+t.TargetField+"."+strconv.Itoa(i+t.StartIndex+1)] = value
			}
		}
	}
	if x, ok := v.(string); ok {
		for i, value := range strings.Split(x, ",") {
			if strings.TrimSpace(t.TargetField) == "" {
				(*req)[chain+DownLineToHump(k)+"."+strconv.Itoa(i+1)] = value
			} else {
				(*req)[chain+t.TargetField+"."+strconv.Itoa(i+1)] = value
			}
		}
	}
	return nil
}

func RequestConvertListN(v interface{}, k string, t RequestConvert, req *map[string]interface{}, chain string, d *schema.ResourceData, forceGet bool, single bool, contentType RequestContentType, schemaChain string, indexes []int) error {
	var (
		err   error
		isSet bool
		m     *schema.Set
		ok    bool
		list  []interface{}
	)

	if m, ok = v.(*schema.Set); ok {
		v = m.List()
		isSet = true
	}
	if list, ok = v.([]interface{}); ok {
		for index, v1 := range list {
			_index := index
			if m1, ok1 := v1.(map[string]interface{}); ok1 {
				if isSet {
					_index = m.F(m1)
				}
				for k2, v2 := range m1 {
					flag := false
					if t.NextLevelConvert != nil && t.NextLevelConvert[k2].ForceGet {
						flag = true
					} else {
						var schemaKey string
						if len(indexes) > 0 {
							schemaKey = fmt.Sprintf("%s.%d.%s", schemaChain+k, indexes[index], k2)
						} else {
							schemaKey = fmt.Sprintf("%s.%d.%s", schemaChain+k, _index, k2)
						}

						if forceGet {
							if t.ForceGet || (d.HasChange(schemaKey) && !d.IsNewResource()) {
								flag = true
							}
						} else {
							if _, ok2 := d.GetOk(schemaKey); ok2 {
								flag = true
							}
						}
					}
					if flag {
						var k3 string
						if single {
							k3 = chain + GetFinalKey(t, k, true) + "."
						} else {
							k3 = chain + GetFinalKey(t, k, true) + "." + strconv.Itoa(index+t.StartIndex+1) + "."
						}
						k4 := schemaChain + k + "." + strconv.Itoa(index) + "."
						switch reflect.TypeOf(v2).Kind() {
						case reflect.Slice:
							if t.NextLevelConvert[k2].Convert != nil {
								err = Convert(d, k2, t.NextLevelConvert[k2].Convert(d, v2), t.NextLevelConvert[k2], 0, req, k3, t.NextLevelConvert[k2].ForceGet, contentType, k4, nil)
							} else {
								err = Convert(d, k2, v2, t.NextLevelConvert[k2], 0, req, k3, t.NextLevelConvert[k2].ForceGet, contentType, k4, nil)
							}

							if err != nil {
								return err
							}
						case reflect.Ptr:
							if _v2, ok2 := v2.(*schema.Set); ok2 {
								var setIndex []int
								for _, mmm := range _v2.List() {
									setIndex = append(setIndex, _v2.F(mmm))
								}
								if t.NextLevelConvert[k2].Convert != nil {
									err = Convert(d, k2, t.NextLevelConvert[k2].Convert(d, _v2.List()), t.NextLevelConvert[k2], 0, req, k3, t.NextLevelConvert[k2].ForceGet, contentType, k4, setIndex)
								} else {
									err = Convert(d, k2, _v2.List(), t.NextLevelConvert[k2], 0, req, k3, t.NextLevelConvert[k2].ForceGet, contentType, k4, setIndex)
								}
								if err != nil {
									return err
								}
								break
							}
						default:
							k3 = k3 + GetFinalKey(t, k2, false)
							(*req)[k3] = v2
						}
					}
				}
				if single {
					break
				}
			}
		}
	}
	return nil
}

//func RequestConvertListUnique(v interface{}, k string, t RequestConvert, req *map[string]interface{}, d *schema.ResourceData, forceGet bool) error {
//	if list, ok := v.([]interface{}); ok {
//		for i, v1 := range list {
//			if m1, ok := v1.(map[string]interface{}); ok {
//				for k2, v2 := range m1 {
//					flag := false
//					schemaKey := fmt.Sprintf("%s.%d.%s", k, i, k2)
//					if forceGet {
//						if t.ForceGet || (d.HasChange(schemaKey) && !d.IsNewResource()) {
//							flag = true
//						}
//					} else {
//						if _, ok := d.GetOk(schemaKey); ok {
//							flag = true
//						}
//					}
//					if flag {
//						k3 := k + "." + k2
//						if target, ok := t.NextLevelConvert[k3]; ok {
//							(*req)[target.TargetField] = v2
//						} else {
//							(*req)[fmt.Sprintf("%s.%s", DownLineToHump(k), DownLineToHump(k2))] = v2
//						}
//					}
//				}
//				break
//			}
//		}
//	}
//	return nil
//}

func RequestConvertWithFilter(v interface{}, k string, t RequestConvert, index int, req *map[string]interface{}) (int, error) {
	if x, ok := v.([]interface{}); ok {
		v = schema.NewSet(schema.HashString, x)
	}
	if x, ok := v.(string); ok {
		if strings.TrimSpace(t.TargetField) == "" {
			(*req)["Filter."+strconv.Itoa(1)+".Name"] = DownLineToFilter(k)
		} else {
			(*req)["Filter."+strconv.Itoa(1)+".Name"] = t.TargetField
		}
		(*req)["Filter."+strconv.Itoa(index)+".Value."+strconv.Itoa(1)] = x
	}
	if x, ok := v.(*schema.Set); ok {
		for i, value := range (*x).List() {
			if i == 0 {
				if strings.TrimSpace(t.TargetField) == "" {
					(*req)["Filter."+strconv.Itoa(index)+".Name"] = DownLineToFilter(k)
				} else {
					(*req)["Filter."+strconv.Itoa(index)+".Name"] = t.TargetField
				}

			}
			(*req)["Filter."+strconv.Itoa(index)+".Value."+strconv.Itoa(i+1)] = value
		}
		index = index + 1
	}
	return index, nil
}

func RequestConvertListFilter(v interface{}, k string, t RequestConvert, index int, req *map[string]interface{}) (int, error) {
	var err error
	if list, ok := v.([]interface{}); ok {
		for _, v1 := range list {
			if m1, ok := v1.(map[string]interface{}); ok {
				for k2, v2 := range m1 {
					if v3, ok := v2.(*schema.Set); ok {
						for i, v4 := range (*v3).List() {
							if i == 0 {
								k = k + "." + k2
								if target, ok := t.NextLevelConvert[k]; ok {
									(*req)["Filter."+strconv.Itoa(index)+".Name"] = target
								} else {
									(*req)["Filter."+strconv.Itoa(index)+".Name"] = DownLineToFilter(k)
								}

							}
							(*req)["Filter."+strconv.Itoa(index)+".Value."+strconv.Itoa(i+1)] = v4
						}
						index = index + 1
					} else {
						index, err = RequestConvertListFilter(v2, k, t, index, req)
						if err != nil {
							return index, err
						}
					}
				}
			}
		}
	}
	return index, nil
}

func GetFinalKey(t RequestConvert, k string, isRoot bool) string {
	if isRoot {
		if t.TargetField == "" {
			return DownLineToHump(k)
		}
		return t.TargetField
	} else {
		if target, ok := t.NextLevelConvert[k]; ok {
			if target.TargetField == "" {
				return DownLineToHump(k)
			}
			return target.TargetField
		} else {
			return DownLineToHump(k)
		}
	}
}

func DefaultMapValue(source *map[string]interface{}, key string, defaultStruct map[string]interface{}) {
	if def, ok := (*source)[key]; !ok {
		(*source)[key] = defaultStruct
	} else {
		if ele, ok1 := def.(map[string]interface{}); ok1 {
			for k, v := range defaultStruct {
				if v1, ok3 := v.(map[string]interface{}); ok3 {
					next := (*source)[key].(map[string]interface{})
					DefaultMapValue(&next, k, v1)
				}

				if _, ok2 := ele[k]; !ok2 {
					(*source)[key].(map[string]interface{})[k] = v
				}
			}
		}
	}
}
