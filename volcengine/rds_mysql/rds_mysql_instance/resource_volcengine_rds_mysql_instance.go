package rds_mysql_instance

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Rds Mysql Instance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_instance.default mysql-72da4258c2c7
```

*/

func ResourceVolcengineRdsMysqlInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsMysqlInstanceCreate,
		Read:   resourceVolcengineRdsMysqlInstanceRead,
		Update: resourceVolcengineRdsMysqlInstanceUpdate,
		Delete: resourceVolcengineRdsMysqlInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"db_engine_version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance type. Value:\nMySQL_5_7\nMySQL_8_0.",
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
				ForceNew:    true,
				Description: "The available zone of secondary node.",
			},
			"storage_space": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     100,
				Description: "Instance storage space. Value range: [20, 3000], unit: GB, increments every 100GB. Default value: 100.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Subnet ID of the RDS instance.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance name. Cannot start with a number or a dash\nCan only contain Chinese characters, letters, numbers, underscores and dashes\nThe length is limited between 1 ~ 128.",
			},
			"lower_case_table_names": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether the table name is case sensitive, the default value is 1.\nRanges:\n0: Table names are stored as fixed and table names are case-sensitive.\n1: Table names will be stored in lowercase and table names are not case sensitive.",
			},
			"db_time_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the RDS instance.",
			},
			"tags": ve.TagsSchema(),
			"connection_pool_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Connection pool type. Value range:\nDirect: Direct connection mode.\nTransaction: Transaction-level connection pool (default).",
			},
			"charge_info": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
				Description: "Payment methods.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"charge_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							ValidateFunc: validation.StringInSlice([]string{
								"PostPaid",
								"PrePaid",
							}, false),
							Description: "Payment type. Value:\nPostPaid - Pay-As-You-Go\nPrePaid - Yearly and monthly (default).",
						},
						"auto_renew": {
							Type:             schema.TypeBool,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: RdsMysqlInstanceImportDiffSuppress,
							Description:      "Whether to automatically renew in prepaid scenarios.",
						},
						"period_unit": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     validation.StringInSlice([]string{"Month", "Year"}, false),
							DiffSuppressFunc: RdsMysqlInstanceImportDiffSuppress,
							Description:      "The purchase cycle in the prepaid scenario.\nMonth - monthly subscription (default)\nYear - Package year.",
						},
						"period": {
							Type:             schema.TypeInt,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: RdsMysqlInstanceImportDiffSuppress,
							Description:      "Purchase duration in prepaid scenarios. Default: 1.",
						},
					},
				},
			},
			"allow_list_ids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Set:         schema.HashString,
				Description: "Allow list Ids of the RDS instance.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"parameters": {
				Type:        schema.TypeSet,
				Optional:    true,
				Set:         parameterHash,
				Description: "Parameter of the RDS instance. This field can only be added or modified. Deleting this field is invalid.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"parameter_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Parameter name.",
						},
						"parameter_value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Parameter value.",
						},
					},
				},
			},
			"maintenance_window": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Description: "Specify the maintainable time period of the instance when creating the instance. " +
					"This field is optional. If not set, " +
					"it defaults to 18:00Z - 21:59Z of every day within a week (that is, 02:00 - 05:59 Beijing time).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"maintenance_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Maintenance period of an instance. Format: HH:mmZ-HH:mmZ (UTC time).",
						},
						"day_kind": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Maintenance cycle granularity, values: Week: Week. Month: Month.",
						},
						"day_of_week": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Specify the maintainable time period of a certain day of the week." +
								" The values are: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday." +
								" Multiple selections are allowed. If this value is not specified or is empty, " +
								"it defaults to specifying all seven days of the week.",
						},
					},
				},
			},
		},
	}
	dataSource := DataSourceVolcengineRdsMysqlInstances().Schema["rds_mysql_instances"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineRdsMysqlInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsMysqlInstanceService := NewRdsMysqlInstanceService(meta.(*ve.SdkClient))
	err = rdsMysqlInstanceService.Dispatcher.Create(rdsMysqlInstanceService, d, ResourceVolcengineRdsMysqlInstance())
	if err != nil {
		return fmt.Errorf("error on creating RDS mysql instance %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlInstanceRead(d, meta)
}

func resourceVolcengineRdsMysqlInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	rdsMysqlInstanceService := NewRdsMysqlInstanceService(meta.(*ve.SdkClient))
	err = rdsMysqlInstanceService.Dispatcher.Read(rdsMysqlInstanceService, d, ResourceVolcengineRdsMysqlInstance())
	if err != nil {
		return fmt.Errorf("error on reading RDS mysql instance %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsMysqlInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsMysqlInstanceService := NewRdsMysqlInstanceService(meta.(*ve.SdkClient))
	err = rdsMysqlInstanceService.Dispatcher.Update(rdsMysqlInstanceService, d, ResourceVolcengineRdsMysqlInstance())
	if err != nil {
		return fmt.Errorf("error on updating RDS mysql instance %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlInstanceRead(d, meta)
}

func resourceVolcengineRdsMysqlInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	rdsMysqlInstanceService := NewRdsMysqlInstanceService(meta.(*ve.SdkClient))
	err = rdsMysqlInstanceService.Dispatcher.Delete(rdsMysqlInstanceService, d, ResourceVolcengineRdsMysqlInstance())
	if err != nil {
		return fmt.Errorf("error on deleting RDS mysql instance %q, %w", d.Id(), err)
	}
	return err
}
