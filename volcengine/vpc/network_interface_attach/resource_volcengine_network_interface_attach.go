package network_interface_attach

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Network interface attach can be imported using the network_interface_id:instance_id.
```
$ terraform import volcengine_network_interface_attach.default eni-bp1fg655nh68xyz9***:i-wijfn35c****
```

*/

func ResourceVolcengineNetworkInterfaceAttach() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVolcengineNetworkInterfaceAttachDelete,
		Create: resourceVolcengineNetworkInterfaceAttachCreate,
		Read:   resourceVolcengineNetworkInterfaceAttachRead,
		Importer: &schema.ResourceImporter{
			State: networkInterfaceAttachImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"network_interface_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the ENI.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the instance to which the ENI is bound.",
			},
		},
	}
}

func resourceVolcengineNetworkInterfaceAttachCreate(d *schema.ResourceData, meta interface{}) error {
	networkInterfaceAttachService := NewNetworkInterfaceAttachService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(networkInterfaceAttachService, d, ResourceVolcengineNetworkInterfaceAttach()); err != nil {
		return fmt.Errorf("error on creating network interface attach %q, %w", d.Id(), err)
	}
	return resourceVolcengineNetworkInterfaceAttachRead(d, meta)
}

func resourceVolcengineNetworkInterfaceAttachRead(d *schema.ResourceData, meta interface{}) error {
	networkInterfaceAttachService := NewNetworkInterfaceAttachService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(networkInterfaceAttachService, d, ResourceVolcengineNetworkInterfaceAttach()); err != nil {
		return fmt.Errorf("error on reading network interface attach %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineNetworkInterfaceAttachDelete(d *schema.ResourceData, meta interface{}) error {
	networkInterfaceAttachService := NewNetworkInterfaceAttachService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(networkInterfaceAttachService, d, ResourceVolcengineNetworkInterfaceAttach()); err != nil {
		return fmt.Errorf("error on deleting network interface attach %q, %w", d.Id(), err)
	}
	return nil
}
