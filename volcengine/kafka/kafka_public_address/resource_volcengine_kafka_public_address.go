package kafka_public_address

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KafkaPublicAddress can be imported using the instance_id:eip_id, e.g.
```
$ terraform import volcengine_kafka_public_address.default instance_id:eip_id
```

*/

func ResourceVolcengineKafkaPublicAddress() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKafkaPublicAddressCreate,
		Read:   resourceVolcengineKafkaPublicAddressRead,
		Delete: resourceVolcengineKafkaPublicAddressDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("eip_id", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("instance_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of kafka instance.",
				ForceNew:    true,
			},
			"eip_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of eip.",
				ForceNew:    true,
			},
			"endpoint_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The endpoint type of instance.",
			},
			"network_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The network type of instance.",
			},
			"public_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The public endpoint of instance.",
			},
		},
	}
	return resource
}

func resourceVolcengineKafkaPublicAddressCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaInternetEnablerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKafkaPublicAddress())
	if err != nil {
		return fmt.Errorf("error on creating kafka public address %q, %s", d.Id(), err)
	}
	return resourceVolcengineKafkaPublicAddressRead(d, meta)
}

func resourceVolcengineKafkaPublicAddressRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaInternetEnablerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKafkaPublicAddress())
	if err != nil {
		return fmt.Errorf("error on reading kafka public address %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKafkaPublicAddressDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaInternetEnablerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKafkaPublicAddress())
	if err != nil {
		return fmt.Errorf("error on deleting kafka public address %q, %s", d.Id(), err)
	}
	return err
}
