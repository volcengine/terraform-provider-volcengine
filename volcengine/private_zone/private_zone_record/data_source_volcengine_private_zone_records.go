package private_zone_record

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcenginePrivateZoneRecords() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcenginePrivateZoneRecordsRead,
		Schema: map[string]*schema.Schema{
			"zid": {
				Type:         schema.TypeInt,
				Optional:     true,
				AtLeastOneOf: []string{"zid", "record_ids"},
				Description:  "The zid of Private Zone.",
			},
			"record_id": {
				Type:     schema.TypeString,
				Optional: true,
				//AtLeastOneOf: []string{"zid", "record_id"},
				Deprecated:  "This field is deprecated, please use `record_ids` instead.",
				Description: "The id of Private Zone Record.",
			},
			"record_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				AtLeastOneOf: []string{"zid", "record_ids"},
				Description:  "The ids of Private Zone Record.",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The host of Private Zone Record.",
			},
			"search_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "LIKE",
				ValidateFunc: validation.StringInSlice([]string{"LIKE", "EXACT"}, false),
				Description:  "The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The domain name of Private Zone Record.",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of Private Zone Record.",
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The value of Private Zone Record.",
			},
			"last_operator": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The last operator account id of Private Zone Record.",
			},
			"line": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The subnet id of Private Zone Record.",
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
			"records": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"record_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the private zone record.",
						},
						"zid": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The zid of the private zone record.",
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
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value of the private zone record.",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ttl of the private zone record. Unit: second.",
						},
						"line": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subnet id of the private zone record. This field is only effected when the `intelligent_mode` of the private zone is true.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The weight of the private zone record.",
						},
						"enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the private zone record is enabling.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The remark of the private zone record.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The created time of the private zone record.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The updated time of the private zone record.",
						},
						"last_operator": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last operator account id of the private zone record.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcenginePrivateZoneRecordsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewPrivateZoneRecordService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcenginePrivateZoneRecords())
}
