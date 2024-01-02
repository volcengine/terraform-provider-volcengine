package vpc_endpoint_service_resource

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VpcEndpointServiceResource can be imported using the serviceId:resourceId, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint_service_resource.default epsvc-2fe630gurkl37k5gfuy33****:clb-bp1o94dp5i6ea****
```
It is not recommended to use this resource for binding resources, it is recommended to use the resources field of vpc_endpoint_service for binding.
If using this resource and vpc_endpoint_service jointly for operations, use lifecycle ignore_changes to suppress changes to the resources field in vpc_endpoint_service.
*/

func ResourceVolcenginePrivatelinkVpcEndpointServiceResource() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivatelinkVpcEndpointServiceResourceCreate,
		Read:   resourceVolcenginePrivatelinkVpcEndpointServiceResourceRead,
		Delete: resourceVolcenginePrivatelinkVpcEndpointServiceResourceDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("service_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("resource_id", items[1]); err != nil {
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
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The id of resource. It is not recommended to use this resource for binding resources, " +
					"it is recommended to use the resources field of vpc_endpoint_service for binding. " +
					"If using this resource and vpc_endpoint_service jointly for operations, " +
					"use lifecycle ignore_changes to suppress changes to the resources field in vpc_endpoint_service.",
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

func resourceVolcenginePrivatelinkVpcEndpointServiceResourceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(aclService, d, ResourceVolcenginePrivatelinkVpcEndpointServiceResource())
	if err != nil {
		return fmt.Errorf("error on creating vpc endpoint service resource %q, %w", d.Id(), err)
	}
	return resourceVolcenginePrivatelinkVpcEndpointServiceResourceRead(d, meta)
}

func resourceVolcenginePrivatelinkVpcEndpointServiceResourceRead(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(aclService, d, ResourceVolcenginePrivatelinkVpcEndpointServiceResource())
	if err != nil {
		return fmt.Errorf("error on reading vpc endpoint service resource %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcenginePrivatelinkVpcEndpointServiceResourceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcenginePrivatelinkVpcEndpointServiceResource())
	if err != nil {
		return fmt.Errorf("error on deleting vpc endpoint service resource %q, %w", d.Id(), err)
	}
	return err
}
