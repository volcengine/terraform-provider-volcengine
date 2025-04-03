package rds_mysql_backup_policy

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsMysqlBackupPolicy can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_backup_policy.default instanceId:backupPolicy
```
Warning:The resource cannot be deleted, and the destroy operation will not perform any actions.
*/

func ResourceVolcengineRdsMysqlBackupPolicy() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsMysqlBackupPolicyCreate,
		Read:   resourceVolcengineRdsMysqlBackupPolicyRead,
		Update: resourceVolcengineRdsMysqlBackupPolicyUpdate,
		Delete: resourceVolcengineRdsMysqlBackupPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("instance_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "The ID of the RDS instance.",
			},
			"data_full_backup_periods": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Description: "Full backup period. It is recommended to select at least 2 days for full backup every week." +
					" Multiple values are separated by English commas (,). " +
					"Values: Monday. Tuesday. Wednesday. Thursday. Friday. Saturday. Sunday. " +
					"When modifying the data backup policy, this parameter needs to be passed in.",
			},
			"data_backup_retention_day": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Data backup retention days, value range: 7 to 365 days. Default retention is 7 days.",
			},
			"data_full_backup_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Time window for executing backup tasks is one hour. " +
					"Format: HH:mmZ-HH:mmZ (UTC time). " +
					"Explanation: This parameter needs to be passed in when modifying the data backup policy.",
			},
			"data_incr_backup_periods": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Description: "Incremental backup period." +
					" Multiple values are separated by commas (,). " +
					"Values: Monday. Tuesday. Wednesday. Thursday. Friday. Saturday. Sunday." +
					"Description: The incremental backup period cannot conflict with the full backup." +
					" When modifying the data backup policy, this parameter needs to be passed in.",
			},
			"binlog_file_counts_enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				Description: "Whether to enable the upper limit of local Binlog retention. " +
					"Values: true: Enabled. false: Disabled. " +
					"Description:When modifying the log backup policy, " +
					"this parameter needs to be passed in.",
			},
			"binlog_limit_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Description: "Number of local Binlog retained, ranging from 6 to 1000, " +
					"in units of pieces. Automatically delete local logs that exceed the retained number. " +
					"Explanation: When modifying the log backup policy, this parameter needs to be passed in.",
			},
			"binlog_local_retention_hour": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Description: "Local Binlog retention duration, " +
					"with a value ranging from 0 to 168, in hours. " +
					"Local logs exceeding the retention duration will be automatically deleted. " +
					"When set to 0, local logs will not be automatically deleted. " +
					"Note: When modifying the log backup policy, this parameter needs to be passed.",
			},
			"binlog_space_limit_enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				Description: "Whether to enable automatic cleanup of Binlog when space is too large." +
					" When the total storage space occupancy rate of the instance exceeds 80% or the remaining space is less than 5GB, " +
					"the system will automatically start cleaning up the earliest local Binlog until the " +
					"total space occupancy rate is lower than 80% and the remaining space is greater than 5GB. " +
					"true: Enabled. false: Disabled. " +
					"Description: This parameter needs to be passed in when modifying the log backup policy.",
			},
			"binlog_storage_percentage": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Description: "Maximum storage space usage rate can be set to 20% - 50%. " +
					"After exceeding this limit, the earliest Binlog file will be automatically " +
					"deleted until the space usage rate is lower than this ratio. " +
					"Local Binlog space usage rate = Local Binlog size / Total available (purchased) instance space size. " +
					"When modifying the log backup policy, this parameter needs to be passed in. " +
					"Explanation: When modifying the log backup policy, this parameter needs to be passed in.",
			},
			"log_backup_retention_day": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("retention_policy_synced").(bool) {
						return true
					}
					return false
				},
				Description: "Binlog backup retention period. The value range is 7 to 365, in days. " +
					"Explanation: When modifying the log backup policy, this parameter needs to be passed in.",
			},
			"lock_ddl_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Description: "Maximum waiting time for DDL. The default value is 30. " +
					"The minimum value is 10. The maximum value is 1440. The unit is minutes. " +
					"Description: Only instances of MySQL 8.0 version support this setting.",
			},
			"data_full_backup_start_utc_hour": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Description: "The start point (UTC time) of the time window for starting the full backup task." +
					" The time window length is 1 hour. " +
					"Explanation: Both DataFullBackupStartUTCHour and DataFullBackupTime can be used to indicate the full backup time period of an instance." +
					" DataFullBackupStartUTCHour has higher priority. " +
					"If both fields are returned at the same time, DataFullBackupStartUTCHour shall prevail.",
			},
			"hourly_incr_backup_enable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable high-frequency backup function. Values:\ntrue: Yes.\nfalse: No.",
			},
			"incr_backup_hour_period": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("hourly_incr_backup_enable").(bool) {
						return false
					}
					return true
				},
				Description: "Frequency of performing high-frequency incremental backups. " +
					"Values: 2: Perform an incremental backup every 2 hours. " +
					"4: Perform an incremental backup every 4 hours. " +
					"6: Perform an incremental backup every 6 hours. " +
					"12: Perform an incremental backup every 12 hours. " +
					"Description: This parameter takes effect only when HourlyIncrBackupEnable is set to true.",
			},
			"data_backup_encryption_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable encryption for data backup. Values:\ntrue: Yes.\nfalse: No.",
			},
			"binlog_backup_encryption_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Is encryption enabled for log backups? Values:\ntrue: Yes.\nfalse: No.",
			},
			"data_keep_policy_after_released": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Policy for retaining a backup of an instance after it is released." +
					" The values are: Last: Keep the last backup. " +
					"Default value. All: Keep all backups of the instance.",
			},
			"data_keep_days_after_released": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Backup retention days when an instance is released. Currently, only a value of 7 is supported.",
			},
			"data_backup_all_retention": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to retain all data backups before releasing the instance. Values:\ntrue: Yes.\nfalse: No.",
			},
			"binlog_backup_all_retention": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("retention_policy_synced").(bool) {
						return true
					}
					return false
				},
				Description: "Whether to retain all log backups before releasing an instance." +
					" Values:\ntrue: Yes.\nfalse: No. " +
					"Description: BinlogBackupAllRetention is ineffective when the value of RetentionPolicySynced is true.",
			},
			"binlog_backup_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable log backup function. Values:\ntrue: Yes.\nfalse: No.",
			},
			"retention_policy_synced": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				Description: "Is the retention policy for log backups the same as that for data backups?\n" +
					"Explanation: When the value is true, LogBackupRetentionDay and BinlogBackupAllRetention are ignored.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsMysqlBackupPolicyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlBackupPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsMysqlBackupPolicy())
	if err != nil {
		return fmt.Errorf("error on creating rds_mysql_backup_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlBackupPolicyRead(d, meta)
}

func resourceVolcengineRdsMysqlBackupPolicyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlBackupPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsMysqlBackupPolicy())
	if err != nil {
		return fmt.Errorf("error on reading rds_mysql_backup_policy %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsMysqlBackupPolicyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlBackupPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsMysqlBackupPolicy())
	if err != nil {
		return fmt.Errorf("error on updating rds_mysql_backup_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlBackupPolicyRead(d, meta)
}

func resourceVolcengineRdsMysqlBackupPolicyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlBackupPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsMysqlBackupPolicy())
	if err != nil {
		return fmt.Errorf("error on deleting rds_mysql_backup_policy %q, %s", d.Id(), err)
	}
	return err
}
