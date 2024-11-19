package image_share_permission

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ImageSharePermission can be imported using the image_id:account_id, e.g.
```
$ terraform import volcengine_image_share_permission.default resource_id
```

*/

func ResourceVolcengineImageSharePermission() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineImageSharePermissionCreate,
		Read:   resourceVolcengineImageSharePermissionRead,
		Delete: resourceVolcengineImageSharePermissionDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("image_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("account_id", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"image_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the image.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The share account id of the image.",
			},
		},
	}
	return resource
}

func resourceVolcengineImageSharePermissionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewImageSharePermissionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineImageSharePermission())
	if err != nil {
		return fmt.Errorf("error on creating image_share_permission %q, %s", d.Id(), err)
	}
	return resourceVolcengineImageSharePermissionRead(d, meta)
}

func resourceVolcengineImageSharePermissionRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewImageSharePermissionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineImageSharePermission())
	if err != nil {
		return fmt.Errorf("error on reading image_share_permission %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineImageSharePermissionUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewImageSharePermissionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineImageSharePermission())
	if err != nil {
		return fmt.Errorf("error on updating image_share_permission %q, %s", d.Id(), err)
	}
	return resourceVolcengineImageSharePermissionRead(d, meta)
}

func resourceVolcengineImageSharePermissionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewImageSharePermissionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineImageSharePermission())
	if err != nil {
		return fmt.Errorf("error on deleting image_share_permission %q, %s", d.Id(), err)
	}
	return err
}
