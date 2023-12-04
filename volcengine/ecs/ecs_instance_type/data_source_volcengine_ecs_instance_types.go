package ecs_instance_type

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineEcsInstanceTypes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineEcsInstanceTypesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
				Description: "A list of instance type IDs. " +
					"When the number of ids is greater than 10, only the first 10 are effective.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"instance_types": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_type_family": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance type family.",
						},
						"gpu": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The GPU device info of Instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gpu_devices": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "GPU device information list.",
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
												"memory": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Graphics memory information.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"size": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The Memory Size of GPU device.",
															},
															"encrypted_size": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The Encrypted Memory Size of GPU device.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"baseline_credit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The CPU benchmark performance that can be provided steadily by on-demand instances is determined by the instance type.",
						},
						"initial_credit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The CPU credits obtained at once when creating a on-demand performance instance are fixed at 30 credits per vCPU.",
						},
						"instance_type_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the instance type.",
						},
						"processor": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "CPU information of instance specifications.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpus": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of ECS instance CPU cores.",
									},
									"model": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "CPU model.",
									},
									"base_frequency": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "CPU clock speed, unit: GHz.",
									},
									"turbo_frequency": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "CPU Turbo Boost, unit: GHz.",
									},
								},
							},
						},
						"rdma": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "RDMA Specification Information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rdma_network_interfaces": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of RDMA network cards.",
									},
								},
							},
						},
						"memory": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Memory information of instance specifications.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"size": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Memory size, unit: MiB.",
									},
									"encrypted_size": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The Encrypted Memory Size of GPU device.",
									},
								},
							},
						},
						"network": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Network information of instance specifications.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"maximum_network_interfaces": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Maximum number of elastic network interfaces supported for attachment.",
									},
									"maximum_private_ipv4_addresses_per_network_interface": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Maximum number of IPv4 addresses for a single elastic network interface.",
									},
									"maximum_queues_per_network_interface": {
										Type:     schema.TypeInt,
										Computed: true,
										Description: "Maximum queue number for a single elastic network interface, " +
											"including the queue number supported by the primary network interface and the auxiliary network interface.",
									},
									"maximum_throughput_kpps": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Network packet sending and receiving capacity (in+out), unit: Kpps.",
									},
									"baseline_bandwidth_mbps": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Network benchmark bandwidth capacity (out/in), unit: Mbps.",
									},
									"maximum_bandwidth_mbps": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Peak network bandwidth capacity (out/in), unit: Mbps.",
									},
								},
							},
						},
						"local_volumes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Local disk configuration information corresponding to instance specifications.",
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
									"count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of local disks mounted on the instance.",
									},
								},
							},
						},
						"volume": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Cloud disk information for instance specifications.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"supported_volume_types": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "List of supported volume types.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"maximum_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum number of volumes.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineEcsInstanceTypesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewEcsInstanceTypeService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineEcsInstanceTypes())
}
