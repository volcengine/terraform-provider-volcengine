package instance

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ESCloud Instance can be imported using the id, e.g.
```
$ terraform import volcengine_escloud_instance.default n769ewmjjqyqh5dv
```

*/

func ResourceVolcengineESCloudInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineESCloudInstanceCreate,
		Read:   resourceVolcengineESCloudInstanceRead,
		Update: resourceVolcengineESCloudInstanceUpdate,
		Delete: resourceVolcengineESCloudInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"instance_configuration": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,

				Description: "The configuration of ESCloud instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"V6_7", "V7_10"}, false),
							Description:  "The version of ESCloud instance, the value is V6_7 or V7_10.",
						},
						"region_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return d.Id() != ""
							},
							Description: "The region ID of ESCloud instance.",
						},
						"zone_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return d.Id() != ""
							},
							Description: "The available zone ID of ESCloud instance.",
						},
						"zone_number": {
							Type:        schema.TypeInt,
							Required:    true,
							ForceNew:    true,
							Description: "The zone count of the ESCloud instance used.",
						},
						"enable_https": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: "Whether Https access is enabled.",
						},
						"admin_user_name": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"admin"}, false),
							Description:  "The name of administrator account(should be admin).",
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
							ForceNew: true,
							ValidateFunc: validation.StringInSlice([]string{
								"PostPaid",
								"PrePaid",
							}, false),
							Description: "The charge type of ESCloud instance, the value can be PostPaid or PrePaid.",
						},
						"configuration_code": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Configuration code used for billing.",
						},
						"enable_pure_master": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: "Whether the Master node is independent.",
						},
						"node_specs_assigns": {
							Type:             schema.TypeList,
							Required:         true,
							Description:      "The number and configuration of various ESCloud instance node. Kibana NodeSpecsAssign should not be modified.",
							DiffSuppressFunc: nodeSpecsAssignsDiffSuppressFunc,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"Master", "Hot", "Kibana"}, false),
										Description:  "The type of node, the value is `Master` or `Hot` or `Kibana`.",
									},
									"number": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The number of node.",
									},
									"resource_spec_name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of compute resource spec, the value is `kibana.x2.small` or `es.x4.medium` or `es.x4.large` or `es.x4.xlarge` or `es.x2.2xlarge` or `es.x4.2xlarge` or `es.x2.3xlarge`.",
									},
									"storage_spec_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of storage spec. Kibana NodeSpecsAssign should not specify this field.",
									},
									"storage_size": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The size of storage. Kibana NodeSpecsAssign should not specify this field.",
									},
								},
							},
						},
						"instance_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of ESCloud instance.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The ID of subnet, the subnet must belong to the AZ selected.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The project name  to which the ESCloud instance belongs.",
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
							Description: "The maintainable date for the instance. Works only on modified scenes.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// 创建时不存在这个参数，修改时存在这个参数
								return d.Id() == ""
							},
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"force_restart_after_scale": {
							Type:             schema.TypeBool,
							Optional:         true,
							DiffSuppressFunc: forceRestartAfterScaleDiffSuppressFunc,
							Description: "Whether to force restart when changes are made. " +
								"If true, it means that the cluster will be forced to restart without paying attention to instance availability. Works only on modified the node_specs_assigns field.",
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceVolcengineESCloudInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewESCloudInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineESCloudInstance())
	if err != nil {
		return fmt.Errorf("Error on creating ESCloud instance %q,%s", d.Id(), err)
	}
	return resourceVolcengineESCloudInstanceRead(d, meta)
}

func resourceVolcengineESCloudInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewESCloudInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineESCloudInstance())
	if err != nil {
		return fmt.Errorf("error on updating ESCloud instance  %q, %s", d.Id(), err)
	}
	return resourceVolcengineESCloudInstanceRead(d, meta)
}

func resourceVolcengineESCloudInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewESCloudInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineESCloudInstance())
	if err != nil {
		return fmt.Errorf("error on deleting ecs instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineESCloudInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewESCloudInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineESCloudInstance())
	if err != nil {
		return fmt.Errorf("Error on reading ESCloud instance %q,%s", d.Id(), err)
	}
	return err
}
