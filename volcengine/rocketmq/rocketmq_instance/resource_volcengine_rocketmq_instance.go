package rocketmq_instance

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RocketmqInstance can be imported using the id, e.g.
```
$ terraform import volcengine_rocketmq_instance.default resource_id
```

*/

func ResourceVolcengineRocketmqInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRocketmqInstanceCreate,
		Read:   resourceVolcengineRocketmqInstanceRead,
		Update: resourceVolcengineRocketmqInstanceUpdate,
		Delete: resourceVolcengineRocketmqInstanceDelete,
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
				Description: "The zone id of the rocketmq instance. Support specifying multiple availability zones.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet id of the rocketmq instance.",
			},
			"version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The version of the rocketmq instance. Valid values: `4.8`.",
			},
			"compute_spec": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The compute spec of the rocketmq instance.",
			},
			"storage_space": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The storage space of the rocketmq instance.",
			},
			"auto_scale_queue": {
				Type:     schema.TypeBool,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return !(d.Id() != "" && d.HasChanges("compute_spec", "storage_space"))
				},
				Description: "Whether to create queue automatically when the spec of the instance is changed. This field is effective only when modifying `compute_field` and `storage_space`.",
			},
			"file_reserved_time": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The reserved time of messages on the RocketMQ server of the message queue. Messages that exceed the reserved time will be cleared after expiration. The unit is in hours. Valid value range is 1~72.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The instance name of the rocketmq instance.",
			},
			"instance_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The instance description of the rocketmq instance.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the rocketmq instance.",
			},
			"tags": ve.TagsSchema(),
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
							Description: "The charge type of the rocketmq instance. Valid values: `PostPaid`, `PrePaid`.",
						},
						"auto_renew": {
							Type:             schema.TypeBool,
							Optional:         true,
							Default:          false,
							DiffSuppressFunc: rocketMqInstanceImportDiffSuppress,
							Description:      "Whether to automatically renew in prepaid scenarios. Default is false.",
						},
						"period_unit": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "Monthly",
							DiffSuppressFunc: rocketMqInstanceImportDiffSuppress,
							Description:      "The purchase cycle in the prepaid scenario. Valid values: `Monthly`, `Yearly`. Default is `Monthly`.",
						},
						"period": {
							Type:             schema.TypeInt,
							Optional:         true,
							Default:          1,
							DiffSuppressFunc: rocketMqInstanceImportDiffSuppress,
							Description:      "Purchase duration in prepaid scenarios. When PeriodUnit is specified as `Monthly`, the value range is 1-9. When PeriodUnit is specified as `Yearly`, the value range is 1-3. Default is 1.",
						},
					},
				},
			},
			//"allow_list_ids": {
			//	Type:        schema.TypeSet,
			//	Optional:    true,
			//	Computed:    true,
			//	Set:         schema.HashString,
			//	Description: "Allow list Ids of the rocketmq instance.",
			//	Elem: &schema.Schema{
			//		Type: schema.TypeString,
			//	},
			//},

			// computed fields
			"instance_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the rocketmq instance.",
			},
			"region_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The region id of the rocketmq instance.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The vpc id of the rocketmq instance.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account id of the rocketmq instance.",
			},
			"eip_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The eip id of the rocketmq instance.",
			},
			"ssl_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ssl mode of the rocketmq instance.",
			},
			"enable_ssl": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the ssl authentication is enabled for the rocketmq instance.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the rocketmq instance.",
			},
			"used_topic_number": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The used topic number of the rocketmq instance.",
			},
			"used_storage_space": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The used storage space of the rocketmq instance.",
			},
			"available_queue_number": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The available queue number of the rocketmq instance.",
			},
			"used_queue_number": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The used queue number of the rocketmq instance.",
			},
			"used_group_number": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The used group number of the rocketmq instance.",
			},
			"apply_private_dns_to_public": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the private dns to public function is enabled for the rocketmq instance.",
			},
			"connection_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The connection information of the rocketmq.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The endpoint type of the rocketmq.",
						},
						"network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The network type of the rocketmq.",
						},
						"internal_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The internal endpoint of the rocketmq.",
						},
						"public_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public endpoint of the rocketmq.",
						},
						"endpoint_address_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The endpoint address ip of the rocketmq.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineRocketmqInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRocketmqInstance())
	if err != nil {
		return fmt.Errorf("error on creating rocketmq_instance %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqInstanceRead(d, meta)
}

func resourceVolcengineRocketmqInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRocketmqInstance())
	if err != nil {
		return fmt.Errorf("error on reading rocketmq_instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRocketmqInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRocketmqInstance())
	if err != nil {
		return fmt.Errorf("error on updating rocketmq_instance %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqInstanceRead(d, meta)
}

func resourceVolcengineRocketmqInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqInstanceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRocketmqInstance())
	if err != nil {
		return fmt.Errorf("error on deleting rocketmq_instance %q, %s", d.Id(), err)
	}
	return err
}

func rocketMqInstanceImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	//在计费方式为PostPaid的时候 period的变化会被忽略
	if d.Get("charge_info.0.charge_type").(string) == "PostPaid" {
		return true
	}
	if !d.HasChange("charge_info.0.charge_type") {
		return true
	}

	return false
}
