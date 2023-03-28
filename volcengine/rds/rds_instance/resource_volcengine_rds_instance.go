package rds_instance

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RDS Instance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_instance.default mysql-42b38c769c4b
```

*/

func ResourceVolcengineRdsInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsInstanceCreate,
		Read:   resourceVolcengineRdsInstanceRead,
		Delete: resourceVolcengineRdsInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Set the name of the instance. The naming rules are as follows:\n\nCannot start with a number, a dash (-).\nIt can only contain Chinese characters, letters, numbers, underscores (_) and underscores (-).\nThe length needs to be within 1~128 characters.",
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The region of the RDS instance.",
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The available zone of the RDS instance.",
			},
			"db_engine": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  "Database type. Value:\nMySQL (default).",
				ValidateFunc: validation.StringInSlice([]string{"MySQL"}, false),
			},
			"db_engine_version": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Instance type. Value:\nMySQL_Community_5_7\nMySQL_8_0.",
				ValidateFunc: validation.StringInSlice([]string{"MySQL_Community_5_7", "MySQL_8_0"}, false),
			},
			"instance_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Instance type. Value:\nHA: High availability version.",
				ValidateFunc: validation.StringInSlice([]string{"HA"}, false),
			},
			"instance_spec_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance specification name, you can specify the specification name of the instance to be created. Value:\nrds.mysql.1c2g\nrds.mysql.2c4g\nrds.mysql.4c8g\nrds.mysql.4c16g\nrds.mysql.8c32g\nrds.mysql.16c64g\nrds.mysql.16c128g\nrds.mysql.32c128g\nrds.mysql.32c256g.",
			},
			"storage_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Instance storage type. Value:\nLocalSSD: Local SSD disk.",
				ValidateFunc: validation.StringInSlice([]string{"LocalSSD"}, false),
			},
			"storage_space_gb": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The storage space(GB) of the RDS instance.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The vpc ID of the RDS instance.",
			},
			"super_account_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Fill in the high-privileged user account name. The naming rules are as follows:\nUnique name.\nStart with a letter and end with a letter or number.\nConsists of lowercase letters, numbers, or underscores (_).\nThe length is 2~32 characters.\n[Keywords](https://www.volcengine.com/docs/6313/66162) are not allowed for account names.",
			},
			"supper_account_password": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Description:   "Set a high-privilege account password. The rules are as follows:\nOnly uppercase and lowercase letters, numbers and the following special characters _#!@$%^*()+=-.\nThe length needs to be within 8~32 characters.\nContains at least 3 of uppercase letters, lowercase letters, numbers or special characters.",
				ConflictsWith: []string{"super_account_password"},
				Deprecated:    "supper_account_password is deprecated, use super_account_password instead",
			},
			"super_account_password": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Description:   "Set a high-privilege account password. The rules are as follows:\nOnly uppercase and lowercase letters, numbers and the following special characters _#!@$%^*()+=-.\nThe length needs to be within 8~32 characters.\nContains at least 3 of uppercase letters, lowercase letters, numbers or special characters.",
				ConflictsWith: []string{"supper_account_password"},
			},
			"charge_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Billing type. Value:\nPostPaid: Postpaid (pay-as-you-go).\nPrepaid: Prepaid (yearly and monthly).",
				ValidateFunc: validation.StringInSlice([]string{"PostPaid", "Prepaid"}, false),
			},
			"auto_renew": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether to automatically renew. Default: false. Value:\ntrue: yes.\nfalse: no. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"prepaid_period": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The purchase cycle in the prepaid scenario. Value:\nMonth: monthly subscription.\nYear: yearly subscription. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"used_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The purchase time of RDS instance. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Select the project to which the instance belongs. If this parameter is left blank, the new instance will not be added to any project. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Subnet ID. The subnet must belong to the selected Availability Zone.",
			},
			"connection_info": {
				Type:        schema.TypeList,
				Computed:    true,
				MaxItems:    1,
				Description: "The connection info ot the RDS instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"internal_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The internal domain of the RDS instance.",
						},
						"internal_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The interval port of the RDS instance.",
						},
						"public_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public domain of the RDS instance.",
						},
						"public_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public port of the RDS instance.",
						},
						"enable_read_write_splitting": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether read-write separation is enabled.",
						},
						"enable_read_only": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether global read-only is enabled.",
						},
					},
				},
			},
		},
	}
}

func resourceVolcengineRdsInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsInstanceService := NewRdsInstanceService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Create(rdsInstanceService, d, ResourceVolcengineRdsInstance())
	if err != nil {
		return fmt.Errorf("error on creating RDS instance %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsInstanceRead(d, meta)
}

func resourceVolcengineRdsInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	rdsInstanceService := NewRdsInstanceService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Read(rdsInstanceService, d, ResourceVolcengineRdsInstance())
	if err != nil {
		return fmt.Errorf("error on reading RDS instance %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	rdsInstanceService := NewRdsInstanceService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Delete(rdsInstanceService, d, ResourceVolcengineRdsInstance())
	if err != nil {
		return fmt.Errorf("error on deleting RDS instance %q, %w", d.Id(), err)
	}
	return err
}
