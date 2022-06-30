package scaling_configuration

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ScalingConfiguration can be imported using the id, e.g.
```
$ terraform import volcengine_scaling_configuration.default scc-ybkuck3mx8cm9tm5yglz
```

*/

func ResourceVolcengineScalingConfiguration() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineScalingConfigurationCreate,
		Read:   resourceVolcengineScalingConfigurationRead,
		Update: resourceVolcengineScalingConfigurationUpdate,
		Delete: resourceVolcengineScalingConfigurationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"active": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "The active flag of the scaling configuration. when set true, the scaling group which the scaling configuration belongs will use it.",
			},
			"enable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The enable flag of the scaling group. when set true, the scaling group which the scaling configuration belongs will be enabled.",
			},
			"substitute": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: substituteDiffSuppressFunc,
				Description:      "The id of the substitute scaling configuration. when the active flag set false with the scaling configuration lifecycle state active, it must a valid substitute.",
			},
			"lifecycle_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the scaling configuration.",
			},
			"scaling_configuration_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the scaling configuration.",
			},
			"scaling_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the scaling group to which the scaling configuration belongs.",
			},
			"image_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ECS image id which the scaling configuration set.",
			},
			"instance_types": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of the ECS instance type which the scaling configuration set.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ECS instance name which the scaling configuration set.",
			},
			"instance_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ECS instance description which the scaling configuration set.",
			},
			"host_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ECS hostname which the scaling configuration set.",
			},
			"password": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				AtLeastOneOf: []string{"password", "key_pair_name"},
				Description:  "The ECS password which the scaling configuration set.",
			},
			"key_pair_name": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"password", "key_pair_name"},
				Description:  "The ECS key pair name which the scaling configuration set.",
			},
			"security_enhancement_strategy": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Active",
				ValidateFunc: validation.StringInSlice([]string{"Active", "InActive"}, false),
				Description:  "The Ecs security enhancement strategy which the scaling configuration set.",
			},
			"volumes": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Description: "The list of volume of the scaling configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volume_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of volume.",
						},
						"size": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The size of volume.",
						},
						"delete_with_instance": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "The delete with instance flag of volume.",
						},
					},
				},
			},
			"security_group_ids": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of the security group id of the networkInterface which the scaling configuration set.",
			},
			"eip_bandwidth": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "The EIP bandwidth which the scaling configuration set.",
			},
			"eip_isp": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: eipDiffSuppressFunc,
				ValidateFunc:     validation.StringInSlice([]string{"BGP", "ChinaMobile", "ChinaUnicom", "ChinaTelecom"}, false),
				Description:      "The EIP ISP which the scaling configuration set.",
			},
			"eip_billing_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: eipDiffSuppressFunc,
				ValidateFunc:     validation.StringInSlice([]string{"PostPaidByBandwidth", "PostPaidByTraffic"}, false),
				Description:      "The EIP billing type which the scaling configuration set.",
			},
			"user_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ECS user data which the scaling configuration set.",
			},
		},
	}
	dataSource := DataSourceVolcengineScalingConfigurations().Schema["scaling_configurations"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineScalingConfigurationCreate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingConfigurationService := NewScalingConfigurationService(meta.(*ve.SdkClient))
	err = scalingConfigurationService.Dispatcher.Create(scalingConfigurationService, d, ResourceVolcengineScalingConfiguration())
	if err != nil {
		return fmt.Errorf("error on creating ScalingConfiguration %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingConfigurationRead(d, meta)
}

func resourceVolcengineScalingConfigurationRead(d *schema.ResourceData, meta interface{}) (err error) {
	scalingConfigurationService := NewScalingConfigurationService(meta.(*ve.SdkClient))
	err = scalingConfigurationService.Dispatcher.Read(scalingConfigurationService, d, ResourceVolcengineScalingConfiguration())
	if err != nil {
		return fmt.Errorf("error on reading ScalingConfiguration %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineScalingConfigurationUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingConfigurationService := NewScalingConfigurationService(meta.(*ve.SdkClient))
	err = scalingConfigurationService.Dispatcher.Update(scalingConfigurationService, d, ResourceVolcengineScalingConfiguration())
	if err != nil {
		return fmt.Errorf("error on updating ScalingConfiguration %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingConfigurationRead(d, meta)
}

func resourceVolcengineScalingConfigurationDelete(d *schema.ResourceData, meta interface{}) (err error) {
	scalingConfigurationService := NewScalingConfigurationService(meta.(*ve.SdkClient))
	err = scalingConfigurationService.Dispatcher.Delete(scalingConfigurationService, d, ResourceVolcengineScalingConfiguration())
	if err != nil {
		return fmt.Errorf("error on deleting ScalingConfiguration %q, %s", d.Id(), err)
	}
	return err
}
