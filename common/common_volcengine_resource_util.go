package common

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceSpecial(resource *schema.Resource, data map[string]interface{}, convert map[string]ResponseConvert) {
	if len(resource.Schema) > 0 {
		mappings := map[string]string{}
		//构造一下反向映射
		if len(convert) > 0 {
			for k, v := range convert {
				if v.TargetField != "" {
					mappings[v.TargetField] = k
				}
			}
		}

		for k, v := range resource.Schema {
			var (
				key       string
				isMapping bool
				ok        bool
				val       interface{}
				next      *schema.Resource
			)
			if key, ok = mappings[k]; ok {
				isMapping = true
			} else {
				key = DownLineToHump(k)
			}
			switch v.Type {
			case schema.TypeSet:
				if val, ok = data[key]; !ok {
					setSchemaSetSpecial(key, k, data, isMapping)
				} else {
					if val == nil {
						setSchemaSetSpecial(key, k, data, isMapping)
					} else {
						break
					}
				}
				//不存在转换映射的情况下 判断一下schemaKey是否存在 不存在则set
				if !isMapping {
					if val, ok = data[k]; !ok {
						setSchemaSetSpecial(key, k, data, isMapping)
					} else {
						if val == nil {
							setSchemaSetSpecial(key, k, data, isMapping)
						}
					}
				}
			default:
				// do nothing
			}
			if v.Elem != nil {
				if next, ok = v.Elem.(*schema.Resource); ok {
					if val, ok = data[key]; ok {
						if nextData, ok2 := val.(map[string]interface{}); ok2 {
							resourceSpecial(next, nextData, convert)
							data[key] = nextData
						}
					}
				}

			}
		}
	}
}

func setSchemaSetSpecial(targetKey string, sourceKey string, data map[string]interface{}, isMapping bool) {
	data[targetKey] = new(schema.Set)
	if !isMapping {
		//不存在转换映射的情况下 自动补齐一个下划线类型的key上去
		data[sourceKey] = new(schema.Set)
	}
}
