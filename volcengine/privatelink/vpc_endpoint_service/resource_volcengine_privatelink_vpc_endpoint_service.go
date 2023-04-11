package vpc_endpoint_service

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VpcEndpointService can be imported using the id, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint_service.default epsvc-2fe630gurkl37k5gfuy33****
```

*/

func ResourceVolcenginePrivatelinkVpcEndpointService() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivateLinkVpcEndpointServiceCreate,
		Read:   resourceVolcenginePrivateLinkVpcEndpointServiceRead,
		Update: resourceVolcenginePrivateLinkVpcEndpointServiceUpdate,
		Delete: resourceVolcenginePrivateLinkVpcEndpointServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"auto_accept_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether auto accept node connect.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of service.",
			},
			// 创建 service 时候，必须传入一个 resource；在修改 resource 的时候，必须保留一个，不能全部删除
			"resources": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The resources info. When create vpc endpoint service, the resource must exist.",
				Set:         resourceHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of resource.",
						},
						"resource_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of resource.",
						},
					},
				},
			},
			"service_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Id of service.",
			},
			"service_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of service.",
			},
			"service_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The domain of service.",
			},
			"service_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of service.",
			},
			"service_resource_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type of service.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of service.",
			},
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of service.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of service.",
			},
			"zone_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The IDs of zones.",
			},
		},
	}
	return resource
}

var resourceHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v#%v", m["resource_type"], m["resource_id"]))
	return hashcode.String(buf.String())
}

func resourceVolcenginePrivateLinkVpcEndpointServiceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(aclService, d, ResourceVolcenginePrivatelinkVpcEndpointService())
	if err != nil {
		return fmt.Errorf("error on creating vpc endpoint service %q, %w", d.Id(), err)
	}
	return resourceVolcenginePrivateLinkVpcEndpointServiceRead(d, meta)
}

func resourceVolcenginePrivateLinkVpcEndpointServiceRead(d *schema.ResourceData, meta interface{}) (err error) {
	aclService := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(aclService, d, ResourceVolcenginePrivatelinkVpcEndpointService())
	if err != nil {
		return fmt.Errorf("error on reading vpc endpoint service %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcenginePrivateLinkVpcEndpointServiceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcenginePrivatelinkVpcEndpointService())
	if err != nil {
		return fmt.Errorf("error on updating vpc endoint service %q, %w", d.Id(), err)
	}
	return resourceVolcenginePrivateLinkVpcEndpointServiceRead(d, meta)
}

func resourceVolcenginePrivateLinkVpcEndpointServiceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcenginePrivatelinkVpcEndpointService())
	if err != nil {
		return fmt.Errorf("error on deleting vpc endpoint service%q, %w", d.Id(), err)
	}
	return err
}
