package private_zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcenginePrivateZones() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcenginePrivateZonesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Private Zone IDs.",
			},
			"zone_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of Private Zone.",
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The region of Private Zone.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The vpc id associated to Private Zone.",
			},
			"recursion_mode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the recursion mode of Private Zone is enabled.",
			},
			"line_mode": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The line mode of Private Zone, specified whether the intelligent mode and the load balance function is enabled.",
			},
			"search_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "LIKE",
				ValidateFunc: validation.StringInSlice([]string{"LIKE", "EXACT"}, false),
				Description:  "The search mode of query. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.",
			},

			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
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
			"private_zones": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						//"id": {
						//	Type:        schema.TypeString,
						//	Computed:    true,
						//	Description: "The id of the private zone.",
						//},
						"zid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the private zone.",
						},
						"zone_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the private zone.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The remark of the private zone.",
						},
						"record_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The record count of the private zone.",
						},
						"recursion_mode": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the recursion mode of the private zone is enabled.",
						},
						"line_mode": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The line mode of the private zone, specified whether the intelligent mode and the load balance function is enabled.",
						},
						"last_operator": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id of the last operator who created the private zone.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The created time of the private zone.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The updated time of the private zone.",
						},
						"region": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The region of the private zone.",
						},
						"bind_vpcs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The Bind vpc info of the private zone.",
							Elem: schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the bind vpc.",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region of the bind vpc.",
									},
									"region_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region name of the bind vpc.",
									},
									"account_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The account id of the bind vpc.",
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

func dataSourceVolcenginePrivateZonesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewPrivateZoneService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcenginePrivateZones())
}
