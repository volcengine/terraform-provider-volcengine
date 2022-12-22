package ecs_instance

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ECS Instance can be imported using the id, e.g.
If Import,The data_volumes is sort by volume name
```
$ terraform import volcengine_ecs_instance.default i-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVolcengineEcsInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEcsInstanceCreate,
		Read:   resourceVolcengineEcsInstanceRead,
		Update: resourceVolcengineEcsInstanceUpdate,
		Delete: resourceVolcengineEcsInstanceDelete,
		Exists: resourceVolcengineEcsInstanceExist,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The available zone ID of ECS instance.",
			},
			"image_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Image ID of ECS instance.",
			},
			"instance_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance type of ECS instance.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of ECS instance.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of ECS instance.",
			},
			"host_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The host name of ECS instance.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The password of ECS instance.",
			},
			"key_pair_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The ssh key name of ECS instance.",
			},
			"instance_charge_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"PostPaid",
					"PrePaid",
				}, false),
				Description: "The charge type of ECS instance, the value can be `PrePaid` or `PostPaid`.",
			},
			"user_data": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: UserDateImportDiffSuppress,
				Description:      "The user data of ECS instance, this field must be encrypted with base64.",
			},
			"security_enhancement_strategy": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Active",
					"InActive",
				}, false),
				Default:     "Active",
				Description: "The security enhancement strategy of ECS instance. The value can be Active or InActive. Default is Active.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"hpc_cluster_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The hpc cluster ID of ECS instance.",
			},
			"period": {
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          12,
				DiffSuppressFunc: EcsInstanceImportDiffSuppress,
				Description:      "The period of ECS instance.Only effective when instance_charge_type is PrePaid. Default is 12. Unit is Month.",
			},
			//"period_unit": {
			//	Type:     schema.TypeString,
			//	Optional: true,
			//	Default:  "Month",
			//	ValidateFunc: validation.StringInSlice([]string{
			//		"Month",
			//	}, false),
			//	DiffSuppressFunc: ve.EcsInstanceImportDiffSuppress,
			//	Description:      "The period unit of ECS instance.Only effective when instance_charge_type is PrePaid. Default is Month.",
			//},
			"auto_renew": {
				Type:             schema.TypeBool,
				Optional:         true,
				ForceNew:         true,
				Default:          true,
				DiffSuppressFunc: EcsInstanceImportDiffSuppress,
				Description:      "The auto renew flag of ECS instance.Only effective when instance_charge_type is PrePaid. Default is true.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"auto_renew_period": {
				Type:             schema.TypeInt,
				Optional:         true,
				ForceNew:         true,
				Default:          1,
				DiffSuppressFunc: EcsInstanceImportDiffSuppress,
				Description:      "The auto renew period of ECS instance.Only effective when instance_charge_type is PrePaid. Default is 1.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},

			"include_data_volumes": {
				Type:             schema.TypeBool,
				Optional:         true,
				Default:          false,
				DiffSuppressFunc: EcsInstanceImportDiffSuppress,
				Description:      "The include data volumes flag of ECS instance.Only effective when change instance charge type.include_data_volumes.",
			},

			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet ID of primary networkInterface.",
			},

			"security_group_ids": {
				Type:        schema.TypeSet,
				Required:    true,
				MaxItems:    5,
				MinItems:    1,
				Description: "The security group ID set of primary networkInterface.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},

			"network_interface_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of primary networkInterface.",
			},

			"primary_ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The private ip address of primary networkInterface.",
			},

			"system_volume_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of system volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.",
			},

			"system_volume_size": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The size of system volume.",
			},

			"system_volume_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of system volume.",
			},

			"deployment_set_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of Ecs Deployment Set.",
			},

			"data_volumes": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    15,
				MinItems:    1,
				Description: "The data volumes collection of  ECS instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volume_type": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The type of volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.",
						},
						"size": {
							Type:        schema.TypeInt,
							Required:    true,
							ForceNew:    true,
							Description: "The size of volume.",
						},
						"delete_with_instance": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							ForceNew:    true,
							Description: "The delete with instance flag of volume.",
						},
					},
				},
			},

			"secondary_network_interfaces": {
				Type:        schema.TypeList,
				Optional:    true,
				MinItems:    1,
				Description: "The secondary networkInterface detail collection of ECS instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The subnet ID of secondary networkInterface.",
						},
						"security_group_ids": {
							Type:        schema.TypeSet,
							Required:    true,
							ForceNew:    true,
							MaxItems:    5,
							MinItems:    1,
							Description: "The security group ID set of secondary networkInterface.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"primary_ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The private ip address of secondary networkInterface.",
						},
					},
				},
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The ProjectName of the VPC.",
			},
			"tags": ve.TagsSchema(),
		},
	}
	dataSource := DataSourceVolcengineEcsInstances().Schema["instances"].Elem.(*schema.Resource).Schema
	delete(dataSource, "network_interfaces")
	delete(dataSource, "volumes")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineEcsInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	instanceService := NewEcsService(meta.(*ve.SdkClient))
	err = instanceService.Dispatcher.Create(instanceService, d, ResourceVolcengineEcsInstance())
	if err != nil {
		return fmt.Errorf("error on creating ecs instance  %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsInstanceRead(d, meta)
}

func resourceVolcengineEcsInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	instanceService := NewEcsService(meta.(*ve.SdkClient))
	err = instanceService.Dispatcher.Read(instanceService, d, ResourceVolcengineEcsInstance())
	if err != nil {
		return fmt.Errorf("error on reading ecs instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEcsInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	instanceService := NewEcsService(meta.(*ve.SdkClient))
	err = instanceService.Dispatcher.Update(instanceService, d, ResourceVolcengineEcsInstance())
	if err != nil {
		return fmt.Errorf("error on updating ecs instance  %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsInstanceRead(d, meta)
}

func resourceVolcengineEcsInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	instanceService := NewEcsService(meta.(*ve.SdkClient))
	err = instanceService.Dispatcher.Delete(instanceService, d, ResourceVolcengineEcsInstance())
	if err != nil {
		return fmt.Errorf("error on deleting ecs instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEcsInstanceExist(d *schema.ResourceData, meta interface{}) (flag bool, err error) {
	err = resourceVolcengineEcsInstanceRead(d, meta)
	if err != nil {
		if strings.Contains(err.Error(), "notfound") || strings.Contains(err.Error(), "not found") ||
			strings.Contains(err.Error(), "not exist") || strings.Contains(err.Error(), "not associate") ||
			strings.Contains(err.Error(), "invalid") || strings.Contains(err.Error(), "not_found") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
