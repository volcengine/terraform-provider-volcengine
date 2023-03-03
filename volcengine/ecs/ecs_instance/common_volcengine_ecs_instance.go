package ecs_instance

import (
	"encoding/base64"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
