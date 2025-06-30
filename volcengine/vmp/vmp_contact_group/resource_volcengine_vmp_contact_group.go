package vmp_contact_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VMP Contact Group can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_contact_group.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
```

*/

func ResourceVolcengineVmpContactGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVmpContactGroupCreate,
		Read:   resourceVolcengineVmpContactGroupRead,
		Update: resourceVolcengineVmpContactGroupUpdate,
		Delete: resourceVolcengineVmpContactGroupDelete,
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
				Description: "The name of the contact group.",
			},
			"contact_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "A list of contact IDs.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of contact group.",
			},
		},
	}
	return resource
}

func resourceVolcengineVmpContactGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineVmpContactGroup())
	if err != nil {
		return fmt.Errorf("error on creating contact group %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpContactGroupRead(d, meta)
}

func resourceVolcengineVmpContactGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineVmpContactGroup())
	if err != nil {
		return fmt.Errorf("error on reading contact group %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVmpContactGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineVmpContactGroup())
	if err != nil {
		return fmt.Errorf("error on updating contact group %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpContactGroupRead(d, meta)
}

func resourceVolcengineVmpContactGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineVmpContactGroup())
	if err != nil {
		return fmt.Errorf("error on deleting contact group %q, %s", d.Id(), err)
	}
	return err
}
