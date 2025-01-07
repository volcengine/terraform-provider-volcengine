package rocketmq_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRocketmqGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRocketmqGroupsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of rocketmq instance.",
			},
			"group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of rocketmq group. This field support fuzzy query.",
			},
			"group_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of rocketmq group. Valid values: `TCP`.",
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

			"rocketmq_groups": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the rocketmq group.",
						},
						"group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the rocketmq group.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the rocketmq group.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the rocketmq group.",
						},
						"is_sub_same": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the subscription relationship of consumer instance groups within the group is consistent.",
						},
						"message_delay_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The message delay time of the rocketmq group. The unit is milliseconds.",
						},
						"message_model": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The message model of the rocketmq group.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the rocketmq group.",
						},
						"total_consume_rate": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The total consume rate of the rocketmq group. The unit is per second.",
						},
						"total_diff": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total amount of unconsumed messages.",
						},
						"consumed_topics": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The consumed topic information of the rocketmq group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"topic_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the rocketmq topic.",
									},
									"queue_num": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The queue number of the rocketmq topic.",
									},
									"sub_string": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The sub string of the rocketmq topic.",
									},
								},
							},
						},
						"consumed_clients": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The consumed topic information of the rocketmq group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the consumed client.",
									},
									"client_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The address of the consumed client.",
									},
									"language": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The language of the consumed client.",
									},
									"version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The version of the consumed client.",
									},
									"diff": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The amount of message.",
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

func dataSourceVolcengineRocketmqGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRocketmqGroupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRocketmqGroups())
}
