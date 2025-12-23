package rds_postgresql_endpoint_public_address

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlEndpointPublicAddress can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_endpoint_public_address.default resource_id
```

*/

func ResourceVolcengineRdsPostgresqlEndpointPublicAddress() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlEndpointPublicAddressCreate,
		Read:   resourceVolcengineRdsPostgresqlEndpointPublicAddressRead,
		Update: resourceVolcengineRdsPostgresqlEndpointPublicAddressUpdate,
		Delete: resourceVolcengineRdsPostgresqlEndpointPublicAddressDelete,
		Importer: &schema.ResourceImporter{State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
			parts := strings.Split(d.Id(), ":")
			if len(parts) != 3 {
				return []*schema.ResourceData{d}, fmt.Errorf("import id must be 'instance_id:endpoint_id:eip_id'")
			}
			_ = d.Set("instance_id", parts[0])
			_ = d.Set("endpoint_id", parts[1])
			_ = d.Set("eip_id", parts[2])
			return []*schema.ResourceData{d}, nil
		}},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the RDS PostgreSQL instance.",
			},
			"endpoint_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Endpoint ID.",
			},
			"eip_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "EIP ID to bind for public access.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlEndpointPublicAddressCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlEndpointPublicAddressService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlEndpointPublicAddress())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_endpoint_public_address %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlEndpointPublicAddressRead(d, meta)
}

func resourceVolcengineRdsPostgresqlEndpointPublicAddressRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlEndpointPublicAddressService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlEndpointPublicAddress())
	if err != nil {
		return fmt.Errorf("error on reading rds_postgresql_endpoint_public_address %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlEndpointPublicAddressUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlEndpointPublicAddressService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsPostgresqlEndpointPublicAddress())
	if err != nil {
		return fmt.Errorf("error on updating rds_postgresql_endpoint_public_address %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlEndpointPublicAddressRead(d, meta)
}

func resourceVolcengineRdsPostgresqlEndpointPublicAddressDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlEndpointPublicAddressService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlEndpointPublicAddress())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_endpoint_public_address %q, %s", d.Id(), err)
	}
	return err
}
