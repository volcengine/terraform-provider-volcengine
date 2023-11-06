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
				MaxItems: 10,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of the ECS instance type which the scaling configuration set. The maximum number of instance types is 10.",
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
				Description:  "The Ecs security enhancement strategy which the scaling configuration set. Valid values: Active, InActive.",
			},
			"volumes": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				MaxItems:    15,
				Description: "The list of volume of the scaling configuration. The number of supported volumes ranges from 1 to 15.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volume_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of volume.",
						},
						"size": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(10, 8192),
							Description:  "The size of volume. System disk value range: 10 - 500. The value range of the data disk: 10 - 8192.",
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
			"security_group_ids": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 5,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of the security group id of the networkInterface which the scaling configuration set." +
					" A maximum of 5 security groups can be bound at the same time, and the value ranges from 1 to 5.",
			},
			"eip_bandwidth": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(1, 500),
				Description: "The EIP bandwidth which the scaling configuration set. " +
					"When the value of Eip.BillingType is PostPaidByBandwidth, the value is 1 to 500. When the value of Eip.BillingType is PostPaidByTraffic, the value is 1 to 200.",
			},
			"eip_isp": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: eipDiffSuppressFunc,
				ValidateFunc:     validation.StringInSlice([]string{"BGP", "ChinaMobile", "ChinaUnicom", "ChinaTelecom"}, false),
				Description:      "The EIP ISP which the scaling configuration set. Valid values: BGP, ChinaMobile, ChinaUnicom, ChinaTelecom.",
			},
			"eip_billing_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: eipDiffSuppressFunc,
				ValidateFunc:     validation.StringInSlice([]string{"PostPaidByBandwidth", "PostPaidByTraffic"}, false),
				Description:      "The EIP billing type which the scaling configuration set. Valid values: PostPaidByBandwidth, PostPaidByTraffic.",
			},
			"user_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ECS user data which the scaling configuration set.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The label of the instance created by the scaling configuration. Up to 20 tags are supported.",
				MaxItems:    20,
				Set:         ve.TagsHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Key of Tags.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Value of Tags.",
						},
					},
				},
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project to which the instance created by the scaling configuration belongs.",
			},
			"hpc_cluster_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the HPC cluster to which the instance belongs. Valid only when InstanceTypes.N specifies High Performance Computing GPU Type.",
			},
			"spot_strategy": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "NoSpot",
				Description: "The preemption policy of the instance. Valid Value: NoSpot (default), SpotAsPriceGo.",
				ValidateFunc: validation.StringInSlice([]string{
					"NoSpot",
					"SpotAsPriceGo",
				}, false),
			},
			"ipv6_address_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Assign IPv6 address to instance network card. Possible values:\n0: Do not assign IPv6 address.\n1: Assign IPv6 address and the system will automatically assign an IPv6 subnet for you.",
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
	err = ve.DefaultDispatcher().Create(scalingConfigurationService, d, ResourceVolcengineScalingConfiguration())
	if err != nil {
		return fmt.Errorf("error on creating ScalingConfiguration %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingConfigurationRead(d, meta)
}

func resourceVolcengineScalingConfigurationRead(d *schema.ResourceData, meta interface{}) (err error) {
	scalingConfigurationService := NewScalingConfigurationService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(scalingConfigurationService, d, ResourceVolcengineScalingConfiguration())
	if err != nil {
		return fmt.Errorf("error on reading ScalingConfiguration %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineScalingConfigurationUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingConfigurationService := NewScalingConfigurationService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(scalingConfigurationService, d, ResourceVolcengineScalingConfiguration())
	if err != nil {
		return fmt.Errorf("error on updating ScalingConfiguration %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingConfigurationRead(d, meta)
}

func resourceVolcengineScalingConfigurationDelete(d *schema.ResourceData, meta interface{}) (err error) {
	scalingConfigurationService := NewScalingConfigurationService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(scalingConfigurationService, d, ResourceVolcengineScalingConfiguration())
	if err != nil {
		return fmt.Errorf("error on deleting ScalingConfiguration %q, %s", d.Id(), err)
	}
	return err
}
