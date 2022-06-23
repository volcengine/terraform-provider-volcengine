package customer_gateway

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
CustomerGateway can be imported using the id, e.g.
```
$ terraform import vestack_customer_gateway.default cgw-2byswc356dybk2dx0eed2****
```

*/

func ResourceVestackCustomerGateway() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackCustomerGatewayCreate,
		Read:   resourceVestackCustomerGatewayRead,
		Update: resourceVestackCustomerGatewayUpdate,
		Delete: resourceVestackCustomerGatewayDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsIPAddress,
				Description:  "The IP address of the customer gateway.",
			},
			"customer_gateway_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the customer gateway.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the customer gateway.",
			},
		},
	}
	dataSource := DataSourceVestackCustomerGateways().Schema["customer_gateways"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVestackCustomerGatewayCreate(d *schema.ResourceData, meta interface{}) (err error) {
	customerGatewayService := NewCustomerGatewayService(meta.(*ve.SdkClient))
	err = customerGatewayService.Dispatcher.Create(customerGatewayService, d, ResourceVestackCustomerGateway())
	if err != nil {
		return fmt.Errorf("error on creating Customer Gateway %q, %s", d.Id(), err)
	}
	return resourceVestackCustomerGatewayRead(d, meta)
}

func resourceVestackCustomerGatewayRead(d *schema.ResourceData, meta interface{}) (err error) {
	customerGatewayService := NewCustomerGatewayService(meta.(*ve.SdkClient))
	err = customerGatewayService.Dispatcher.Read(customerGatewayService, d, ResourceVestackCustomerGateway())
	if err != nil {
		return fmt.Errorf("error on reading Customer Gateway %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackCustomerGatewayUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	customerGatewayService := NewCustomerGatewayService(meta.(*ve.SdkClient))
	err = customerGatewayService.Dispatcher.Update(customerGatewayService, d, ResourceVestackCustomerGateway())
	if err != nil {
		return fmt.Errorf("error on updating Customer Gateway %q, %s", d.Id(), err)
	}
	return resourceVestackCustomerGatewayRead(d, meta)
}

func resourceVestackCustomerGatewayDelete(d *schema.ResourceData, meta interface{}) (err error) {
	customerGatewayService := NewCustomerGatewayService(meta.(*ve.SdkClient))
	err = customerGatewayService.Dispatcher.Delete(customerGatewayService, d, ResourceVestackCustomerGateway())
	if err != nil {
		return fmt.Errorf("error on deleting Customer Gateway %q, %s", d.Id(), err)
	}
	return err
}
