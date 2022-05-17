package ecs_instance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
ECS Instance can be imported using the id, e.g.
```
$ terraform import vestack_ecs_instance.default i-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVestackEcsInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackEcsInstanceCreate,
		Read:   resourceVestackEcsInstanceRead,
		Update: resourceVestackEcsInstanceUpdate,
		Delete: resourceVestackEcsInstanceDelete,
		Exists: resourceVestackEcsInstanceExist,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
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
				Description: "The charge type of ECS instance.",
			},
			"user_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The user data of ECS instance.",
			},
			"security_enhancement_strategy": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Active",
					"InActive",
				}, false),
				Default:          "Active",
				DiffSuppressFunc: ve.EcsInstanceImportDiffSuppress,
				Description:      "The security enhancement strategy of ECS instance.Default is true.",
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
				DiffSuppressFunc: ve.EcsInstanceImportDiffSuppress,
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
				DiffSuppressFunc: ve.EcsInstanceImportDiffSuppress,
				Description:      "The auto renew flag of ECS instance.Only effective when instance_charge_type is PrePaid. Default is true.",
			},
			"auto_renew_period": {
				Type:             schema.TypeInt,
				Optional:         true,
				ForceNew:         true,
				Default:          1,
				DiffSuppressFunc: ve.EcsInstanceImportDiffSuppress,
				Description:      "The auto renew period of ECS instance.Only effective when instance_charge_type is PrePaid. Default is 1.",
			},

			"include_data_volumes": {
				Type:             schema.TypeBool,
				Optional:         true,
				Default:          false,
				DiffSuppressFunc: ve.EcsInstanceImportDiffSuppress,
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

			"system_volume_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of system volume.",
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

			"data_volumes": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    15,
				MinItems:    1,
				Description: "The data volume collection of  ECS instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volume_type": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The type of volume.",
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
					},
				},
			},
		},
	}
	dataSource := DataSourceVestackEcsInstances().Schema["instances"].Elem.(*schema.Resource).Schema
	delete(dataSource, "network_interfaces")
	delete(dataSource, "volumes")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVestackEcsInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	instanceService := NewEcsService(meta.(*ve.SdkClient))
	err = instanceService.Dispatcher.Create(instanceService, d, ResourceVestackEcsInstance())
	if err != nil {
		return fmt.Errorf("error on creating ecs instance  %q, %s", d.Id(), err)
	}
	return resourceVestackEcsInstanceRead(d, meta)
}

func resourceVestackEcsInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	instanceService := NewEcsService(meta.(*ve.SdkClient))
	err = instanceService.Dispatcher.Read(instanceService, d, ResourceVestackEcsInstance())
	if err != nil {
		return fmt.Errorf("error on reading ecs instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackEcsInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	instanceService := NewEcsService(meta.(*ve.SdkClient))
	err = instanceService.Dispatcher.Update(instanceService, d, ResourceVestackEcsInstance())
	if err != nil {
		return fmt.Errorf("error on updating ecs instance  %q, %s", d.Id(), err)
	}
	return resourceVestackEcsInstanceRead(d, meta)
}

func resourceVestackEcsInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	instanceService := NewEcsService(meta.(*ve.SdkClient))
	err = instanceService.Dispatcher.Delete(instanceService, d, ResourceVestackEcsInstance())
	if err != nil {
		return fmt.Errorf("error on deleting ecs instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackEcsInstanceExist(d *schema.ResourceData, meta interface{}) (flag bool, err error) {
	err = resourceVestackEcsInstanceRead(d, meta)
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
