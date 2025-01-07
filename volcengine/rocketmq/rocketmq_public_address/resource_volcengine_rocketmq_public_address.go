package rocketmq_public_address

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RocketmqPublicAddress can be imported using the instance_id:eip_id, e.g.
```
$ terraform import volcengine_rocketmq_public_address.default resource_id
```

*/

func ResourceVolcengineRocketmqPublicAddress() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRocketmqPublicAddressCreate,
		Read:   resourceVolcengineRocketmqPublicAddressRead,
		Delete: resourceVolcengineRocketmqPublicAddressDelete,
		Importer: &schema.ResourceImporter{
			State: importRocketmqPublicAddress,
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
				Description: "The id of rocketmq instance.",
			},
			"eip_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the eip.",
			},
			"ssl_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The ssl mode of the rocketmq instance. Valid values: `enforcing`, `permissive`. Default is `permissive`.",
			},
		},
	}
	return resource
}

func resourceVolcengineRocketmqPublicAddressCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqPublicAddressService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRocketmqPublicAddress())
	if err != nil {
		return fmt.Errorf("error on creating rocketmq_public_address %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqPublicAddressRead(d, meta)
}

func resourceVolcengineRocketmqPublicAddressRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqPublicAddressService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRocketmqPublicAddress())
	if err != nil {
		return fmt.Errorf("error on reading rocketmq_public_address %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRocketmqPublicAddressUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqPublicAddressService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRocketmqPublicAddress())
	if err != nil {
		return fmt.Errorf("error on updating rocketmq_public_address %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqPublicAddressRead(d, meta)
}

func resourceVolcengineRocketmqPublicAddressDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqPublicAddressService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRocketmqPublicAddress())
	if err != nil {
		return fmt.Errorf("error on deleting rocketmq_public_address %q, %s", d.Id(), err)
	}
	return err
}

func importRocketmqPublicAddress(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
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
