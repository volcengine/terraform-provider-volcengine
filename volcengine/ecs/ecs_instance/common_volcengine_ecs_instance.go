package ecs_instance

import (
	"encoding/base64"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func EcsInstanceImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if k == "force_restart" {
		return true
	}
	//由于一些字段暂时无法支持从查询中返回 所以现在设立做特殊处理拦截变更 用来适配导入的场景 后续支持后在对导入场景做优化 此模式会导致不一致问题 去除
	//if d.Id() != "" {
	//	if k == "security_enhancement_strategy" {
	//		return true
	//	}
	//	if k == "auto_renew" {
	//		return true
	//	}
	//	if k == "auto_renew_period" {
	//		return true
	//	}
	//}

	if d.Id() == "" {
		if k == "include_data_volumes" {
			return true
		}
	}

	//在计费方式没有发生变化的时候 period的变化会被忽略
	if !d.HasChange("instance_charge_type") && (k == "period" || k == "include_data_volumes") {
		return true
	}

	if d.Get("instance_charge_type").(string) == "PostPaid" && (k == "period" || k == "period_unit") {
		return true
	}

	return false
}

func UserDateImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	if k == "user_data" {
		_, base64DecodeError := base64.StdEncoding.DecodeString(new)
		if base64DecodeError != nil {
			return false
		}
		v := base64.StdEncoding.EncodeToString([]byte(old))
		if v == new {
			return true
		}

	}
	return false
}

func RemoveSystemTags(data []interface{}) ([]interface{}, error) {
	var (
		ok      bool
		result  map[string]interface{}
		results []interface{}
		tags    []interface{}
	)
	for _, d := range data {
		if result, ok = d.(map[string]interface{}); !ok {
			return results, errors.New("The elements in data are not map ")
		}
		tags, ok = result["Tags"].([]interface{})
		if ok {
			tags = ve.FilterSystemTags(tags)
			result["Tags"] = tags
		}
		results = append(results, result)
	}
	return results, nil
}
