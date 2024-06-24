package kafka_topic

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKafkaTopics() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKafkaTopicsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of kafka instance.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of kafka topic. This field supports fuzzy query.",
			},
			"partition_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of partition in kafka topic.",
			},
			"replica_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of replica in kafka topic.",
			},
			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "When a user name is specified, only the access policy of the specified user for this Topic will be returned.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of kafka topic.",
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
			"topics": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the kafka topic.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the kafka topic.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the kafka topic.",
						},
						"partition_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of partition in the kafka topic.",
						},
						"replica_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of replica in the kafka topic.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the kafka topic.",
						},
						"parameters": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The parameters of the kafka topic.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"min_insync_replica_number": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The min number of sync replica.",
									},
									"message_max_byte": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The max byte of message.",
									},
									"log_retention_hours": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The retention hours of log.",
									},
								},
							},
						},
						"all_authority": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the kafka topic is configured to be accessible by all users.",
						},
						"access_policies": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The access policies info of the kafka topic.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of SASL user.",
									},
									"access_policy": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The access policy of SASL user.",
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

func dataSourceVolcengineKafkaTopicsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKafkaTopicService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineKafkaTopics())
}
