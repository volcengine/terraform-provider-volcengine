package hot_key

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineHotKeys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineHotKeysRead,
		Schema: map[string]*schema.Schema{
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
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of Instance.",
			},
			"query_start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Query the start time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).",
			},
			"query_end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Query the end time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).",
			},
			"key_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the data type used to filter the query results of hot keys.",
			},
			"shard_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specify the list of shard ids used to filter the query results of hot keys.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hot_key": {
				Description: "The List of hot Key details.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the database to which the hot Key belongs.",
						},
						"node_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The node ID to which the hot Key belongs.",
						},
						"key_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the hot Key.",
						},
						"key_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of hot Key.",
						},
						"shard_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The shard ID to which the hot Key belongs.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The collection time of the hot Key, in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).",
						},
						"query_count": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The number of accesses to the hot Key.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineHotKeysRead(d *schema.ResourceData, meta interface{}) error {
	service := NewHotKeyService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineHotKeys())
}
