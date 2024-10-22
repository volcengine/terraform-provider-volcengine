package vedb_mysql_instance

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VedbMysqlInstance can be imported using the id, e.g.
```
$ terraform import volcengine_vedb_mysql_instance.default resource_id
```

*/

func ResourceVolcengineVedbMysqlInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVedbMysqlInstanceCreate,
		Read:   resourceVolcengineVedbMysqlInstanceRead,
		Update: resourceVolcengineVedbMysqlInstanceUpdate,
		Delete: resourceVolcengineVedbMysqlInstanceDelete,
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
				Description: "Database engine version, with a fixed value of MySQL_8_0.",
			},
			"db_minor_version": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "veDB MySQL minor version. For detailed instructions on version numbers, please refer to Version Number Management.\n " +
					"3.0 (default): veDB MySQL stable version, 100% compatible with MySQL 8.0.\n " +
					"3.1: Natively supports HTAP application scenarios and accelerates complex queries.\n " +
					"3.2: Natively supports HTAP application scenarios and accelerates complex queries. " +
					"In addition, it has built-in cold data archiving capabilities. " +
					"It can archive data with low-frequency access to object storage TOS to reduce storage costs.",
			},
			"node_spec": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Node specification code of an instance.",
			},
			"node_number": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Number of instance nodes. The value range is from 2 to 16.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Subnet ID of the veDB Mysql instance.",
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				// 2024-10-16 当前只支持通过控制台修改端口号，api不支持修改
				ForceNew: true,
				Description: "Specify the private network port number for the connection terminal created by default for the instance." +
					" The default value is 3306, and the value range is 1000 to 65534.\n" +
					"Note:\nThis configuration item is only effective for the primary node terminal, default terminal, " +
					"and HTAP cluster terminal. That is, after the instance is created successfully," +
					" for the newly created custom terminal, the port number is still 3306 by default.\n" +
					"After the instance is created successfully, you can also modify the port number at any time. " +
					"Currently, only modification through the console is supported.",
			},
			// 超级账号功能放到account里，这里不暴露账号相关字段
			//"super_account_name": {
			//
			//},
			"charge_type": {
				Type:     schema.TypeString,
				Required: true,
				Description: "Calculate the billing type. When calculating the billing type during instance creation, " +
					"the possible values are as follows:\n" +
					"PostPaid: Pay-as-you-go (postpaid).\n" +
					"PrePaid: Monthly or yearly subscription (prepaid).",
			},
			"storage_charge_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Storage billing type. " +
					"When this parameter is not passed, the storage billing type defaults to be the same as the computing billing type. " +
					"The values are as follows:\n" +
					"PostPaid: Pay-as-you-go (postpaid).\n" +
					"PrePaid: Monthly or yearly subscription (prepaid)." +
					"\nNote\nWhen the computing billing type is PostPaid, " +
					"the storage billing type can only be PostPaid.\n" +
					"When the computing billing type is PrePaid, " +
					"the storage billing type can be PrePaid or PostPaid.",
			},
			"instance_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Instance name. " +
					"Naming rules:\n" +
					"It cannot start with a number or a hyphen (-).\n" +
					"It can only contain Chinese characters, letters, numbers, underscores (_), and hyphens (-).\n" +
					"The length must be within 1 to 128 characters.\n" +
					"Description\nIf the instance name is not filled in, the instance ID will be used as the instance name.\n" +
					"When creating instances in batches, if an instance name is passed in, " +
					"a serial number will be automatically added after the instance name.",
			},
			"db_time_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"lower_case_table_names": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: "Whether table names are case-sensitive. " +
					"The default value is 1. Value range:\n" +
					"0: Table names are case-sensitive. The backend stores them according to the actual table name." +
					"\n1: (default) Table names are not case-sensitive. " +
					"The backend stores them by converting table names to lowercase letters.\n" +
					"Description:\nThis rule cannot be modified after creating an instance." +
					" Please set it reasonably according to business requirements.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Project name of the instance. When this parameter is left blank, the newly created instance is added to the default project by default.",
			},
			"tags": ve.TagsSchema(),
			"auto_renew": {
				Type:             schema.TypeBool,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: VeDBMysqlInstanceImportDiffSuppress,
				Description: "Whether to automatically renew under the prepaid scenario. " +
					"Values:\ntrue: Automatically renew.\nfalse: Do not automatically renew.\n" +
					"Description:\nWhen the value of ChargeType (billing type) is PrePaid (monthly/yearly package), " +
					"this parameter is required.",
			},
			"period_unit": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: VeDBMysqlInstanceImportDiffSuppress,
				Description: "Purchase cycle in prepaid scenarios.\n" +
					"Month: Monthly package.\nYear: Annual package.\nDescription:\n" +
					"When the value of ChargeType (computing billing type) is PrePaid (monthly or annual package), this parameter is required.",
			},
			"period": {
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: VeDBMysqlInstanceImportDiffSuppress,
				Description: "Purchase duration in prepaid scenarios.\n" +
					"Description:\nWhen the value of ChargeType (computing billing type) is PrePaid (monthly/yearly package), " +
					"this parameter is required.",
			},
			"pre_paid_storage_in_gb": {
				//PrePaidStorageInGB
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("storage_charge_type").(string) == "PostPaid" && k == "pre_paid_storage_in_gb" {
						return true
					}
					return false
				},
				Description: "Storage size in prepaid scenarios.\n" +
					"Description: When the value of StorageChargeType (storage billing type) is PrePaid (monthly/yearly prepaid), " +
					"this parameter is required.",
			},
		},
	}
	return resource
}

func resourceVolcengineVedbMysqlInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVedbMysqlInstance())
	if err != nil {
		return fmt.Errorf("error on creating vedb_mysql_instance %q, %s", d.Id(), err)
	}
	return resourceVolcengineVedbMysqlInstanceRead(d, meta)
}

func resourceVolcengineVedbMysqlInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVedbMysqlInstance())
	if err != nil {
		return fmt.Errorf("error on reading vedb_mysql_instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVedbMysqlInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVedbMysqlInstance())
	if err != nil {
		return fmt.Errorf("error on updating vedb_mysql_instance %q, %s", d.Id(), err)
	}
	return resourceVolcengineVedbMysqlInstanceRead(d, meta)
}

func resourceVolcengineVedbMysqlInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVedbMysqlInstance())
	if err != nil {
		return fmt.Errorf("error on deleting vedb_mysql_instance %q, %s", d.Id(), err)
	}
	return err
}

func VeDBMysqlInstanceImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	//在计费方式为PostPaid的时候 period的变化会被忽略
	if d.Get("charge_type").(string) == "PostPaid" && (k == "period" || k == "period_unit" || k == "auto_renew") {
		return true
	}

	return false
}
