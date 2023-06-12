package vpc_endpoint_connection

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
PrivateLink Vpc Endpoint Connection Service can be imported using the endpoint id and service id, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint_connection.default ep-3rel74u229dz45zsk2i6l69qa:epsvc-2byz5mykk9y4g2dx0efs4aqz3
```

*/

func ResourceVolcenginePrivatelinkVpcEndpointConnectionService() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivatelinkVpcEndpointConnectionCreate,
		Read:   resourceVolcenginePrivatelinkVpcEndpointConnectionRead,
		Delete: resourceVolcenginePrivatelinkVpcEndpointConnectionDelete,
		Importer: &schema.ResourceImporter{
			State: vpcConnectionImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"endpoint_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the endpoint.",
			},
			"service_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the security group.",
			},
		},
	}
	dataSource := DataSourceVolcenginePrivatelinkVpcEndpointConnections().Schema["connections"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcenginePrivatelinkVpcEndpointConnectionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcEndpointConnectionService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcenginePrivatelinkVpcEndpointConnectionService())
	if err != nil {
		return fmt.Errorf("error on creating private link VpcEndpointConnection service %q, %w", d.Id(), err)
	}
	return resourceVolcenginePrivatelinkVpcEndpointConnectionRead(d, meta)
}

func resourceVolcenginePrivatelinkVpcEndpointConnectionRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcEndpointConnectionService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcenginePrivatelinkVpcEndpointConnectionService())
	if err != nil {
		return fmt.Errorf("error on reading private link VpcEndpointConnection service %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcenginePrivatelinkVpcEndpointConnectionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcEndpointConnectionService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcenginePrivatelinkVpcEndpointConnectionService())
	if err != nil {
		return fmt.Errorf("error on deleting private link VpcEndpointConnection service %q, %w", d.Id(), err)
	}
	return nil
}

var vpcConnectionImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("endpoint_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("service_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
