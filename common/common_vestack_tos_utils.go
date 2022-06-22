package common

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	TosPath     = "PATH"
	TosDomain   = "DOMAIN"
	TosHeader   = "HEADER"
	TosParam    = "PARAM"
	TosUrlParam = "URL_PARAM"
	TosResponse = "RESPONSE"
)

func convertToTosParams(convert map[string]RequestConvert, condition map[string]interface{}) (result map[string]interface{}, err error) {
	if len(condition) > 0 {
		result = map[string]interface{}{
			TosDomain:   "",
			TosHeader:   make(map[string]string),
			TosPath:     []string{},
			TosParam:    condition,
			TosUrlParam: make(map[string]string),
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
							result[TosDomain] = v1
							delete(condition, k1)
						}
					case UrlParam:
						if v1, ok := condition[k1]; ok {
							if _, ok1 := v1.(string); !ok1 {
								return result, fmt.Errorf("%s must a string type", k)
							}
							result[TosUrlParam].(map[string]string)[k1] = v1.(string)
							delete(condition, k1)
						}
					case HeaderParam:
						if v1, ok := condition[k1]; ok {
							if _, ok1 := v1.(string); !ok1 {
								return result, fmt.Errorf("%s must a string type", k)
							}
							if v1.(string) != "" {
								result[TosHeader].(map[string]string)[k1] = v1.(string)
							}
							delete(condition, k1)
						}
					case PathParam:
						if v1, ok := condition[k1]; ok {
							if _, ok1 := v1.(string); !ok1 {
								return result, fmt.Errorf("%s must a string type", k)
							}
							temp := result[TosPath].([]string)
							temp = append(temp, strconv.Itoa(v.SpecialParam.Index)+":"+v1.(string))
							result[TosPath] = temp
							delete(condition, k1)
						}
					}
				}
			}
			//sort
			if v, ok := result[TosPath]; ok {
				temp := v.([]string)
				sort.Strings(temp)
				var afterSort []string
				for _, v1 := range temp {
					afterSort = append(afterSort, v1[strings.Index(v1, ":")+1:])
				}
				result[TosPath] = afterSort
			}
			//query
			result[TosParam] = condition
		}
	}
	return result, err
}
