package escloud_instance_v2

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
EscloudInstanceV2 can be imported using the id, e.g.
```
$ terraform import volcengine_escloud_instance_v2.default resource_id
```

*/

func ResourceVolcengineEscloudInstanceV2() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEscloudInstanceV2Create,
		Read:   resourceVolcengineEscloudInstanceV2Read,
		Update: resourceVolcengineEscloudInstanceV2Update,
		Delete: resourceVolcengineEscloudInstanceV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"version": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The version of instance. When creating ESCloud instance, the valid value is `V6_7` or `V7_10`. When creating OpenSearch instance, the valid value is `OPEN_SEARCH_2_9`.",
			},
			"zone_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				MinItems: 1,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The zone id of the ESCloud instance. Support specifying multiple availability zones.\n " +
					"The first zone id is the primary availability zone, while the rest are backup availability zones.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of subnet, the subnet must belong to the AZ selected.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of ESCloud instance.",
			},
			"enable_https": {
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    true,
				Description: "Whether Https access is enabled.",
			},
			"admin_password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The password of administrator account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"charge_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"PostPaid",
					"PrePaid",
				}, false),
				Description: "The charge type of ESCloud instance, valid values: `PostPaid`, `PrePaid`.",
			},
			"auto_renew": {
				Type:             schema.TypeBool,
				Optional:         true,
				Default:          false,
				DiffSuppressFunc: esCloudInstanceImportDiffSuppress,
				Description:      "Whether to automatically renew in prepaid scenarios. Default is false.",
			},
			"period": {
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          1,
				DiffSuppressFunc: esCloudInstanceImportDiffSuppress,
				Description:      "Purchase duration in prepaid scenarios. Unit: Monthly.",
			},
			"configuration_code": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Configuration code used for billing.",
			},
			"enable_pure_master": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Whether the Master node is independent.",
			},
			"deletion_protection": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether enable deletion protection for ESCloud instance. Default is false.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name to which the ESCloud instance belongs.",
			},
			"tags": ve.TagsSchema(),
			"network_specs": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "The public network config of the ESCloud instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The type of public network, valid values: `Elasticsearch`, `Kibana`.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Required:    true,
							ForceNew:    true,
							Description: "The bandwidth of the eip. Unit: Mbps.",
						},
						"is_open": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: "Whether the eip is opened.",
						},
						"spec_name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The spec name of public network.",
						},
					},
				},
			},
			"node_specs_assigns": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Description: "The number and configuration of various ESCloud instance node. Kibana NodeSpecsAssign should not be modified.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of node, valid values: `Master`, `Hot`, `Cold`, `Warm`, `Kibana`, `Coordinator`.",
						},
						"number": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The number of node.",
						},
						"resource_spec_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of compute resource spec.",
						},
						"storage_spec_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of storage spec. Kibana NodeSpecsAssign should specify this field to ``.",
						},
						"storage_size": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The size of storage. Unit: GiB. the adjustment step size is 10GiB. Default is 100 GiB. Kibana NodeSpecsAssign should specify this field to 0.",
						},
						"extra_performance": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The extra performance of FlexPL storage spec.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"throughput": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "When your data node chooses to use FlexPL storage type and the storage specification configuration is 500GiB or above, it supports purchasing bandwidth packages to increase disk bandwidth.\nThe unit is MiB, and the adjustment step size is 10MiB.",
									},
								},
							},
						},
					},
				},
			},
			"maintenance_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The maintainable time period for the instance. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"maintenance_day": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Set:         schema.HashString,
				Description: "The maintainable day for the instance. Valid values: `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, `SATURDAY`. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// computed fields
			"es_eip_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The eip id associated with the instance.",
			},
			"es_eip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The eip address of instance.",
			},
			"kibana_eip_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The eip id associated with kibana.",
			},
			"kibana_eip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The eip address of kibana.",
			},
			"main_zone_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The main zone id of instance.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of instance.",
			},
			"es_public_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The es public domain of instance.",
			},
			"es_private_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The es private domain of instance.",
			},
			"es_public_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The es public endpoint of instance.",
			},
			"es_private_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The es private endpoint of instance.",
			},
			"kibana_private_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The kibana private domain of instance.",
			},
			"kibana_public_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The kibana public domain of instance.",
			},
			"cerebro_private_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The cerebro private domain of instance.",
			},
			"cerebro_public_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The cerebro public domain of instance.",
			},
			"es_public_ip_whitelist": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The whitelist of es public ip.",
			},
			"es_private_ip_whitelist": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The whitelist of es private ip.",
			},
			"kibana_public_ip_whitelist": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The whitelist of kibana public ip.",
			},
			"kibana_private_ip_whitelist": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The whitelist of kibana private ip.",
			},
		},
	}
	return resource
}

func resourceVolcengineEscloudInstanceV2Create(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEscloudInstanceV2Service(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineEscloudInstanceV2())
	if err != nil {
		return fmt.Errorf("error on creating escloud_instance_v2 %q, %s", d.Id(), err)
	}
	return resourceVolcengineEscloudInstanceV2Read(d, meta)
}

func resourceVolcengineEscloudInstanceV2Read(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEscloudInstanceV2Service(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineEscloudInstanceV2())
	if err != nil {
		return fmt.Errorf("error on reading escloud_instance_v2 %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEscloudInstanceV2Update(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEscloudInstanceV2Service(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineEscloudInstanceV2())
	if err != nil {
		return fmt.Errorf("error on updating escloud_instance_v2 %q, %s", d.Id(), err)
	}
	return resourceVolcengineEscloudInstanceV2Read(d, meta)
}

func resourceVolcengineEscloudInstanceV2Delete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEscloudInstanceV2Service(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineEscloudInstanceV2())
	if err != nil {
		return fmt.Errorf("error on deleting escloud_instance_v2 %q, %s", d.Id(), err)
	}
	return err
}

func esCloudInstanceImportDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	//在计费方式为PostPaid的时候 period的变化会被忽略
	if d.Get("charge_type").(string) == "PostPaid" {
		return true
	}
	if !d.HasChange("charge_type") {
		return true
	}

	return false
}
