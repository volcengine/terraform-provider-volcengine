package rabbitmq_instance

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RabbitmqInstance can be imported using the id, e.g.
```
$ terraform import volcengine_rabbitmq_instance.default resource_id
```

*/

func ResourceVolcengineRabbitmqInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRabbitmqInstanceCreate,
		Read:   resourceVolcengineRabbitmqInstanceRead,
		Update: resourceVolcengineRabbitmqInstanceUpdate,
		Delete: resourceVolcengineRabbitmqInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"zone_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				MinItems: 1,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The zone id of the rabbitmq instance. Support specifying multiple availability zones.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet id of the rabbitmq instance.",
			},
			"version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The version of the rabbitmq instance. Valid values: `3.8.18`, `3.12`.",
			},
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The administrator name of the rabbitmq instance.",
			},
			"user_password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The administrator password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"compute_spec": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The compute specification of the rabbitmq instance.",
			},
			"storage_space": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The storage space of the rabbitmq instance. Unit: GiB. The valid value must be specified as a multiple of 100.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the rabbitmq instance.",
			},
			"instance_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the rabbitmq instance.",
			},
			"charge_info": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The charge information of the rocketmq instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"charge_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The charge type of the rabbitmq instance. Valid values: `PostPaid`, `PrePaid`.",
						},
						"auto_renew": {
							Type:             schema.TypeBool,
							Optional:         true,
							Default:          false,
							DiffSuppressFunc: rabbitMqInstanceImportDiffSuppress,
							Description:      "Whether to automatically renew in prepaid scenarios. Default is false.",
						},
						"period_unit": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "Month",
							DiffSuppressFunc: rabbitMqInstanceImportDiffSuppress,
							Description:      "The purchase cycle in the prepaid scenario. Valid values: `Month`, `Year`. Default is `Month`.",
						},
						"period": {
							Type:             schema.TypeInt,
							Optional:         true,
							Default:          1,
							DiffSuppressFunc: rabbitMqInstanceImportDiffSuppress,
							Description:      "Purchase duration in prepaid scenarios. When PeriodUnit is specified as `Month`, the value range is 1-9. When PeriodUnit is specified as `Year`, the value range is 1-3. Default is 1.",
						},
					},
				},
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The IAM project name where the rabbitmq instance resides.",
			},
			"tags": ve.TagsSchema(),

			// computed fields
			"instance_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the rabbitmq instance.",
			},
			"arch_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the rabbitmq instance.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the rabbitmq instance.",
			},
			"eip_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The eip id of the rabbitmq instance.",
			},
			"apply_private_dns_to_public": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether enable the public network parsing function of the rabbitmq instance.",
			},
			"init_user_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The WebUI admin user name of the rabbitmq instance.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account id of the rabbitmq instance.",
			},
			"region_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The region id of the rabbitmq instance.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The vpc id of the rabbitmq instance.",
			},
			"used_storage_space": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The used storage space of the rabbitmq instance. Unit: GiB.",
			},
			"endpoints": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The endpoint info of the rabbitmq instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The endpoint type of the rabbitmq instance.",
						},
						"internal_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The internal endpoint of the rabbitmq instance.",
						},
						"network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network type of the rabbitmq instance.",
						},
						"public_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public endpoint of the rabbitmq instance.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineRabbitmqInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRabbitmqInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRabbitmqInstance())
	if err != nil {
		return fmt.Errorf("error on creating rabbitmq_instance %q, %s", d.Id(), err)
	}
	return resourceVolcengineRabbitmqInstanceRead(d, meta)
}

func resourceVolcengineRabbitmqInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRabbitmqInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRabbitmqInstance())
	if err != nil {
		return fmt.Errorf("error on reading rabbitmq_instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRabbitmqInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRabbitmqInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRabbitmqInstance())
	if err != nil {
		return fmt.Errorf("error on updating rabbitmq_instance %q, %s", d.Id(), err)
	}
	return resourceVolcengineRabbitmqInstanceRead(d, meta)
}

func resourceVolcengineRabbitmqInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRabbitmqInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRabbitmqInstance())
	if err != nil {
		return fmt.Errorf("error on deleting rabbitmq_instance %q, %s", d.Id(), err)
	}
	return err
}

func rabbitMqInstanceImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	//在计费方式为PostPaid的时候 period的变化会被忽略
	if d.Get("charge_info.0.charge_type").(string) == "PostPaid" {
		return true
	}
	if !d.HasChange("charge_info.0.charge_type") {
		return true
	}

	return false
}
