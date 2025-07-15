package ecs_invocation

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
EcsInvocation can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_invocation.default ivk-ychnxnm45dl8j0mm****
```

*/

func ResourceVolcengineEcsInvocation() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEcsInvocationCreate,
		Read:   resourceVolcengineEcsInvocationRead,
		Update: resourceVolcengineEcsInvocationUpdate,
		Delete: resourceVolcengineEcsInvocationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"command_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The command id of the ecs invocation.",
			},
			"instance_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list of ECS instance IDs.",
			},
			"invocation_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the ecs invocation.",
			},
			"invocation_description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The description of the ecs invocation.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The username of the ecs command. When this field is not specified, use the value of the field with the same name in ecs command as the default value.",
			},
			"working_dir": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The working directory of the ecs invocation. When this field is not specified, use the value of the field with the same name in ecs command as the default value.",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The timeout of the ecs command. Unit: seconds. Valid value range: 30~86400. Default is 60.",
			},
			"repeat_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Once",
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Once",
					"Rate",
					"Fixed",
				}, false),
				Description: "The repeat mode of the ecs invocation. Valid values: `Once`, `Rate`, `Fixed`. Default is `Once`.",
			},
			"frequency": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("repeat_mode").(string) == "Rate" {
						return false
					}
					return true
				},
				Description: "The frequency of the ecs invocation. This field is valid and required when the value of the repeat_mode field is `Rate`.",
			},
			"launch_time": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("repeat_mode").(string) == "Once" {
						return true
					}
					return false
				},
				Description: "The launch time of the ecs invocation. RFC3339 format. This field is valid and required when the value of the repeat_mode field is `Rate` or `Fixed`.",
			},
			"recurrence_end_time": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("repeat_mode").(string) == "Rate" {
						return false
					}
					return true
				},
				Description: "The recurrence end time of the ecs invocation. RFC3339 format. This field is valid and required when the value of the repeat_mode field is `Rate`.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the ecs command.",
			},
			"tags": ve.TagsSchema(),
			"parameters": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "The custom parameters of the ecs command. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The name of the parameter.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The value of the parameter.",
						},
					},
				},
			},

			// computed fields
			"invocation_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the ecs invocation.",
			},
			"start_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The start time of the ecs invocation.",
			},
			"end_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The end time of the ecs invocation.",
			},
		},
	}
	return resource
}

func resourceVolcengineEcsInvocationCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsInvocationService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineEcsInvocation())
	if err != nil {
		return fmt.Errorf("error on creating ecs invocation %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsInvocationRead(d, meta)
}

func resourceVolcengineEcsInvocationRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsInvocationService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineEcsInvocation())
	if err != nil {
		return fmt.Errorf("error on reading ecs invocation %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEcsInvocationUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsInvocationService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineEcsInvocation())
	if err != nil {
		return fmt.Errorf("error on updating ecs invocation %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsInvocationRead(d, meta)
}

func resourceVolcengineEcsInvocationDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsInvocationService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineEcsInvocation())
	if err != nil {
		return fmt.Errorf("error on deleting ecs invocation %q, %s", d.Id(), err)
	}
	return err
}
