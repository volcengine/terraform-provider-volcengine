package rds_mssql_instance

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Rds Mssql Instance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mssql_instance.default resource_id
```

*/

func ResourceVolcengineRdsMssqlInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineSqlserverInstanceCreate,
		Read:   resourceVolcengineSqlserverInstanceRead,
		Update: resourceVolcengineSqlserverInstanceUpdate,
		Delete: resourceVolcengineSqlserverInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"db_engine_version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.",
			},
			"instance_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The Instance type. When the value of the `db_engine_version` is `SQLServer_2019_Std`, the value of this field can be `HA` or `Basic`." +
					"When the value of the `db_engine_version` is `SQLServer_2019_Ent`, the value of this field can be `Cluster` or `Basic`." +
					"When the value of the `db_engine_version` is `SQLServer_2019_Web`, the value of this field can be `Basic`.",
			},
			"subnet_id": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MinItems: 1,
				MaxItems: 2,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The subnet id of the instance node. When creating an instance that includes primary and backup nodes and needs to deploy primary and backup nodes across availability zones, you can specify two subnet_id. " +
					"By default, the first is the primary node availability zone, and the second is the backup node availability zone.",
			},
			"node_spec": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The node specification.",
			},
			"storage_space": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Storage space size, measured in GiB. The range of values is 20GiB to 4000GiB, with a step size of 10GiB.",
			},
			"super_account_password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				ForceNew:    true,
				Description: "The super account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Name of the instance.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name.",
			},
			"tags": ve.TagsSchema(),
			"charge_info": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
				Description: "The charge info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"charge_type": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The charge type. Valid values: `PostPaid`, `PrePaid`.",
						},
						"auto_renew": {
							Type:             schema.TypeBool,
							Optional:         true,
							ForceNew:         true,
							Computed:         true,
							DiffSuppressFunc: rdsMssqlInstanceDiffSuppress,
							Description:      "Whether to enable automatic renewal in the prepaid scenario. This parameter can be set when the ChargeType is `Prepaid`.",
						},
						"period": {
							Type:             schema.TypeInt,
							Optional:         true,
							ForceNew:         true,
							Computed:         true,
							DiffSuppressFunc: rdsMssqlInstanceDiffSuppress,
							Description:      "Purchase duration in a prepaid scenario. This parameter is required when the ChargeType is `Prepaid`.",
						},
						"charge_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Charge start time.",
						},
						"charge_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Charge end time.",
						},
						"charge_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge status.",
						},
						"overdue_reclaim_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Expected release time when overdue fees are shut down.",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time for Disconnection due to Unpaid Fees.",
						},
					},
				},
			},
			"backup_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The time window for starting the backup task is one hour interval. " +
					"\nThis field is valid and required when updating the backup plan of instance.",
			},
			"full_backup_period": {
				Type:     schema.TypeSet,
				Set:      schema.HashString,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Description: "Full backup cycle. Multiple values separated by commas. " +
					"The values are as follows: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday. " +
					"\nThis field is valid and required when updating the backup plan of instance.",
			},
			"backup_retention_period": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "Data backup retention days, value range: 7~30. " +
					"\nThis field is valid and required when updating the backup plan of instance.",
			},
		},
	}
	return resource
}

func resourceVolcengineSqlserverInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMssqlInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsMssqlInstance())
	if err != nil {
		return fmt.Errorf("error on creating rds_mssql_instance %q, %s", d.Id(), err)
	}
	return resourceVolcengineSqlserverInstanceRead(d, meta)
}

func resourceVolcengineSqlserverInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMssqlInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsMssqlInstance())
	if err != nil {
		return fmt.Errorf("error on reading rds_mssql_instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineSqlserverInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMssqlInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsMssqlInstance())
	if err != nil {
		return fmt.Errorf("error on updating rds_mssql_instance %q, %s", d.Id(), err)
	}
	return resourceVolcengineSqlserverInstanceRead(d, meta)
}

func resourceVolcengineSqlserverInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	return fmt.Errorf("This resource does not support deletion. " +
		"If you want to remove it from terraform state, " +
		"please use `terraform state rm volcengine_rds_mssql_instance.resource_name` command. ")
}

func rdsMssqlInstanceDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	//在计费方式为PostPaid的时候 period的变化会被忽略
	if d.Get("charge_info.0.charge_type").(string) == "PostPaid" && (k == "charge_info.0.period" || k == "charge_info.0.auto_renew") {
		return true
	}

	return false
}
