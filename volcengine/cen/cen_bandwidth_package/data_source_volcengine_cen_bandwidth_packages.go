package cen_bandwidth_package

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCenBandwidthPackages() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCenBandwidthPackagesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of cen bandwidth package IDs.",
			},
			"cen_bandwidth_package_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of cen bandwidth package names.",
			},
			"cen_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A cen id.",
			},
			"local_geographic_region_set_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A local geographic region set id.",
			},
			"peer_geographic_region_set_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A peer geographic region set id.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of cen bandwidth package.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of cen bandwidth package query.",
			},
			"tags": ve.TagsSchema(),
			"bandwidth_packages": {
				Description: "The collection of cen bandwidth package query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the cen bandwidth package.",
						},
						"cen_bandwidth_package_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the cen bandwidth package.",
						},
						"cen_bandwidth_package_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the cen bandwidth package.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the cen bandwidth package.",
						},
						"local_geographic_region_set_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The local geographic region set id of the cen bandwidth package.",
						},
						"peer_geographic_region_set_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The peer geographic region set id of the cen bandwidth package.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The bandwidth of the cen bandwidth package.",
						},
						"remaining_bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The remain bandwidth of the cen bandwidth package.",
						},
						"cen_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The cen IDs of the bandwidth package.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the cen bandwidth package.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the cen bandwidth package.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the cen bandwidth package.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID of the cen bandwidth package.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The business status of the cen bandwidth package.",
						},
						"billing_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The billing type of the cen bandwidth package.",
						},
						"expired_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expired time of the cen bandwidth package.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deleted time of the cen bandwidth package.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCenBandwidthPackagesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCenBandwidthPackageService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineCenBandwidthPackages())
}
