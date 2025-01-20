package rds_mysql_backup

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsMysqlBackup can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_backup.default instanceId:backupId
```

*/

func ResourceVolcengineRdsMysqlBackup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsMysqlBackupCreate,
		Read:   resourceVolcengineRdsMysqlBackupRead,
		Delete: resourceVolcengineRdsMysqlBackupDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("instance_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("backup_id", items[1]); err != nil {
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
				ForceNew:    true,
				Description: "The id of the instance.",
			},
			"backup_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the backup.",
			},
			"backup_type": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Backup type. Currently, only full backup is supported. The value is Full.",
			},
			"backup_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Description: "Backup method. Value range: Full, " +
					"full backup under physical backup type. Default value. " +
					"DumpAll: full database backup under logical backup type. " +
					"Prerequisite: If you need to create a full database backup of logical backup type, " +
					"that is, when the value of BackupType is DumpAll, " +
					"the backup type should be set to logical backup, " +
					"that is, the value of BackupMethod should be Logical. " +
					"If you need to create a database table backup of logical backup type, " +
					"you do not need to pass in this field. " +
					"You only need to specify the database and table to be backed up in the BackupMeta field.",
			},
			"backup_meta": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Description: "When creating a library table backup of logical backup type, " +
					"it is used to specify the library table information to be backed up.\n" +
					"Prerequisite: When the value of BackupMethod is Logical, and the BackupType field is not passed." +
					"\nMutual exclusion situation: When the value of the BackupType field is DumpAll, " +
					"this field is not effective.\nQuantity limit: When creating a specified library table backup, " +
					"the upper limit of the number of libraries is 5000, " +
					"and the upper limit of the number of tables in each library is 5000.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("backup_type").(string) == "DumpAll" {
						return true
					}
					if d.Get("backup_method").(string) == "Logical" {
						_, ok := d.GetOk("backup_type")
						if !ok {
							return false
						}
					}
					return true
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Specify the database that needs to be backed up.",
						},
						"table_names": {
							Type:     schema.TypeSet,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Specify the tables to be backed up in the specified database." +
								" When this field is empty, it defaults to full database backup.",
						},
					},
				},
			},
			//"backup_policy": {
			//	Type:        schema.TypeList,
			//	MaxItems:    1,
			//	Optional:    true,
			//	Computed:    true,
			//	Description: "Data backup strategy for instances.",
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			"data_full_backup_periods": {
			//				Type: schema.TypeSet,
			//				Elem: &schema.Schema{
			//					Type: schema.TypeString,
			//				},
			//				Optional: true,
			//				Description: "Full backup period. It is recommended to select at least 2 days for full backup every week." +
			//					" Multiple values are separated by English commas (,). " +
			//					"Values: Monday. Tuesday. Wednesday. Thursday. Friday. Saturday. Sunday. " +
			//					"When modifying the data backup policy, this parameter needs to be passed in.",
			//			},
			//			"data_backup_retention_day": {
			//				Type:        schema.TypeInt,
			//				Optional:    true,
			//				Description: "Data backup retention days, value range: 7 to 365 days. Default retention is 7 days.",
			//			},
			//			"data_full_backup_time": {
			//				Type:     schema.TypeString,
			//				Optional: true,
			//				Description: "Time window for executing backup tasks is one hour. " +
			//					"Format: HH:mmZ-HH:mmZ (UTC time). " +
			//					"Explanation: This parameter needs to be passed in when modifying the data backup policy.",
			//			},
			//			"data_incr_backup_periods": {
			//				Type: schema.TypeSet,
			//				Elem: &schema.Schema{
			//					Type: schema.TypeString,
			//				},
			//				Optional: true,
			//				Description: "Incremental backup period." +
			//					" Multiple values are separated by commas (,). " +
			//					"Values: Monday. Tuesday. Wednesday. Thursday. Friday. Saturday. Sunday." +
			//					"Description: The incremental backup period cannot conflict with the full backup." +
			//					" When modifying the data backup policy, this parameter needs to be passed in.",
			//			},
			//			"binlog_file_counts_enable": {
			//				Type:     schema.TypeBool,
			//				Optional: true,
			//				Description: "Whether to enable the upper limit of local Binlog retention. " +
			//					"Values: true: Enabled. false: Disabled. " +
			//					"Description:When modifying the log backup policy, " +
			//					"this parameter needs to be passed in.",
			//			},
			//			"binlog_limit_count": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Description: "Number of local Binlog retained, ranging from 6 to 1000, " +
			//					"in units of pieces. Automatically delete local logs that exceed the retained number. " +
			//					"Explanation: When modifying the log backup policy, this parameter needs to be passed in.",
			//			},
			//			"binlog_local_retention_hour": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Description: "Local Binlog retention duration, " +
			//					"with a value ranging from 0 to 168, in hours. " +
			//					"Local logs exceeding the retention duration will be automatically deleted. " +
			//					"When set to 0, local logs will not be automatically deleted. " +
			//					"Note: When modifying the log backup policy, this parameter needs to be passed.",
			//			},
			//			"binlog_space_limit_enable": {
			//				Type:     schema.TypeBool,
			//				Optional: true,
			//				Description: "Whether to enable automatic cleanup of Binlog when space is too large." +
			//					" When the total storage space occupancy rate of the instance exceeds 80% or the remaining space is less than 5GB, " +
			//					"the system will automatically start cleaning up the earliest local Binlog until the " +
			//					"total space occupancy rate is lower than 80% and the remaining space is greater than 5GB. " +
			//					"true: Enabled. false: Disabled. " +
			//					"Description: This parameter needs to be passed in when modifying the log backup policy.",
			//			},
			//			"binlog_storage_percentage": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Description: "Maximum storage space usage rate can be set to 20% - 50%. " +
			//					"After exceeding this limit, the earliest Binlog file will be automatically " +
			//					"deleted until the space usage rate is lower than this ratio. " +
			//					"Local Binlog space usage rate = Local Binlog size / Total available (purchased) instance space size. " +
			//					"When modifying the log backup policy, this parameter needs to be passed in. " +
			//					"Explanation: When modifying the log backup policy, this parameter needs to be passed in.",
			//			},
			//			"log_backup_retention_day": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Description: "Binlog backup retention period. The value range is 7 to 365, in days. " +
			//					"Explanation: When modifying the log backup policy, this parameter needs to be passed in.",
			//			},
			//			"lock_ddl_time": {
			//				Type:     schema.TypeInt,
			//				Optional: true,
			//				Description: "Maximum waiting time for DDL. The default value is 30. " +
			//					"The minimum value is 10. The maximum value is 1440. The unit is minutes. " +
			//					"Description: Only instances of MySQL 8.0 version support this setting.",
			//			},
			//		},
			//	},
			//},
		},
	}
	return resource
}

func resourceVolcengineRdsMysqlBackupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsMysqlBackup())
	if err != nil {
		return fmt.Errorf("error on creating rds_mysql_backup %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlBackupRead(d, meta)
}

func resourceVolcengineRdsMysqlBackupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsMysqlBackup())
	if err != nil {
		return fmt.Errorf("error on reading rds_mysql_backup %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsMysqlBackupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsMysqlBackup())
	if err != nil {
		return fmt.Errorf("error on deleting rds_mysql_backup %q, %s", d.Id(), err)
	}
	return err
}
