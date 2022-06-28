package scaling_group

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
ScalingGroup can be imported using the id, e.g.
```
$ terraform import vestack_scaling_group.default scg-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVestackScalingGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackScalingGroupCreate,
		Read:   resourceVestackScalingGroupRead,
		Update: resourceVetackScalingGroupUpdate,
		Delete: resourceVetackScalingGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"scaling_group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the scaling group.",
			},
			"default_cooldown": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				Description:  "The default cooldown interval of the scaling group.",
				ValidateFunc: validation.IntBetween(5, 86400),
			},
			"subnet_ids": {
				Type:        schema.TypeList, // 子网顺序会影响优先级策略，需要为list类型
				Required:    true,
				Description: "The list of the subnet id to which the ENI is connected.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"desire_instance_number": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				Description:  "The desire instance number of the scaling group.",
				ValidateFunc: validation.IntAtLeast(-1),
			},
			"min_instance_number": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "The min instance number of the scaling group.",
				ValidateFunc: validation.IntAtLeast(0),
			},
			"max_instance_number": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "The max instance number of the scaling group.",
				ValidateFunc: validation.IntAtLeast(0),
			},
			"instance_terminate_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The instance terminate policy of the scaling group.",
				ValidateFunc: validation.StringInSlice([]string{
					"OldestInstance",
					"NewestInstance",
					"OldestScalingConfigurationWithOldestInstance",
					"OldestScalingConfigurationWithNewestInstance",
				}, false),
			},
			"db_instance_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list of the DB instance ids.",
			},
			"multi_az_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"PRIORITY",
					"BALANCE",
				}, false),
				Description: "The multi az policy of the scaling group.",
			},
		},
	}
	dataSource := DataSourceVestackScalingGroups().Schema["scaling_groups"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVestackScalingGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingGroupService := NewScalingGroupService(meta.(*ve.SdkClient))
	err = scalingGroupService.Dispatcher.Create(scalingGroupService, d, ResourceVestackScalingGroup())
	if err != nil {
		return fmt.Errorf("error on creating ScalingGroup %q, %s", d.Id(), err)
	}
	return resourceVestackScalingGroupRead(d, meta)
}

func resourceVestackScalingGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	scalingGroupService := NewScalingGroupService(meta.(*ve.SdkClient))
	err = scalingGroupService.Dispatcher.Read(scalingGroupService, d, ResourceVestackScalingGroup())
	if err != nil {
		return fmt.Errorf("error on reading ScalingGroup %q, %s", d.Id(), err)
	}
	return err
}

func resourceVetackScalingGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingGroupService := NewScalingGroupService(meta.(*ve.SdkClient))
	err = scalingGroupService.Dispatcher.Update(scalingGroupService, d, ResourceVestackScalingGroup())
	if err != nil {
		return fmt.Errorf("error on updating ScalingGroup %q, %s", d.Id(), err)
	}
	return resourceVestackScalingGroupRead(d, meta)
}

func resourceVetackScalingGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	scalingGroupService := NewScalingGroupService(meta.(*ve.SdkClient))
	err = scalingGroupService.Dispatcher.Delete(scalingGroupService, d, ResourceVestackScalingGroup())
	if err != nil {
		return fmt.Errorf("error on deleting ScalingGroup %q, %s", d.Id(), err)
	}
	return err
}
