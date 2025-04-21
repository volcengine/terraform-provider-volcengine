package dns_backup_schedule

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
DnsBackupSchedule can be imported using the id, e.g.
```
$ terraform import volcengine_dns_backup_schedule.default resource_id
```

*/

func ResourceVolcengineDnsBackupSchedule() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineDnsBackupScheduleCreate,
		Read:   resourceVolcengineDnsBackupScheduleRead,
		Update: resourceVolcengineDnsBackupScheduleUpdate,
		Delete: resourceVolcengineDnsBackupScheduleDelete,
		Importer: &schema.ResourceImporter{
			State: dnsBackupScheduleImporter,
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
				Description: "The ID of the domain for which you want to update the backup schedule.",
			},
			"schedule": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntNotInSlice([]int{0}),
				Description: "The backup schedule. 0: Turn off automatic backup. " +
					"1: Automatic backup once per hour. " +
					"2: Automatic backup once per day. " +
					"3: Automatic backup once per month.",
			},
			"count_limit": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Maximum number of backups per domain.",
			},
		},
	}
	return resource
}

func resourceVolcengineDnsBackupScheduleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsBackupScheduleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineDnsBackupSchedule())
	if err != nil {
		return fmt.Errorf("error on creating dns_backup_schedule %q, %s", d.Id(), err)
	}
	return resourceVolcengineDnsBackupScheduleRead(d, meta)
}

func resourceVolcengineDnsBackupScheduleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsBackupScheduleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineDnsBackupSchedule())
	if err != nil {
		return fmt.Errorf("error on reading dns_backup_schedule %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineDnsBackupScheduleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsBackupScheduleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineDnsBackupSchedule())
	if err != nil {
		return fmt.Errorf("error on updating dns_backup_schedule %q, %s", d.Id(), err)
	}
	return resourceVolcengineDnsBackupScheduleRead(d, meta)
}

func resourceVolcengineDnsBackupScheduleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsBackupScheduleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineDnsBackupSchedule())
	if err != nil {
		return fmt.Errorf("error on deleting dns_backup_schedule %q, %s", d.Id(), err)
	}
	return err
}
