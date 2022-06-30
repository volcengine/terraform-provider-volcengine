package common

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// HumpToDownLine Convert InstanceName -> instance_name
func HumpToDownLine(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}
	var s1 string
	if len(s) == 1 {
		s1 = strings.ToLower(s[:1])
		return s1
	}
	for k, v := range s {
		if k == 0 {
			s1 = strings.ToLower(s[0:1])
			continue
		}
		if v >= 65 && v <= 90 {
			v1 := "_" + strings.ToLower(s[k:k+1])
			s1 = s1 + v1
		} else {
			s1 = s1 + s[k:k+1]
		}
	}
	return s1
}

// DownLineToHump Convert instance_name -> InstanceName
func DownLineToHump(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}
	var s1 []string
	ss := strings.Split(s, "_")
	for _, v := range ss {
		vv := strings.ToUpper(v[:1]) + v[1:]
		s1 = append(s1, vv)
	}
	return strings.Join(s1, "")
}

// DownLineToHumpAndFirstLower Convert instance_name -> instanceName
func DownLineToHumpAndFirstLower(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}
	var s1 []string
	ss := strings.Split(s, "_")
	var index int
	for _, v := range ss {
		if index == 0 {
			s1 = append(s1, v)
		} else {
			vv := strings.ToUpper(v[:1]) + v[1:]
			s1 = append(s1, vv)
		}

		index++
	}
	return strings.Join(s1, "")
}

// DownLineToFilter instance_name ->instance-name
func DownLineToFilter(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}
	var s1 []string
	ss := strings.Split(s, "_")
	for i, v := range ss {
		if i < len(ss)-1 {
			vv := v + "-"
			s1 = append(s1, vv)
		} else {
			s1 = append(s1, v)
		}
	}
	return strings.Join(s1, "")
}

// DownLineToSpace instance_name ->instance name
func DownLineToSpace(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}
	var s1 []string
	ss := strings.Split(s, "_")
	for i, v := range ss {
		if i < len(ss)-1 {
			vv := v + " "
			s1 = append(s1, vv)
		} else {
			s1 = append(s1, v)
		}
	}
	return strings.Join(s1, "")
}

func ObtainSdkValue(keyPattern string, obj interface{}) (interface{}, error) {
	keys := strings.Split(keyPattern, ".")
	root := obj
	for index, k := range keys {
		if reflect.ValueOf(root).Kind() == reflect.Map {
			root = root.(map[string]interface{})[k]
			if root == nil {
				return root, nil
			}

		} else if reflect.ValueOf(root).Kind() == reflect.Slice {
			i, err := strconv.Atoi(k)
			if err != nil {
				return nil, fmt.Errorf("keyPattern %s index %d must number", keyPattern, index)
			}
			if len(root.([]interface{})) <= i {
				return nil, nil
			}
			root = root.([]interface{})[i]
		}
	}
	return root, nil
}

func MergeDateSourceToResource(source map[string]*schema.Schema, target *map[string]*schema.Schema) {
	for k, v := range source {
		if v1, ok := (*target)[k]; ok && (v.Elem == nil || v1.Elem == nil || v1.Optional || v1.Required) {
			continue
		}
		n := &schema.Schema{
			Type:        v.Type,
			Description: v.Description,
			Computed:    true,
		}
		if v.Elem != nil {
			switch v.Elem.(type) {
			case *schema.Resource:
				nextSource := v.Elem.(*schema.Resource).Schema
				nextTarget := make(map[string]*schema.Schema)
				if v1, ok := (*target)[k]; ok {
					if v1.Elem != nil {
						switch v1.Elem.(type) {
						case *schema.Resource:
							nextTarget = v1.Elem.(*schema.Resource).Schema
						}
					}
				}
				MergeDateSourceToResource(nextSource, &nextTarget)
				n.Elem = &schema.Resource{
					Schema: nextTarget,
				}
			case *schema.Schema:
				n.Elem = v.Elem
			}
		}

		if v.Set != nil {
			n.Set = v.Set
		}
		(*target)[k] = n
	}
}
