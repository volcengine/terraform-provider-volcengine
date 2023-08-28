package common

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	BypassPath         = "PATH"
	BypassDomain       = "DOMAIN"
	BypassHeader       = "HEADER"
	BypassParam        = "PARAM"
	BypassUrlParam     = "URL_PARAM"
	BypassResponse     = "RESPONSE"
	BypassFilePath     = "FILE_PATH"
	BypassResponseData = "RESPONSE_DATA"
)

func convertToBypassParams(convert map[string]RequestConvert, condition map[string]interface{}) (result map[string]interface{}, err error) {
	if len(condition) > 0 {
		result = map[string]interface{}{
			BypassDomain:   "",
			BypassHeader:   make(map[string]string),
			BypassPath:     []string{},
			BypassParam:    condition,
			BypassUrlParam: make(map[string]string),
		}

		if len(convert) > 0 {

			for k, v := range convert {
				k1 := DownLineToHump(k)
				if v.TargetField != "" {
					k1 = v.TargetField
				}
				if v.SpecialParam != nil {
					switch v.SpecialParam.Type {
					case DomainParam:
						if v1, ok := condition[k1]; ok {
							if _, ok1 := v1.(string); !ok1 {
								return result, fmt.Errorf("%s must a string type", k)
							}
							result[BypassDomain] = v1
							delete(condition, k1)
						}
					case UrlParam:
						if v1, ok := condition[k1]; ok {
							if _, ok1 := v1.(string); !ok1 {
								return result, fmt.Errorf("%s must a string type", k)
							}
							result[BypassUrlParam].(map[string]string)[k1] = v1.(string)
							delete(condition, k1)
						}
					case HeaderParam:
						if v1, ok := condition[k1]; ok {
							if _, ok1 := v1.(string); !ok1 {
								return result, fmt.Errorf("%s must a string type", k)
							}
							if v1.(string) != "" {
								result[BypassHeader].(map[string]string)[k1] = v1.(string)
							}
							delete(condition, k1)
						}
					case PathParam:
						if v1, ok := condition[k1]; ok {
							if _, ok1 := v1.(string); !ok1 {
								return result, fmt.Errorf("%s must a string type", k)
							}
							temp := result[BypassPath].([]string)
							temp = append(temp, strconv.Itoa(v.SpecialParam.Index)+":"+v1.(string))
							result[BypassPath] = temp
							delete(condition, k1)
						}
					case FilePathParam:
						if v1, ok := condition[k1]; ok {
							if _, ok1 := v1.(string); !ok1 {
								return result, fmt.Errorf("%s must a string type", k)
							}
							result[BypassFilePath] = v1
							delete(condition, k1)
						}
					}
				}
			}
			//sort
			if v, ok := result[BypassPath]; ok {
				temp := v.([]string)
				sort.Strings(temp)
				var afterSort []string
				for _, v1 := range temp {
					afterSort = append(afterSort, v1[strings.Index(v1, ":")+1:])
				}
				result[BypassPath] = afterSort
			}
			//query
			result[BypassParam] = condition
		}
	}
	return result, err
}
