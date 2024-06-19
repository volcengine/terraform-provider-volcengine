package kafka_topic_partition

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKafkaTopicPartitions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKafkaTopicPartitionsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of kafka instance.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of kafka topic.",
			},
			"under_insync_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to only query the list of partitions that have out-of-sync replicas, the default value is false.",
			},
			"partition_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashInt,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "The index number of partition.",
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
			"partitions": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"partition_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The index number of partition.",
						},
						"leader": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The leader info of partition.",
						},
						"start_offset": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The start offset of partition leader.",
						},
						"end_offset": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The end offset of partition leader.",
						},
						"message_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of message.",
						},
						"replicas": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Computed:    true,
							Description: "The replica info.",
						},
						"insync_replicas": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Computed:    true,
							Description: "The insync replica info.",
						},
						"under_insync_replicas": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Computed:    true,
							Description: "The under insync replica info.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineKafkaTopicPartitionsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKafkaTopicPartitionService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineKafkaTopicPartitions())
}
