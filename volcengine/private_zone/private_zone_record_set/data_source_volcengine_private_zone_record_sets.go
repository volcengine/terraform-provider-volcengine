package private_zone_record_set

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcenginePrivateZoneRecordSets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcenginePrivateZoneRecordSetsRead,
		Schema: map[string]*schema.Schema{
			"zid": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The zid of Private Zone.",
			},
			"record_set_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of Private Zone Record Set.",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The host of Private Zone Record Set.",
			},
			"search_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "LIKE",
				ValidateFunc: validation.StringInSlice([]string{"LIKE", "EXACT"}, false),
				Description:  "The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.",
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
			"record_sets": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"record_set_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the private zone record set.",
						},
						"fqdn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Complete domain name of the private zone record.",
						},
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The host of the private zone record.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the private zone record.",
						},
						"line": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet id of the private zone record. This field is only effected when the `intelligent_mode` of the private zone is true.",
						},
						"weight_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable the load balance of the private zone record set.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcenginePrivateZoneRecordSetsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewPrivateZoneRecordSetService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcenginePrivateZoneRecordSets())
}
