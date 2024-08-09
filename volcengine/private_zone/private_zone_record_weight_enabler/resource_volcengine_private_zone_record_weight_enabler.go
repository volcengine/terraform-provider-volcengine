package private_zone_record_weight_enabler

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
PrivateZoneRecordWeightEnabler can be imported using the zid:record_set_id, e.g.
```
$ terraform import volcengine_private_zone_record_weight_enabler.default resource_id
```

*/

func ResourceVolcenginePrivateZoneRecordWeightEnabler() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivateZoneRecordWeightEnablerCreate,
		Read:   resourceVolcenginePrivateZoneRecordWeightEnablerRead,
		Update: resourceVolcenginePrivateZoneRecordWeightEnablerUpdate,
		Delete: resourceVolcenginePrivateZoneRecordWeightEnablerDelete,
		Importer: &schema.ResourceImporter{
			State: weightEnablerImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"zid": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The zid of the private zone record set.",
			},
			"record_set_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the private zone record set.",
			},
			"weight_enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether to enable the load balance of the private zone record set.",
			},
		},
	}
	return resource
}

func resourceVolcenginePrivateZoneRecordWeightEnablerCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneRecordWeightEnablerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcenginePrivateZoneRecordWeightEnabler())
	if err != nil {
		return fmt.Errorf("error on creating private_zone_record_weight_enabler %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneRecordWeightEnablerRead(d, meta)
}

func resourceVolcenginePrivateZoneRecordWeightEnablerRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneRecordWeightEnablerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcenginePrivateZoneRecordWeightEnabler())
	if err != nil {
		return fmt.Errorf("error on reading private_zone_record_weight_enabler %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcenginePrivateZoneRecordWeightEnablerUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneRecordWeightEnablerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcenginePrivateZoneRecordWeightEnabler())
	if err != nil {
		return fmt.Errorf("error on updating private_zone_record_weight_enabler %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneRecordWeightEnablerRead(d, meta)
}

func resourceVolcenginePrivateZoneRecordWeightEnablerDelete(d *schema.ResourceData, meta interface{}) (err error) {
	log.Printf("[DEBUG] deleting a volcengine_private_zone_record_weight_enabler resource only stops managing record weight, The Record Weight enabled status is left in its current state.")
	service := NewPrivateZoneRecordWeightEnablerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcenginePrivateZoneRecordWeightEnabler())
	if err != nil {
		return fmt.Errorf("error on deleting private_zone_record_weight_enabler %q, %s", d.Id(), err)
	}
	return err
}

var weightEnablerImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	zid, err := strconv.Atoi(items[0])
	if err != nil {
		return []*schema.ResourceData{data}, fmt.Errorf(" ZID cannot convert to int ")
	}
	if err := data.Set("zid", zid); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("record_set_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
