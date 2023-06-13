package shard

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsShards() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsShardsRead,
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of topic.",
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
			"shards": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of topic.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of shard.",
						},
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modify time.",
						},
						"stop_write_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The stop write time.",
						},
						"shard_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of shard.",
						},
						"inclusive_begin_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The begin key info.",
						},
						"exclusive_end_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The end key info.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsShardsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsShards())
}
