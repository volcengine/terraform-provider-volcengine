package private_zone_record

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
PrivateZoneRecord can be imported using the id, e.g.
```
$ terraform import volcengine_private_zone_record.default resource_id
```

*/

func ResourceVolcenginePrivateZoneRecord() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivateZoneRecordCreate,
		Read:   resourceVolcenginePrivateZoneRecordRead,
		Update: resourceVolcenginePrivateZoneRecordUpdate,
		Delete: resourceVolcenginePrivateZoneRecordDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
				Description: "The zid of the private zone record.",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The host of the private zone record.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the private zone record. Valid values: `A`, `AAAA`, `CNAME`, `MX`, `PTR`.",
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The value of the private zone record. Record values need to be set based on the value of the `type`.",
			},
			"weight": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The weight of the private zone record. This field is only effected when the `load_balance_mode` of the private zone is true and the `weight_enabled` of the record_set is true. Default is 1.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The ttl of the private zone record. Unit: second. Default is 600.",
			},
			//"line": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Computed:    true,
			//	Description: "The subnet id of the private zone record. This field is only effected when the `intelligent_mode` of the private zone is true. Default is `Default`.",
			//},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The remark of the private zone record.",
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				Description: "Whether to enable the private zone record. This field is only effected when modify this resource.",
			},
		},
	}
	return resource
}

func resourceVolcenginePrivateZoneRecordCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneRecordService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcenginePrivateZoneRecord())
	if err != nil {
		return fmt.Errorf("error on creating private_zone_record %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneRecordRead(d, meta)
}

func resourceVolcenginePrivateZoneRecordRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneRecordService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcenginePrivateZoneRecord())
	if err != nil {
		return fmt.Errorf("error on reading private_zone_record %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcenginePrivateZoneRecordUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneRecordService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcenginePrivateZoneRecord())
	if err != nil {
		return fmt.Errorf("error on updating private_zone_record %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneRecordRead(d, meta)
}

func resourceVolcenginePrivateZoneRecordDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneRecordService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcenginePrivateZoneRecord())
	if err != nil {
		return fmt.Errorf("error on deleting private_zone_record %q, %s", d.Id(), err)
	}
	return err
}
