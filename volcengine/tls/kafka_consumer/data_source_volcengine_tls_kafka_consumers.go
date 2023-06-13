package kafka_consumer

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsKafkaConsumers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsKafkaConsumerRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of topic IDs.",
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
			"data": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Topic.",
						},
						"allow_consume": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether allow consume.",
						},
						"consume_topic": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The topic of consume.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsKafkaConsumerRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsKafkaConsumers())
}
