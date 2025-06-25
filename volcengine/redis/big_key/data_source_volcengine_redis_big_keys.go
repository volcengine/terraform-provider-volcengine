package big_key

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineBigKeys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineBigKeysRead,
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
			"order_by": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the sorting conditions of the query results.",
			},
			"big_key": {
				Description: "Details of the big Key.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the database to which the big Key belongs.",
						},
						"key_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the big Key.",
						},
						"key_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of big Key.",
						},
						"value_len": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The number of elements contained in the large Key.",
						},
						"value_size": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The memory usage of large keys, unit: Byte.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineBigKeysRead(d *schema.ResourceData, meta interface{}) error {
	service := NewBigKeyService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineBigKeys())
}
