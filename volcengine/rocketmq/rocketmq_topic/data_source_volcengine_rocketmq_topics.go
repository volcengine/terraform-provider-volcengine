package rocketmq_topic

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRocketmqTopics() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRocketmqTopicsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of rocketmq instance.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the rocketmq topic. This field support fuzzy query.",
			},
			"message_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The type of the rocketmq message. Setting this parameter means filtering the Topic list based on the specified message type. The value explanation is as follows:\n0: Regular message\n1: Transaction message\n2: Partition order message\n3: Global sequential message\n4: Delay message.",
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

			"rocketmq_topics": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of rocketmq instance.",
						},
						"topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the rocketmq topic.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the rocketmq topic.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the rocketmq topic.",
						},
						"message_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The type of the rocketmq message.",
						},
						"queue_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of the rocketmq topic queue.",
						},
						"access_policies": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The access policies of the rocketmq topic.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The access key of the rocketmq key.",
									},
									"authority": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The authority of the rocketmq key for the current topic.",
									},
								},
							},
						},
						"queues": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The queues information of the rocketmq topic.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"queue_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the rocketmq queue.",
									},
									"start_offset": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The start offset of the rocketmq queue.",
									},
									"end_offset": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The end offset of the rocketmq queue.",
									},
									"message_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The message count of the rocketmq queue.",
									},
									"last_update_timestamp": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The last update timestamp of the rocketmq queue.",
									},
								},
							},
						},
						"groups": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The groups information of the rocketmq topic.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the rocketmq group.",
									},
									"message_model": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The message model of the rocketmq group.",
									},
									"sub_string": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The sub string of the rocketmq group.",
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

func dataSourceVolcengineRocketmqTopicsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRocketmqTopicService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRocketmqTopics())
}
