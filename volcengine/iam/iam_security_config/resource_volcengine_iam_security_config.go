package iam_security_config

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Iam SecurityConfig key don't support import

*/

func ResourceVolcengineIamSecurityConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineIamSecurityConfigCreate,
		Read:   resourceVolcengineIamSecurityConfigRead,
		Delete: resourceVolcengineIamSecurityConfigDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The user name.",
			},
			"safe_auth_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of safe auth, Ensure the setting scope is for a single sub-account only.",
			},
			"safe_auth_exempt_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Default:     10,
				Description: "The exempt duration of safe auth, Ensure the setting scope is for a single sub-account only.",
			},
			"safe_auth_close": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The status of safe auth, Ensure the setting scope is for a single sub-account only.",
			},
			"user_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The user id.",
			},
		},
	}
}

func resourceVolcengineIamSecurityConfigCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamSecurityConfigService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineIamSecurityConfig())
	if err != nil {
		return fmt.Errorf("error on creating security config %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamSecurityConfigRead(d, meta)
}

func resourceVolcengineIamSecurityConfigRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamSecurityConfigService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineIamSecurityConfig())
	if err != nil {
		return fmt.Errorf("error on reading security config %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineIamSecurityConfigDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamSecurityConfigService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineIamSecurityConfig())
	if err != nil {
		return fmt.Errorf("error on deleting security config %q, %s", d.Id(), err)
	}
	return err
}
