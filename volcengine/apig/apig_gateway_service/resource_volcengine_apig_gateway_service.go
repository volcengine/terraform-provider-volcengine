package apig_gateway_service

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ApigGatewayService can be imported using the id, e.g.
```
$ terraform import volcengine_apig_gateway_service.default resource_id
```

*/

func ResourceVolcengineApigGatewayService() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineApigGatewayServiceCreate,
		Read:   resourceVolcengineApigGatewayServiceRead,
		Update: resourceVolcengineApigGatewayServiceUpdate,
		Delete: resourceVolcengineApigGatewayServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The gateway id of api gateway service.",
			},
			"service_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of api gateway service.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The comments of api gateway service.",
			},
			"protocol": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The protocol of api gateway service.",
			},
			"auth_spec": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The auth spec of the api gateway service.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether the api gateway service enable auth.",
						},
					},
				},
			},

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the api gateway service.",
			},
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The error message of the api gateway service.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the api gateway service.",
			},
		},
	}
	return resource
}

func resourceVolcengineApigGatewayServiceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigGatewayServiceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineApigGatewayService())
	if err != nil {
		return fmt.Errorf("error on creating apig_gateway_service %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigGatewayServiceRead(d, meta)
}

func resourceVolcengineApigGatewayServiceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigGatewayServiceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineApigGatewayService())
	if err != nil {
		return fmt.Errorf("error on reading apig_gateway_service %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineApigGatewayServiceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigGatewayServiceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineApigGatewayService())
	if err != nil {
		return fmt.Errorf("error on updating apig_gateway_service %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigGatewayServiceRead(d, meta)
}

func resourceVolcengineApigGatewayServiceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigGatewayServiceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineApigGatewayService())
	if err != nil {
		return fmt.Errorf("error on deleting apig_gateway_service %q, %s", d.Id(), err)
	}
	return err
}
