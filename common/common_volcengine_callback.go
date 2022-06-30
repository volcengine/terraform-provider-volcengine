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
	Action         string
	BeforeCall     BeforeCallFunc
	ExecuteCall    ExecuteCallFunc
	CallError      CallErrorFunc
	AfterCall      AfterCallFunc
	Convert        map[string]RequestConvert
	ConvertMode    RequestConvertMode
	SdkParam       *map[string]interface{}
	RequestIdField string
	Refresh        *StateRefresh
	ExtraRefresh   map[ResourceService]*StateRefresh
	ContentType    RequestContentType
	LockId         LockId
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

func sortAndStartTransJson(source map[string]interface{}) map[string]interface{} {
	target := make(map[string]interface{})
	var a []string
	for k := range source {
		a = append(a, k)
	}
	sort.Strings(a)

	for _, k := range a {
		k1, v1 := transToJson(k, source[k], "", &target)
		target[k1] = v1
	}
	return target
}

func transToJson(key string, value interface{}, chain string, top *map[string]interface{}) (string, interface{}) {
	var (
		index int
		err   error
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
		k, v := transToJson(nextKey, value, chain, top)
		index, err = strconv.Atoi(k)
		if err == nil {
			return keys[0], getAndSetSlice(chain, index-1, v, top)
		} else {
			return keys[0], getAndSetMap(chain, k, v, top)
		}
	} else {
		return key, value
	}
}

func getAndSetSlice(pattern string, index int, value interface{}, top *map[string]interface{}) []interface{} {
	exist, _ := ObtainSdkValue(pattern, *top)
	if exist != nil {
		exist1, _ := ObtainSdkValue(pattern+"."+strconv.Itoa(index), *top)
		if exist1 == nil {
			return append(exist.([]interface{}), value)
		}
		return exist.([]interface{})
	}
	return []interface{}{value}
}

func getAndSetMap(pattern string, key string, value interface{}, top *map[string]interface{}) map[string]interface{} {
	exist, _ := ObtainSdkValue(pattern, *top)
	if exist != nil {
		next := exist.(map[string]interface{})
		next[key] = value
		return next
	}
	return map[string]interface{}{
		key: value,
	}
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
	if calls != nil {
		for _, fn := range calls {
			if fn.ExecuteCall != nil {
				var (
					resp *map[string]interface{}
				)
				doExecute := true
				if fn.BeforeCall != nil {
					doExecute, err = fn.BeforeCall(d, client, fn)
				}
				if doExecute {
					switch fn.ContentType {
					case ContentTypeDefault:
						break
					case ContentTypeJson:
						jsonParam := sortAndStartTransJson(*fn.SdkParam)
						fn.SdkParam = &jsonParam
						break
					}
					if fn.LockId != nil {
						key := fn.LockId(d)
						if key != "" {
							TryLock(key)
						}
					}
					resp, err = fn.ExecuteCall(d, client, fn)
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
					stateConf := service.RefreshResourceState(d, fn.Refresh.Target, fn.Refresh.Timeout, d.Id())
					if stateConf != nil {
						_, err = stateConf.WaitForState()
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
