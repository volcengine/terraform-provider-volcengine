package sqlserver_instance

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
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
				Description: "Compatible version. Currently only supports the value SQLServer_2019_Std, which represents SQL Server 2019 Standard Edition.",
			},
			"instance_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance type. Currently only supports the value HA, which represents high availability type.",
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
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet id.",
			},
			"super_account_password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				ForceNew:    true,
				Description: "The super account password.",
			},
			"db_time_zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Time zone. Currently only supports the value of China Standard Time.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Name of the instance.",
			},
			"server_collation": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "SQL Server sorting rules, currently only support the Chinese_PRC_CI_AS (default) value.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name.",
			},
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
							Description: "The charge type.",
						},
						"auto_renew": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "Whether to enable automatic renewal in the prepaid scenario. This parameter can be set when ChargeType is Prepaid.",
						},
						"period_unit": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "Purchase cycle in prepaid scenarios. This parameter can be set when ChargeType is Prepaid.\nMonth: represents monthly (default).\nYear: represents yearly.",
						},
						"period": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "Purchase duration in a prepaid scenario.",
						},
						"number": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Description: "Example purchase quantity. Default value: 1, range of values is [1,10].",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Tags.",
				Set:         tagsHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The Key of Tags.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The Value of Tags.",
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
		"please use `terraform state rm volcengine_sqlserver_instance.resource_id` command. ")
}

var tagsHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v#%v", m["key"], m["value"]))
	return hashcode.String(buf.String())
}
