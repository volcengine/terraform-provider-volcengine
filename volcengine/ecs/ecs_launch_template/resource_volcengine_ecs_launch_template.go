package ecs_launch_template

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
LaunchTemplate can be imported using the LaunchTemplateId, e.g.
When the instance launch template is modified, a new version will be created.
When the number of versions reaches the upper limit (30), the oldest version that is not the default version will be deleted.
```
$ terraform import volcengine_ecs_launch_template.default lt-ysxc16auaugh9zfy****
```

*/

func ResourceVolcengineEcsLaunchTemplate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEcsLaunchTemplateCreate,
		Read:   resourceVolcengineEcsLaunchTemplateRead,
		Update: resourceVolcengineEcsLaunchTemplateUpdate,
		Delete: resourceVolcengineEcsLaunchTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"launch_template_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the launch template.",
			},
			"version_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The latest version description of the launch template.",
			},
			"instance_type_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The compute type of the instance.",
			},
			"image_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The image ID.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the instance.",
			},
			"instance_charge_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The charge type of the instance and volume.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the instance.",
			},
			"host_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The host name of the instance.",
			},
			"unique_suffix": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether the ordered suffix is automatically added to Hostname and InstanceName when multiple instances are created.",
			},
			"suffix_index": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The index of the ordered suffix.",
			},
			"key_pair_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "When you log in to the instance using the SSH key pair, enter the name of the key pair.",
			},
			"security_enhancement_strategy": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Active", "InActive"}, false),
				Description:  "Whether to open the security reinforcement.",
			},
			"volumes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of volume of the scaling configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volume_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of volume.",
						},
						"size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The size of volume.",
						},
						"delete_with_instance": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "The delete with instance flag of volume. Valid values: true, false. Default value: true.",
						},
					},
				},
			},
			"network_interfaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of network interfaces. When creating an instance, it is supported to bind auxiliary network cards at the same time. The first one is the primary network card, and the others are secondary network cards.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The private network subnet ID of the instance, when creating the instance, supports binding the secondary NIC at the same time.",
						},
						"security_group_ids": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							MaxItems:    5,
							MinItems:    1,
							Optional:    true,
							Description: "The security group ID associated with the NIC.",
						},
					},
				},
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The vpc id.",
			},
			"eip_bandwidth": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				ValidateFunc: validation.IntBetween(1, 500),
				Description:  "The EIP bandwidth which the scaling configuration set.",
			},
			"eip_isp": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "BGP",
				ValidateFunc: validation.StringInSlice([]string{"BGP", "ChinaMobile", "ChinaUnicom", "ChinaTelecom"}, false),
				Description:  "The EIP ISP which the scaling configuration set. Valid values: BGP, ChinaMobile, ChinaUnicom, ChinaTelecom.",
			},
			"eip_billing_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"PostPaidByBandwidth", "PostPaidByTraffic"}, false),
				Description:  "The EIP billing type which the scaling configuration set. Valid values: PostPaidByBandwidth, PostPaidByTraffic.",
			},
			"user_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance custom data. The set custom data must be Base64 encoded, and the size of the custom data before Base64 encoding cannot exceed 16KB.",
			},
			"hpc_cluster_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The hpc cluster id.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The zone id.",
			},
			"launch_template_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The launch template id.",
			},
		},
	}
	return resource
}

func resourceVolcengineEcsLaunchTemplateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsLaunchTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineEcsLaunchTemplate())
	if err != nil {
		return fmt.Errorf("error creating launch template service: %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsLaunchTemplateRead(d, meta)
}

func resourceVolcengineEcsLaunchTemplateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsLaunchTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineEcsLaunchTemplate())
	if err != nil {
		return fmt.Errorf("error reading launch template service: %q, %s", d.Id(), err)
	}
	return nil
}

func resourceVolcengineEcsLaunchTemplateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsLaunchTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineEcsLaunchTemplate())
	if err != nil {
		return fmt.Errorf("error updating launch template service: %q, %s", d.Id(), err)
	}
	return nil
}

func resourceVolcengineEcsLaunchTemplateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsLaunchTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineEcsLaunchTemplate())
	if err != nil {
		return fmt.Errorf("error deleting launch template service: %q, %s", d.Id(), err)
	}
	return nil
}
