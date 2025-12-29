package vmp_silence_policy_enable_disable

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The VmpSilencePolicyEnableDisable is not support import.


*/

// ResourceVolcengineVmpSilencePolicyEnableDisable 定义 VMP 静默策略启用/禁用资源 Schema 与 CRUD 方法
func ResourceVolcengineVmpSilencePolicyEnableDisable() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVmpSilencePolicyEnableDisableCreate,
		Read:   resourceVolcengineVmpSilencePolicyEnableDisableRead,
		Delete: resourceVolcengineVmpSilencePolicyEnableDisableDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
	return resource
}

// resourceVolcengineVmpSilencePolicyEnableDisableCreate 执行启用动作（EnableSilencePolicies）并设置资源状态
func resourceVolcengineVmpSilencePolicyEnableDisableCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpSilencePolicyEnableDisableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVmpSilencePolicyEnableDisable())
	if err != nil {
		return fmt.Errorf("error on creating vmp_silence_policy_enable_disable %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpSilencePolicyEnableDisableRead(d, meta)
}

// resourceVolcengineVmpSilencePolicyEnableDisableRead 读取指定静默策略的当前状态
func resourceVolcengineVmpSilencePolicyEnableDisableRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpSilencePolicyEnableDisableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVmpSilencePolicyEnableDisable())
	if err != nil {
		return fmt.Errorf("error on reading vmp_silence_policy_enable_disable %q, %s", d.Id(), err)
	}
	return err
}

// resourceVolcengineVmpSilencePolicyEnableDisableDelete 执行禁用动作（DisableSilencePolicies）
func resourceVolcengineVmpSilencePolicyEnableDisableDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpSilencePolicyEnableDisableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVmpSilencePolicyEnableDisable())
	if err != nil {
		return fmt.Errorf("error on deleting vmp_silence_policy_enable_disable %q, %s", d.Id(), err)
	}
	return err
}
