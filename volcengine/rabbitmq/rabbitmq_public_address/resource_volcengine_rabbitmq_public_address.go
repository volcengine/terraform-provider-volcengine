package rabbitmq_public_address

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RabbitmqPublicAddress can be imported using the instance_id:eip_id, e.g.
```
$ terraform import volcengine_rabbitmq_public_address.default resource_id
```

*/

func ResourceVolcengineRabbitmqPublicAddress() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRabbitmqPublicAddressCreate,
		Read:   resourceVolcengineRabbitmqPublicAddressRead,
		Delete: resourceVolcengineRabbitmqPublicAddressDelete,
		Importer: &schema.ResourceImporter{
			State: importRabbitmqPublicAddress,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of rabbitmq instance.",
			},
			"eip_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the eip.",
			},
		},
	}
	return resource
}

func resourceVolcengineRabbitmqPublicAddressCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRabbitmqPublicAddressService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRabbitmqPublicAddress())
	if err != nil {
		return fmt.Errorf("error on creating rabbitmq_public_address %q, %s", d.Id(), err)
	}
	return resourceVolcengineRabbitmqPublicAddressRead(d, meta)
}

func resourceVolcengineRabbitmqPublicAddressRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRabbitmqPublicAddressService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRabbitmqPublicAddress())
	if err != nil {
		return fmt.Errorf("error on reading rabbitmq_public_address %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRabbitmqPublicAddressUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRabbitmqPublicAddressService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRabbitmqPublicAddress())
	if err != nil {
		return fmt.Errorf("error on updating rabbitmq_public_address %q, %s", d.Id(), err)
	}
	return resourceVolcengineRabbitmqPublicAddressRead(d, meta)
}

func resourceVolcengineRabbitmqPublicAddressDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRabbitmqPublicAddressService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRabbitmqPublicAddress())
	if err != nil {
		return fmt.Errorf("error on deleting rabbitmq_public_address %q, %s", d.Id(), err)
	}
	return err
}

func importRabbitmqPublicAddress(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form InstanceId:EipId")
	}
	err = data.Set("instance_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("eip_id", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
