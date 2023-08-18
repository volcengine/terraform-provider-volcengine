package common

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Callback struct {
	Call SdkCall
	Err  error
}

type SdkCall struct {
	Action           string
	BeforeCall       BeforeCallFunc
	ExecuteCall      ExecuteCallFunc
	CallError        CallErrorFunc
	AfterCall        AfterCallFunc
	Convert          map[string]RequestConvert
	ConvertMode      RequestConvertMode
	SdkParam         *map[string]interface{}
	RequestIdField   string
	Refresh          *StateRefresh
	ExtraRefresh     map[ResourceService]*StateRefresh
	AfterRefresh     CallFunc
	ContentType      RequestContentType
	LockId           LockId
	AfterLocked      CallFunc
	ServiceCategory  ServiceCategory
	ExtraRefreshCall ExtraRefreshFunc

	//common inner use
	refreshState func(*schema.ResourceData, []string, time.Duration, string) *resource.StateChangeConf
}

type StateRefresh struct {
	Target     []string
	Timeout    time.Duration
	ResourceId string
}

type CallErrorFunc func(d *schema.ResourceData, client *SdkClient, call SdkCall, baseErr error) error
type ExecuteCallFunc func(d *schema.ResourceData, client *SdkClient, call SdkCall) (*map[string]interface{}, error)
type AfterCallFunc func(d *schema.ResourceData, client *SdkClient, resp *map[string]interface{}, call SdkCall) error
type BeforeCallFunc func(d *schema.ResourceData, client *SdkClient, call SdkCall) (bool, error)
type ReadResourceFunc func(d *schema.ResourceData, resourceId string) (map[string]interface{}, error)
type CallFunc func(d *schema.ResourceData, client *SdkClient, call SdkCall) error
type ExtraRefreshFunc func(d *schema.ResourceData, client *SdkClient, call SdkCall) (map[ResourceService]*StateRefresh, error)

type LockId func(d *schema.ResourceData) string

func (c *SdkCall) InitWriteCall(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) (err error) {
	var (
		param map[string]interface{}
	)
	param, err = ResourceDateToRequest(resourceData, resource, isUpdate, c.Convert, c.ConvertMode, c.ContentType)
	if err != nil {
		return err
	}

	if c.SdkParam != nil {
		for k, v := range param {
			(*c.SdkParam)[k] = v
		}
	} else {
		c.SdkParam = &param
	}

	if isUpdate && c.RequestIdField != "" {
		(*c.SdkParam)[c.RequestIdField] = resourceData.Id()
	}
	return err
}

func volcengineSort(source []string) (result []string) {
	mapping := make(map[string]int)
	for _, k := range source {
		if strings.Contains(k, ".") {
			keys := strings.Split(k, ".")
			var key string
			for _, v1 := range keys {
				if n, err := strconv.Atoi(v1); err != nil {
					key = key + v1 + "."
				} else {
					if _n, ok := mapping[key]; !ok {
						mapping[key] = len(strconv.Itoa(n))
					} else {
						if len(strconv.Itoa(n)) > _n {
							mapping[key] = len(strconv.Itoa(n))
						}
					}
				}
			}

		}
	}
	for _, k := range source {
		if strings.Contains(k, ".") {
			keys := strings.Split(k, ".")
			var (
				key    string
				newKey string
			)
			for _, v1 := range keys {
				if n, err := strconv.Atoi(v1); err != nil {
					key = key + v1 + "."
					newKey = newKey + v1 + "."
				} else {
					if len(strconv.Itoa(n)) < mapping[key] {
						var zero string
						for i := 0; i < mapping[key]-len(strconv.Itoa(n)); i++ {
							zero = zero + "0"
						}
						newKey = newKey + zero + strconv.Itoa(n) + "."
					} else {
						newKey = newKey + strconv.Itoa(n) + "."
					}
				}
			}
			result = append(result, newKey[0:len(newKey)-1])
		} else {
			result = append(result, k)
		}
	}

	sort.Strings(result)

	//去除补位
	for i, k := range result {
		if strings.Contains(k, ".") {
			keys := strings.Split(k, ".")
			var (
				newKey string
			)
			for _, v1 := range keys {
				if strings.HasPrefix(v1, "0") {
					b := []byte(v1)
					var (
						newB      []byte
						hasNoZero bool
					)

					for _, v2 := range b {
						if v2 != '0' && !hasNoZero {
							hasNoZero = true
							newB = append(newB, v2)
						} else if hasNoZero {
							newB = append(newB, v2)
						}
					}
					newKey = newKey + string(newB) + "."

				} else {
					newKey = newKey + v1 + "."
				}
			}
			result[i] = newKey[0 : len(newKey)-1]
		}
	}

	return result
}

func SortAndStartTransJson(source map[string]interface{}) (map[string]interface{}, error) {
	target := make(map[string]interface{})
	var a []string
	for k := range source {
		a = append(a, k)
	}
	sort.Strings(a)

	a = volcengineSort(a)

	for _, k := range a {
		k1, v1, err := transToJson(k, source[k], "", &target)
		if err != nil {
			return target, err
		}
		target[k1] = v1
	}
	return target, nil
}

func transToJson(key string, value interface{}, chain string, top *map[string]interface{}) (string, interface{}, error) {
	var (
		index int
		err   error
		k     string
		v     interface{}
		v1    []interface{}
		v2    map[string]interface{}
	)
	if strings.Contains(key, ".") {
		keys := strings.Split(key, ".")
		nextKey := key[len(keys[0])+1:]
		if chain == "" {
			chain = keys[0]
		} else {
			index, err = strconv.Atoi(keys[0])
			if err != nil {
				chain = chain + "." + keys[0]
			} else {
				chain = chain + "." + strconv.Itoa(index-1)
			}
		}
		k, v, err = transToJson(nextKey, value, chain, top)
		if err != nil {
			return k, v, err
		}
		index, err = strconv.Atoi(k)
		if err == nil {
			v1, err = getAndSetSlice(chain, index-1, v, top)
			return keys[0], v1, err
		} else {
			v2, err = getAndSetMap(chain, k, v, top)
			return keys[0], v2, err
		}
	} else {
		return key, value, nil
	}
}

func getAndSetSlice(pattern string, index int, value interface{}, top *map[string]interface{}) ([]interface{}, error) {
	exist, _ := ObtainSdkValue(pattern, *top)
	if exist != nil {
		exist1, _ := ObtainSdkValue(pattern+"."+strconv.Itoa(index), *top)
		if exist1 == nil {
			if index > len(exist.([]interface{})) {
				return exist.([]interface{}), fmt.Errorf("%s index is not serial", pattern+"."+strconv.Itoa(index+1))
			}
			return append(exist.([]interface{}), value), nil
		}
		return exist.([]interface{}), nil
	}
	return []interface{}{value}, nil
}

func getAndSetMap(pattern string, key string, value interface{}, top *map[string]interface{}) (map[string]interface{}, error) {
	exist, _ := ObtainSdkValue(pattern, *top)
	if exist != nil {
		next := exist.(map[string]interface{})
		next[key] = value
		return next, nil
	}
	return map[string]interface{}{
		key: value,
	}, nil
}

func (c *SdkCall) InitReadCall(resourceData *schema.ResourceData, resource *schema.Resource) (err error) {
	var param map[string]interface{}
	param, err = ResourceDateToRequest(resourceData, resource, false, c.Convert, RequestConvertInConvert, c.ContentType)
	if err != nil {
		return err
	}
	if c.SdkParam != nil {
		for k, v := range param {
			(*c.SdkParam)[k] = v
		}
	} else {
		c.SdkParam = &param
	}
	return err
}

func CallProcess(calls []SdkCall, d *schema.ResourceData, client *SdkClient, service ResourceService) (err error) {
	if len(calls) > 0 {
		for _, fn := range calls {
			if fn.ExecuteCall != nil {
				var (
					resp *map[string]interface{}
				)
				doExecute := true

				switch fn.ServiceCategory {
				case ServiceBypass:
					var trans map[string]interface{}
					trans, err = convertToBypassParams(fn.Convert, *fn.SdkParam)
					if err != nil {
						return err
					}
					fn.SdkParam = &trans
				case DefaultServiceCategory:
					//do nothing
				}

				if fn.BeforeCall != nil {
					doExecute, err = fn.BeforeCall(d, client, fn)
				}
				if doExecute {
					switch fn.ContentType {
					case ContentTypeDefault:
						//do nothing
					case ContentTypeJson:
						var jsonParam map[string]interface{}
						jsonParam, err = SortAndStartTransJson(*fn.SdkParam)
						if err != nil {
							return err
						}
						fn.SdkParam = &jsonParam
					}
					if fn.LockId != nil {
						key := fn.LockId(d)
						if key != "" {
							TryLock(key)
						}
					}
					if fn.AfterLocked != nil {
						err = fn.AfterLocked(d, client, fn)
					}
					if err == nil {
						resp, err = fn.ExecuteCall(d, client, fn)
					}
				}
				if err != nil {
					if fn.CallError != nil {
						err = fn.CallError(d, client, fn, err)
					}
				}
				if doExecute && fn.AfterCall != nil && err == nil {
					err = fn.AfterCall(d, client, resp, fn)
				}

				// WaitForState
				if doExecute && fn.Refresh != nil && err == nil {
					var stateConf *resource.StateChangeConf
					if fn.refreshState != nil {
						stateConf = fn.refreshState(d, fn.Refresh.Target, fn.Refresh.Timeout, d.Id())
					} else {
						stateConf = service.RefreshResourceState(d, fn.Refresh.Target, fn.Refresh.Timeout, d.Id())
					}
					if stateConf != nil {
						_, err = stateConf.WaitForState()
					}
				}

				if doExecute && fn.ExtraRefreshCall != nil && err == nil {
					var refreshMap map[ResourceService]*StateRefresh
					refreshMap, err = fn.ExtraRefreshCall(d, client, fn)
					if fn.ExtraRefresh == nil {
						fn.ExtraRefresh = refreshMap
					} else {
						for k, v := range fn.ExtraRefresh {
							fn.ExtraRefresh[k] = v
						}
					}
				}

				if doExecute && fn.ExtraRefresh != nil && err == nil {
					for k, v := range fn.ExtraRefresh {
						stateConf := k.RefreshResourceState(d, v.Target, v.Timeout, v.ResourceId)
						if stateConf != nil {
							_, err = stateConf.WaitForState()
						}
					}
				}

				if doExecute && fn.AfterRefresh != nil && err == nil {
					err = fn.AfterRefresh(d, client, fn)
				}

				if doExecute && fn.LockId != nil {
					key := fn.LockId(d)
					if key != "" {
						ReleaseLock(key)
					}
				}
				if err != nil {
					return err
				}
			}

		}
	}
	return err
}

func CheckResourceUtilRemoved(d *schema.ResourceData, readResourceFunc ReadResourceFunc, timeout time.Duration) error {
	return resource.Retry(timeout, func() *resource.RetryError {
		_, callErr := readResourceFunc(d, d.Id())
		// 能查询成功代表还在删除中，重试
		if callErr == nil {
			return resource.RetryableError(fmt.Errorf("resource still in removing status "))
		} else {
			if ResourceNotFoundError(callErr) {
				return nil
			} else {
				return resource.NonRetryableError(callErr)
			}
		}
	})
}
