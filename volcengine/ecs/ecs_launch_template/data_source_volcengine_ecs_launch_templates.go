package ecs_launch_template

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineEcsLaunchTemplates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineEcsLaunchTemplatesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "A list of launch template ids.",
			},
			"launch_template_names": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "A list of launch template names.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of scaling policy.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of scaling policy query.",
			},
			"launch_templates": {
				Description: "The collection of launch templates.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the launch template.",
						},
						"launch_template_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the launch template.",
						},
						"launch_template_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the launch template.",
						},
						"default_version_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The default version of the launch template.",
						},
						"latest_version_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The latest version of the launch template.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The created time of the launch template.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The updated time of the launch template.",
						},
						"version_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The latest version description of the launch template.",
						},
						"image_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The image id.",
						},
						"security_enhancement_strategy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether to open the security reinforcement.",
						},
						"key_pair_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "When you log in to the instance using the SSH key pair, enter the name of the key pair.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the instance.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the instance.",
						},
						"host_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The host name of the instance.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone ID of the instance.",
						},
						"volumes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of volume of the scaling configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"volume_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of volume.",
									},
									"size": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The size of volume.",
									},
									"delete_with_instance": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "The delete with instance flag of volume. Valid values: true, false. Default value: true.",
									},
								},
							},
						},
						"network_interfaces": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of network interfaces.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subnet_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The private network subnet ID of the instance, when creating the instance, supports binding the secondary NIC at the same time.",
									},
									"security_group_ids": {
										Type:        schema.TypeList,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Computed:    true,
										Description: "The security group ID associated with the NIC.",
									},
								},
							},
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id.",
						},
						"eip_bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The EIP bandwidth which the scaling configuration set.",
						},
						"eip_isp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The EIP ISP which the scaling configuration set. Valid values: BGP, ChinaMobile, ChinaUnicom, ChinaTelecom.",
						},
						"eip_billing_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The EIP billing type which the scaling configuration set. Valid values: PostPaidByBandwidth, PostPaidByTraffic.",
						},
						"hpc_cluster_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The hpc cluster id.",
						},
						"unique_suffix": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the ordered suffix is automatically added to Hostname and InstanceName when multiple instances are created.",
						},
						"suffix_index": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The index of the ordered suffix.",
						},
						"instance_charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge type of the instance and volume.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineEcsLaunchTemplatesRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsLaunchTemplateService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineEcsLaunchTemplates())
}
