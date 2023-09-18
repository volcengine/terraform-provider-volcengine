package customer_gateway

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CustomerGateway can be imported using the id, e.g.
```
$ terraform import volcengine_customer_gateway.default cgw-2byswc356dybk2dx0eed2****
```

*/

func ResourceVolcengineCustomerGateway() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCustomerGatewayCreate,
		Read:   resourceVolcengineCustomerGatewayRead,
		Update: resourceVolcengineCustomerGatewayUpdate,
		Delete: resourceVolcengineCustomerGatewayDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
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
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the VPN customer gateway.",
			},
		},
	}
	dataSource := DataSourceVolcengineCustomerGateways().Schema["customer_gateways"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineCustomerGatewayCreate(d *schema.ResourceData, meta interface{}) (err error) {
	customerGatewayService := NewCustomerGatewayService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(customerGatewayService, d, ResourceVolcengineCustomerGateway())
	if err != nil {
		return fmt.Errorf("error on creating Customer Gateway %q, %s", d.Id(), err)
	}
	return resourceVolcengineCustomerGatewayRead(d, meta)
}

func resourceVolcengineCustomerGatewayRead(d *schema.ResourceData, meta interface{}) (err error) {
	customerGatewayService := NewCustomerGatewayService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(customerGatewayService, d, ResourceVolcengineCustomerGateway())
	if err != nil {
		return fmt.Errorf("error on reading Customer Gateway %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCustomerGatewayUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	customerGatewayService := NewCustomerGatewayService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(customerGatewayService, d, ResourceVolcengineCustomerGateway())
	if err != nil {
		return fmt.Errorf("error on updating Customer Gateway %q, %s", d.Id(), err)
	}
	return resourceVolcengineCustomerGatewayRead(d, meta)
}

func resourceVolcengineCustomerGatewayDelete(d *schema.ResourceData, meta interface{}) (err error) {
	customerGatewayService := NewCustomerGatewayService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(customerGatewayService, d, ResourceVolcengineCustomerGateway())
	if err != nil {
		return fmt.Errorf("error on deleting Customer Gateway %q, %s", d.Id(), err)
	}
	return err
}
