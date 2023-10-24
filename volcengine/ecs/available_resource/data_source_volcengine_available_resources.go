package available_resource

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAvailableResources() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAvailableResourcesRead,
		Schema: map[string]*schema.Schema{
			"destination_resource": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"InstanceType", "DedicatedHost"}, false),
				Description:  "The type of resource to query. Valid values: `InstanceType`, `DedicatedHost`.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of available zone.",
			},
			"instance_type_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of instance type.",
			},
			"instance_charge_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "PostPaid",
				ValidateFunc: validation.StringInSlice([]string{"PostPaid", "PrePaid", "ReservedInstance"}, false),
				Description:  "The charge type of instance. Valid values: `PostPaid`, `PrePaid`, `ReservedInstance`. Default is `PostPaid`.",
			},
			"spot_strategy": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "NoSpot",
				ValidateFunc: validation.StringInSlice([]string{"NoSpot", "SpotAsPriceGo"}, false),
				Description:  "The spot strategy of PostPaid instance. Valid values: `NoSpot`, `SpotAsPriceGo`. Default is `NoSpot`.",
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
			"available_zones": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the region.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the available zone.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource status of the available zone. Valid values: `Available`, `SoldOut`.",
						},
						"available_resources": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The resource information of the available zone.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of resource. Valid values: `InstanceType`, `DedicatedHost`.",
									},
									"supported_resources": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The supported resource information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The value of the resource.",
												},
												"status": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The status of the resource. Valid values: `Available`, `SoldOut`.",
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
		},
	}
}

func dataSourceVolcengineAvailableResourcesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewAvailableResourceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineAvailableResources())
}
