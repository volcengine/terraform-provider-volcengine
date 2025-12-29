package vmp_contact

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VMP Contact can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_contact.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
```

*/

func ResourceVolcengineVmpContact() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVmpContactCreate,
		Read:   resourceVolcengineVmpContactRead,
		Update: resourceVolcengineVmpContactUpdate,
		Delete: resourceVolcengineVmpContactDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the contact.",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The email of the contact.",
			},
			"webhook": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The webhook of contact.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The address of webhook.",
						},
						"token": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The token of webhook.",
						},
					},
				},
			},
			"contact_group_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "A list of contact group ids.",
			},
			"lark_bot_webhook": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The lark bot webhook of contact.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The address of webhook.",
						},
						"secret_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The secret key of webhook.",
						},
					},
				},
			},
			"phone_number": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "The phone number of contact.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"country_code": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The country code of phone number. The value is `+86`.",
						},
						"number": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The number of phone number.",
						},
					},
				},
			},
			"ding_talk_bot_webhook": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The ding talk bot webhook of contact.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The address of webhook.",
						},
						"secret_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The secret key of webhook.",
						},
						"at_mobiles": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The mobiles of user.",
						},
						"at_user_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The ids of user.",
						},
					},
				},
			},
			"we_com_bot_webhook": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The we com bot webhook of contact.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The address of webhook.",
						},
						"at_user_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The ids of user.",
						},
					},
				},
			},

			// computed fields
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of contact.",
			},
			"email_active": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the email of contact active.",
			},
		},
	}
	return resource
}

func resourceVolcengineVmpContactCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineVmpContact())
	if err != nil {
		return fmt.Errorf("error on creating contact %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpContactRead(d, meta)
}

func resourceVolcengineVmpContactRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineVmpContact())
	if err != nil {
		return fmt.Errorf("error on reading contact %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVmpContactUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineVmpContact())
	if err != nil {
		return fmt.Errorf("error on updating contact %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpContactRead(d, meta)
}

func resourceVolcengineVmpContactDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineVmpContact())
	if err != nil {
		return fmt.Errorf("error on deleting contact %q, %s", d.Id(), err)
	}
	return err
}
