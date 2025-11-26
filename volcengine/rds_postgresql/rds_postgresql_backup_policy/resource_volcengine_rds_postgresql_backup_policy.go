package rds_postgresql_backup_policy

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlBackupPolicy can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_backup_policy.default resource_id
```

*/

func ResourceVolcengineRdsPostgresqlBackupPolicy() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlBackupPolicyCreate,
		Read:   resourceVolcengineRdsPostgresqlBackupPolicyRead,
		Update: resourceVolcengineRdsPostgresqlBackupPolicyUpdate,
		Delete: resourceVolcengineRdsPostgresqlBackupPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the PostgreSQL instance.",
			},
			"backup_retention_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of days to retain backups, with a value range of 7 to 365.",
			},
			"full_backup_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The time when the backup task is executed. Format: HH:mmZ-HH:mmZ (UTC time).",
			},
			"full_backup_period": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Full backup period. Separate multiple values with an English comma (,)." +
					"Select at least one day per week for a full backup.",
			},
			"data_incr_backup_periods": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The incremental backup method follows the backup frequency for normal increments, with multiple values separated by English commas (,). " +
					"The selected values must not overlap with the full backup cycle. " +
					"Can select at most six days a week for incremental backup.",
			},
			"hourly_incr_backup_enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: "Whether to enable the high-frequency backup function. " +
					"To disable incremental backup, need to pass an empty string for the parameter data_incr_backup_periods and pass false for the parameter hourly_incr_backup_enable.",
			},
			"increment_backup_frequency": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "The method of incremental backup is the backup frequency for high-frequency increments. " +
					"The Unit: hours. The valid values are 1, 2, 4, 6, and 12.",
			},
			"wal_log_space_limit_enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: "Status of the local remaining available space protection function. " +
					"When enabled, it will automatically start clearing the earliest local WAL logs when the total storage space usage rate of the instance exceeds 80% " +
					"or the remaining space is less than 5GB, until the total space usage rate is below 80% and the remaining space is greater than 5GB.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlBackupPolicyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlBackupPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlBackupPolicy())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_backup_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlBackupPolicyRead(d, meta)
}

func resourceVolcengineRdsPostgresqlBackupPolicyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlBackupPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlBackupPolicy())
	if err != nil {
		return fmt.Errorf("error on reading rds_postgresql_backup_policy %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlBackupPolicyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlBackupPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsPostgresqlBackupPolicy())
	if err != nil {
		return fmt.Errorf("error on updating rds_postgresql_backup_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlBackupPolicyRead(d, meta)
}

func resourceVolcengineRdsPostgresqlBackupPolicyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlBackupPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlBackupPolicy())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_backup_policy %q, %s", d.Id(), err)
	}
	return err
}
