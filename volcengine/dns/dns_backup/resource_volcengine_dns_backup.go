package dns_backup

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
DnsBackup can be imported using the id, e.g.
```
$ terraform import volcengine_dns_backup.default ZID:BackupID
```

*/

func ResourceVolcengineDnsBackup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineDnsBackupCreate,
		Read:   resourceVolcengineDnsBackupRead,
		Delete: resourceVolcengineDnsBackupDelete,
		Importer: &schema.ResourceImporter{
			State: dnsBackupImporter,
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
				Description: "The ID of the domain for which you want to get the backup schedule.",
			},
			"backup_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of backup.",
			},
			"backup_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the backup was created. Timezone is UTC.",
			},
		},
	}
	return resource
}

func resourceVolcengineDnsBackupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineDnsBackup())
	if err != nil {
		return fmt.Errorf("error on creating dns_backup %q, %s", d.Id(), err)
	}
	return resourceVolcengineDnsBackupRead(d, meta)
}

func resourceVolcengineDnsBackupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineDnsBackup())
	if err != nil {
		return fmt.Errorf("error on reading dns_backup %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineDnsBackupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDnsBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineDnsBackup())
	if err != nil {
		return fmt.Errorf("error on deleting dns_backup %q, %s", d.Id(), err)
	}
	return err
}
