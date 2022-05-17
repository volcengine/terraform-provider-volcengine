package network_interface_attach

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Network interface attach can be imported using the network_interface_id:instance_id.
```
$ terraform import vestack_network_interface_attach.default eni-bp1fg655nh68xyz9***:i-wijfn35c****
```

*/

func ResourceVestackNetworkInterfaceAttach() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVestackNetworkInterfaceAttachDelete,
		Create: resourceVestackNetworkInterfaceAttachCreate,
		Read:   resourceVestackNetworkInterfaceAttachRead,
		Importer: &schema.ResourceImporter{
			State: networkInterfaceAttachImporter,
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

func resourceVestackNetworkInterfaceAttachCreate(d *schema.ResourceData, meta interface{}) error {
	networkInterfaceAttachService := NewNetworkInterfaceAttachService(meta.(*ve.SdkClient))
	if err := networkInterfaceAttachService.Dispatcher.Create(networkInterfaceAttachService, d, ResourceVestackNetworkInterfaceAttach()); err != nil {
		return fmt.Errorf("error on creating network interface attach %q, %w", d.Id(), err)
	}
	return resourceVestackNetworkInterfaceAttachRead(d, meta)
}

func resourceVestackNetworkInterfaceAttachRead(d *schema.ResourceData, meta interface{}) error {
	networkInterfaceAttachService := NewNetworkInterfaceAttachService(meta.(*ve.SdkClient))
	if err := networkInterfaceAttachService.Dispatcher.Read(networkInterfaceAttachService, d, ResourceVestackNetworkInterfaceAttach()); err != nil {
		return fmt.Errorf("error on reading network interface attach %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVestackNetworkInterfaceAttachDelete(d *schema.ResourceData, meta interface{}) error {
	networkInterfaceAttachService := NewNetworkInterfaceAttachService(meta.(*ve.SdkClient))
	if err := networkInterfaceAttachService.Dispatcher.Delete(networkInterfaceAttachService, d, ResourceVestackNetworkInterfaceAttach()); err != nil {
		return fmt.Errorf("error on deleting network interface attach %q, %w", d.Id(), err)
	}
	return nil
}
