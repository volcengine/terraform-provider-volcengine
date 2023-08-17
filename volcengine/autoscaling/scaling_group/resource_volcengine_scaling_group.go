package scaling_group

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ScalingGroup can be imported using the id, e.g.
```
$ terraform import volcengine_scaling_group.default scg-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVolcengineScalingGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineScalingGroupCreate,
		Read:   resourceVolcengineScalingGroupRead,
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
				Description:  "The default cooldown interval of the scaling group. Value range: 5 ~ 86400, unit: second. Default value: 300.",
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
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The desire instance number of the scaling group.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					atoi, _ := strconv.Atoi(old)
					return atoi < 0
				},
				ValidateFunc: validation.IntAtLeast(-1),
			},
			"min_instance_number": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "The min instance number of the scaling group. Value range: 0 ~ 100.",
				ValidateFunc: validation.IntBetween(0, 2000),
			},
			"max_instance_number": {
				Type:         schema.TypeInt,
				Required:     true,
				Description:  "The max instance number of the scaling group. Value range: 0 ~ 100.",
				ValidateFunc: validation.IntBetween(0, 2000),
			},
			"instance_terminate_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The instance terminate policy of the scaling group. Valid values: OldestInstance, NewestInstance, OldestScalingConfigurationWithOldestInstance, OldestScalingConfigurationWithNewestInstance. Default value: OldestScalingConfigurationWithOldestInstance.",
				ValidateFunc: validation.StringInSlice([]string{
					"OldestInstance",
					"NewestInstance",
					"OldestScalingConfigurationWithOldestInstance",
					"OldestScalingConfigurationWithNewestInstance",
				}, false),
			},
			"server_group_attributes": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The load balancer server group attributes of the scaling group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:         schema.TypeInt,
							Required:     true,
							Description:  "The port receiving request of the server group. Value range: 1 ~ 65535.",
							ValidateFunc: validation.IntBetween(1, 65535),
						},
						"server_group_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of the server group.",
						},
						"weight": {
							Type:         schema.TypeInt,
							Required:     true,
							Description:  "The weight of the instance. Value range: 0 ~ 100.",
							ValidateFunc: validation.IntBetween(0, 100),
						},
						"load_balancer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The load balancer id.",
						},
					},
				},
				Set: serverGroupAttributeHash,
			},
			"multi_az_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"PRIORITY",
					"BALANCE",
				}, false),
				Description: "The multi az policy of the scaling group. Valid values: PRIORITY, BALANCE. Default value: PRIORITY.",
			},
			"launch_template_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the launch template bound to the scaling group. The launch template and scaling configuration cannot take effect at the same time.",
			},
			"launch_template_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version of the launch template bound to the scaling group. Valid values are the version number, Latest, or Default.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ProjectName of the scaling group.",
			},
			"tags": ve.TagsSchema(),
		},
	}
	dataSource := DataSourceVolcengineScalingGroups().Schema["scaling_groups"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineScalingGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingGroupService := NewScalingGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(scalingGroupService, d, ResourceVolcengineScalingGroup())
	if err != nil {
		return fmt.Errorf("error on creating ScalingGroup %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingGroupRead(d, meta)
}

func resourceVolcengineScalingGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	scalingGroupService := NewScalingGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(scalingGroupService, d, ResourceVolcengineScalingGroup())
	if err != nil {
		return fmt.Errorf("error on reading ScalingGroup %q, %s", d.Id(), err)
	}
	return err
}

func resourceVetackScalingGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingGroupService := NewScalingGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(scalingGroupService, d, ResourceVolcengineScalingGroup())
	if err != nil {
		return fmt.Errorf("error on updating ScalingGroup %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingGroupRead(d, meta)
}

func resourceVetackScalingGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	scalingGroupService := NewScalingGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(scalingGroupService, d, ResourceVolcengineScalingGroup())
	if err != nil {
		return fmt.Errorf("error on deleting ScalingGroup %q, %s", d.Id(), err)
	}
	return err
}

func serverGroupAttributeHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v:", m["port"]))
	buf.WriteString(fmt.Sprintf("%v:", m["server_group_id"]))
	buf.WriteString(fmt.Sprintf("%v:", m["weight"]))
	return hashcode.String(buf.String())
}
