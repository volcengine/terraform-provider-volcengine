package ecs_instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineEcsInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineInstancesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of ECS instance IDs.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The available zone ID of ECS instance.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The VPC ID of ECS instance.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of ECS instance.",
			},
			"primary_ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The primary ip address of ECS instance.",
			},
			"hpc_cluster_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The hpc cluster ID of ECS instance.",
			},
			"key_pair_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The key pair name of ECS instance.",
			},
			"instance_charge_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The charge type of ECS instance.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of ECS instance.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ProjectName of ECS instance.",
			},
			"tags": ve.TagsSchema(),
			"deployment_set_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of DeploymentSet IDs.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of ECS instance query.",
			},
			"instances": {
				Description: "The collection of ECS instance query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of ECS instance.",
						},

						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of ECS instance.",
						},

						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of ECS instance.",
						},

						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The available zone ID of ECS instance.",
						},
						"image_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The image ID of ECS instance.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of ECS instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of ECS instance.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of ECS instance.",
						},
						"host_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The host name of ECS instance.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The VPC ID of ECS instance.",
						},
						"key_pair_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ssh key name of ECS instance.",
						},
						"key_pair_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ssh key ID of ECS instance.",
						},
						"stopped_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The stop mode of ECS instance.",
						},
						"instance_charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge type of ECS instance.",
						},
						"spot_strategy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The spot strategy of ECS instance.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The spec type of ECS instance.",
						},
						"cpus": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of ECS instance CPU cores.",
						},
						"memory_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The memory size of ECS instance.",
						},
						"os_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The os type of ECS instance.",
						},
						"os_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The os name of ECS instance.",
						},

						"network_interfaces": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The networkInterface detail collection of ECS instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_interface_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of networkInterface.",
									},
									"vpc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of networkInterface.",
									},
									"subnet_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The subnet ID of networkInterface.",
									},
									"primary_ip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The private ip address of networkInterface.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of networkInterface.",
									},
									"mac_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The mac address of networkInterface.",
									},
								},
							},
						},
						"volumes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The volume detail collection of volume.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"volume_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of volume.",
									},
									"volume_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Name of volume.",
									},
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
										Description: "The delete with instance flag of volume.",
									},
								},
							},
						},
						"is_gpu": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The Flag of GPU instance.If the instance is GPU,The flag is true.",
						},
						"gpu_devices": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The GPU device info of Instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The Count of GPU device.",
									},
									"product_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Product Name of GPU device.",
									},
									"memory_size": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The Memory Size of GPU device.",
									},
									"encrypted_memory_size": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The Encrypted Memory Size of GPU device.",
									},
								},
							},
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ProjectName of ECS instance.",
						},
						"tags": ve.TagsSchemaComputed(),
						"deployment_set_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of DeploymentSet.",
						},
						"ipv6_address_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of IPv6 addresses of the ECS instance.",
						},
						"ipv6_addresses": {
							Type:        schema.TypeSet,
							Computed:    true,
							Set:         schema.HashString,
							Description: "The  IPv6 address list of the ECS instance.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineInstancesRead(d *schema.ResourceData, meta interface{}) error {
	ecsService := NewEcsService(meta.(*ve.SdkClient))
	return ve.NewRateLimitDispatcher(rateInfo).Data(ecsService, d, DataSourceVolcengineEcsInstances())
}
