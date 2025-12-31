package vmp_alerting_rule_enable_disable

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The VmpAlertingRuleEnableDisable is not support import.


*/

// ResourceVolcengineVmpAlertingRuleEnableDisable 定义 VMP 告警规则启用/禁用资源 Schema 与 CRUD 方法
func ResourceVolcengineVmpAlertingRuleEnableDisable() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVmpAlertingRuleEnableDisableCreate,
		Read:   resourceVolcengineVmpAlertingRuleEnableDisableRead,
		Delete: resourceVolcengineVmpAlertingRuleEnableDisableDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The ids of alerting rule.",
			},
			// computed 字段不返回，直接在 aftercall 中处理
			// "successful_items": {
			// 	Type:     schema.TypeSet,
			// 	Computed: true,
			// 	Elem:     &schema.Schema{Type: schema.TypeString},
			// },
			// "unsuccessful_items": {
			// 	Type:     schema.TypeList,
			// 	Computed: true,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"id":            {Type: schema.TypeString, Computed: true},
			// 			"error_code":    {Type: schema.TypeString, Computed: true},
			// 			"error_message": {Type: schema.TypeString, Computed: true},
			// 		},
			// 	},
			// },
		},
	}
	return resource
}

func resourceVolcengineVmpAlertingRuleEnableDisableCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpAlertingRuleEnableDisableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVmpAlertingRuleEnableDisable())
	if err != nil {
		return fmt.Errorf("error on creating vmp_alerting_rule_enable_disable %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpAlertingRuleEnableDisableRead(d, meta)
}

func resourceVolcengineVmpAlertingRuleEnableDisableRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpAlertingRuleEnableDisableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVmpAlertingRuleEnableDisable())
	if err != nil {
		return fmt.Errorf("error on reading vmp_alerting_rule_enable_disable %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVmpAlertingRuleEnableDisableDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpAlertingRuleEnableDisableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVmpAlertingRuleEnableDisable())
	if err != nil {
		return fmt.Errorf("error on deleting vmp_alerting_rule_enable_disable %q, %s", d.Id(), err)
	}
	return err
}
