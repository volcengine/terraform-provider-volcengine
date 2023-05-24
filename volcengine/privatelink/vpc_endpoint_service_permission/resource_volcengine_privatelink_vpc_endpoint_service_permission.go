package vpc_endpoint_service_permission

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VpcEndpointServicePermission can be imported using the serviceId:permitAccountId, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint_service_permission.default epsvc-2fe630gurkl37k5gfuy33****:2100000000
```

*/

func ResourceVolcenginePrivatelinkVpcEndpointServicePermission() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivatelinkVpcEndpointServicePermissionCreate,
		Read:   resourceVolcenginePrivatelinkVpcEndpointServicePermissionRead,
		Delete: resourceVolcenginePrivatelinkVpcEndpointServicePermissionDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("service_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("permit_account_id", items[1]); err != nil {
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
			"permit_account_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of account.",
			},
			"service_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of service.",
			},
		},
	}
	return resource
}

func resourceVolcenginePrivatelinkVpcEndpointServicePermissionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(aclService, d, ResourceVolcenginePrivatelinkVpcEndpointServicePermission())
	if err != nil {
		return fmt.Errorf("error on creating vpc endpoint service permission %q, %w", d.Id(), err)
	}
	return resourceVolcenginePrivatelinkVpcEndpointServicePermissionRead(d, meta)
}

func resourceVolcenginePrivatelinkVpcEndpointServicePermissionRead(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(aclService, d, ResourceVolcenginePrivatelinkVpcEndpointServicePermission())
	if err != nil {
		return fmt.Errorf("error on reading vpc endpoint service permission %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcenginePrivatelinkVpcEndpointServicePermissionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcenginePrivatelinkVpcEndpointServicePermission())
	if err != nil {
		return fmt.Errorf("error on deleting vpc endpoint service permission %q, %w", d.Id(), err)
	}
	return err
}
