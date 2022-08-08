package rds_instance_v2

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RDS Instance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_instance_v2.default mysql-42b38c769c4b
```

*/

func ResourceVolcengineRdsInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsInstanceCreate,
		Read:   resourceVolcengineRdsInstanceRead,
		Update: resourceVolcengineRdsInstanceUpdate,
		Delete: resourceVolcengineRdsInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"db_engine_version": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Instance type. Value:\nMySQL_5_7\nMySQL_8_0.",
				ValidateFunc: validation.StringInSlice([]string{"MySQL_5_7", "MySQL_8_0"}, false),
			},
			"instance_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Instance type. Value:\nHA: High availability version.",
				ValidateFunc: validation.StringInSlice([]string{"HA"}, false),
			},
			"storage_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance storage type. When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, the value is:\nLocalSSD - local SSD disk\nWhen the database type is veDB_MySQL/veDB_PostgreSQL, the value is:\nDistributedStorage - Distributed Storage.",
			},
			"storage_space": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Instance storage space.\nWhen the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, value range: [20, 3000], unit: GB, increments every 100GB.\nWhen the database type is veDB_MySQL/veDB_PostgreSQL, this parameter does not need to be passed.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Private network (VPC) ID. You can call the DescribeVpcs query and use this parameter to specify the VPC where the instance is to be created.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Subnet ID.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
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
				ForceNew:    true,
				Description: "Time zone. Support UTC -12:00 ~ +13:00.",
			},
			"db_param_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Parameter template ID. It only takes effect when the database type is MySQL/PostgreSQL/SQL_Server.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Subordinate to the project.",
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
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Payment type. Value:\nPostPaid - Pay-As-You-Go\nPrePaid - Yearly and monthly (default).",
						},
						"auto_renew": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "Whether to automatically renew in prepaid scenarios.\nAutorenew_Enable\nAutorenew_Disable (default).",
						},
						"period_unit": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The purchase cycle in the prepaid scenario.\nMonth - monthly subscription (default)\nYear - Package year.",
						},
						"period": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "Purchase duration in prepaid scenarios. Default: 1.",
						},
					},
				},
			},
			"node_info": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Instance specification configuration. This parameter is required for RDS for MySQL, RDS for PostgreSQL and MySQL Sharding. There is one and only one Primary node, one and only one Secondary node, and 0-10 Read-Only nodes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Node ID.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Zone ID.",
						},
						"node_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Node type, the value is \"Primary\", \"Secondary\", \"ReadOnly\".",
						},
						"node_spec": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Masternode specs. Pass\nDescribeDBInstanceSpecs Query the instance specifications that can be sold.",
						},
					},
				},
			},
		},
	}
}

func resourceVolcengineRdsInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsInstanceService := NewRdsInstanceService(meta.(*volc.SdkClient))
	err = rdsInstanceService.Dispatcher.Create(rdsInstanceService, d, ResourceVolcengineRdsInstance())
	if err != nil {
		return fmt.Errorf("error on creating RDS instance %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsInstanceRead(d, meta)
}

func resourceVolcengineRdsInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	rdsInstanceService := NewRdsInstanceService(meta.(*volc.SdkClient))
	err = rdsInstanceService.Dispatcher.Read(rdsInstanceService, d, ResourceVolcengineRdsInstance())
	if err != nil {
		return fmt.Errorf("error on reading RDS instance %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsInstanceService := NewRdsInstanceService(meta.(*volc.SdkClient))
	err = rdsInstanceService.Dispatcher.Update(rdsInstanceService, d, ResourceVolcengineRdsInstance())
	if err != nil {
		return fmt.Errorf("error on updating RDS instance %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsInstanceRead(d, meta)
}

func resourceVolcengineRdsInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	rdsInstanceService := NewRdsInstanceService(meta.(*volc.SdkClient))
	err = rdsInstanceService.Dispatcher.Delete(rdsInstanceService, d, ResourceVolcengineRdsInstance())
	if err != nil {
		return fmt.Errorf("error on deleting RDS instance %q, %w", d.Id(), err)
	}
	return err
}
