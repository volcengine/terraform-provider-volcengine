package rds_postgresql_instance

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlInstance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_instance.default postgres-21a3333b****
```

*/

func ResourceVolcengineRdsPostgresqlInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlInstanceCreate,
		Read:   resourceVolcengineRdsPostgresqlInstanceRead,
		Update: resourceVolcengineRdsPostgresqlInstanceUpdate,
		Delete: resourceVolcengineRdsPostgresqlInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		// NodeInfo 和 VpcId 在 create 中填充
		Schema: map[string]*schema.Schema{
			"db_engine_version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance type. Value: PostgreSQL_11, PostgreSQL_12, PostgreSQL_13, PostgreSQL_14, PostgreSQL_15, PostgreSQL_16, PostgreSQL_17.",
			},
			"node_spec": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The specification of primary node and secondary node.",
			},
			"primary_zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The available zone of primary node.",
			},
			"secondary_zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The available zone of secondary node.",
			},
			"storage_space": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     100,
				Description: "Instance storage space. Value range: [20, 3000], unit: GB, step 10GB. Default value: 100.",
			},
			"modify_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Usually",
				Description: "Spec change type. Usually(default) or Temporary.",
			},
			"rollback_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rollback time for Temporary change, UTC format yyyy-MM-ddTHH:mm:ss.sssZ.",
			},
			"estimate_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				Description: "Whether to initiate a configuration change assessment. Only estimate spec change impact without executing. " +
					"Default value: false.",
			},
			"estimation_result": { // readsource 空值
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The estimated impact on the instance after the current configuration changes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"plans": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Estimated impact on the instance after the current configuration changes.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"effects": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "After changing according to the current configuration, the estimated impact on the read and write connections of the instance.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Subnet ID of the RDS PostgreSQL instance.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance name. Cannot start with a number or a dash. Can only contain Chinese characters, letters, numbers, underscores and dashes. The length is limited between 1 ~ 128.",
			},
			"tags": ve.TagsSchema(),
			"allow_list_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Allow list IDs to bind at creation.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the RDS instance.",
			},
			"charge_info": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				// ForceNew:    true,
				Description: "Payment methods.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"charge_type": {
							Type:     schema.TypeString,
							Required: true,
							Description: "Payment type. Value:\nPostPaid - Pay-As-You-Go\nPrePaid - Yearly and monthly (default). \n" +
								"When the value of this field is `PrePaid`, the postgresql instance cannot be deleted through terraform. Please unsubscribe the instance from the Volcengine console first, and then use `terraform state rm volcengine_rds_postgresql_instance.resource_name` command to remove it from terraform state file and management.",
						},
						"auto_renew": {
							Type:             schema.TypeBool,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: RdsPostgreSQLInstanceImportDiffSuppress,
							Description:      "Whether to automatically renew in prepaid scenarios.",
						},
						"period_unit": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: RdsPostgreSQLInstanceImportDiffSuppress,
							Description:      "The purchase cycle in the prepaid scenario.\nMonth - monthly subscription (default)\nYear - Package year.",
						},
						"period": {
							Type:             schema.TypeInt,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: RdsPostgreSQLInstanceImportDiffSuppress,
							Description:      "Purchase duration in prepaid scenarios. Default: 1.",
						},
						"number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1,
							Description: "Purchase number of the RDS PostgreSQL instance. Range: [1, 20]. Default: 1.",
						},
					},
				},
			},
			"parameters": {
				Type:        schema.TypeSet,
				Optional:    true,
				Set:         parameterHash,
				Description: "Parameter of the RDS PostgreSQL instance. This field can only be added or modified. Deleting this field is invalid.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Parameter name.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Parameter value.",
						},
					},
				},
			},
			"zone_migrations": {
				Type:     schema.TypeList,
				Optional: true,
				Description: "Nodes to migrate AZ. Only Secondary or ReadOnly nodes are allowed. " +
					"If you want to migrate the availability zone of the secondary node, you need to add the zone_migrations field. Modifying the secondary_zone_id directly will not work. " +
					"Cross-AZ instance migration is not supported.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Node ID to migrate.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Target zone ID.",
						},
						"node_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Node type: Secondary or ReadOnly.",
						},
					},
				},
			},
			"src_instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Source instance ID. After setting it, a new instance will be created by restoring from the backup/time point.",
			},
			"backup_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Backup ID (choose either this or restore_time; if both are set, backup_id shall prevail).",
			},
			"restore_time": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The point in time to restore to, in UTC format yyyy-MM-ddTHH:mm:ssZ (choose either this or backup_id).",
			},
		},
	}
	dataSource := DataSourceVolcengineRdsPostgresqlInstances().Schema["instances"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineRdsPostgresqlInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlInstance())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_instance %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlInstanceRead(d, meta)
}

func resourceVolcengineRdsPostgresqlInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlInstance())
	if err != nil {
		return fmt.Errorf("error on reading rds_postgresql_instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsPostgresqlInstance())
	if err != nil {
		return fmt.Errorf("error on updating rds_postgresql_instance %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlInstanceRead(d, meta)
}

func resourceVolcengineRdsPostgresqlInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlInstance())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_instance %q, %s", d.Id(), err)
	}
	return err
}

func RdsPostgreSQLInstanceImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	//在计费方式为PostPaid的时候 period的变化会被忽略
	if d.Get("charge_info.0.charge_type").(string) == "PostPaid" && (k == "charge_info.0.period" || k == "charge_info.0.period_unit" || k == "charge_info.0.auto_renew") {
		return true
	}

	return false
}

var parameterHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v#%v", m["name"], m["value"]))
	return hashcode.String(buf.String())
}
