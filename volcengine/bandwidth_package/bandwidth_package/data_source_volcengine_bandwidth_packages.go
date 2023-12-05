package bandwidth_package

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineBandwidthPackages() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineBandwidthPackagesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "Shared bandwidth package instance ID to be queried.",
			},
			"bandwidth_package_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Shared bandwidth package name to be queried.",
			},
			"isp": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Line types for shared bandwidth packages.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.",
			},
			"security_protection_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Security protection types for shared bandwidth packages.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of the bandwidth package to be queried.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
			},
			"tag_filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of tags.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The key of the tag.",
						},
						"values": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "The values of the tag.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
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
			"packages": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the bandwidth package.",
						},
						"bandwidth_package_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the bandwidth package.",
						},
						"bandwidth_package_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the bandwidth package.",
						},
						"isp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The line type.",
						},
						"billing_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The billing type of the bandwidth package.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The bandwidth of the bandwidth package.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol of the bandwidth package.",
						},
						"security_protection_types": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced`.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the bandwidth package.",
						},
						"tags": ve.TagsSchemaComputed(),
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the bandwidth package.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The business status of the bandwidth package.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the bandwidth package.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the bandwidth package.",
						},
						"expired_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expiration time of the bandwidth package.",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The overdue time of the bandwidth package.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deleted time of the bandwidth package.",
						},
						"eip_addresses": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of public IP information included in the shared bandwidth package.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"eip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The eip address.",
									},
									"allocation_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the eip.",
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

func dataSourceVolcengineBandwidthPackagesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewBandwidthPackageService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineBandwidthPackages())
}
