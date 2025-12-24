package rds_postgresql_backup_download

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlBackupDownload can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_backup_download.default resource_id
```

*/

func ResourceVolcengineRdsPostgresqlBackupDownload() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlBackupDownloadCreate,
		Read:   resourceVolcengineRdsPostgresqlBackupDownloadRead,
		Delete: resourceVolcengineRdsPostgresqlBackupDownloadDelete,
		Importer: &schema.ResourceImporter{State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
			parts := strings.Split(d.Id(), ":")
			if len(parts) < 2 {
				return []*schema.ResourceData{d}, fmt.Errorf("import id must be 'instance_id:backup_id'")
			}
			_ = d.Set("instance_id", parts[0])
			_ = d.Set("backup_id", parts[1])
			service := NewRdsPostgresqlBackupDownloadService(meta.(*ve.SdkClient))
			if err := service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlBackupDownload()); err != nil {
				return []*schema.ResourceData{d}, err
			}
			return []*schema.ResourceData{d}, nil
		}},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the PostgreSQL instance.",
			},
			"backup_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the logical backup to be downloaded.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlBackupDownloadCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlBackupDownloadService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlBackupDownload())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_backup_download %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlBackupDownloadRead(d, meta)
}

func resourceVolcengineRdsPostgresqlBackupDownloadRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlBackupDownloadService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlBackupDownload())
	if err != nil {
		return fmt.Errorf("error on reading rds_postgresql_backup_download %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlBackupDownloadDelete(d *schema.ResourceData, meta interface{}) (err error) {
	return nil
}
