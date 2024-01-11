package cloud_monitor_event_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudMonitorEventRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudMonitorEventRulesRead,
		Schema: map[string]*schema.Schema{
			"rule_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Rule name, search rules by name using fuzzy search.",
			},
			"source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Event source.",
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
			"rules": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the rule.",
						},
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the rule.",
						},
						"event_source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source of the event.",
						},
						"rule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the rule.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the rule.",
						},
						"event_bus_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the event bus.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the account.",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the region.",
						},
						"level": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The level of the rule.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Enable the state of the rule.",
						},
						// resource effective_time
						"effect_start_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The start time of the rule.",
						},
						// resource effective_time
						"effect_end_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The end time of the rule.",
						},
						"event_type": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The event type.",
						},
						"filter_pattern": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Filter mode, also known as event matching rules. Custom matching rules are not currently supported.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Computed:    true,
										Type:        schema.TypeList,
										Description: "The list of corresponding event types in pattern matching, currently set to match any.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"source": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Event source corresponding to pattern matching.",
									},
								},
							},
						},
						"contact_methods": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of contact methods.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "When the alarm notification method is alarm callback, it triggers the callback address.",
						},
						"tls_target": {
							Computed:    true,
							Type:        schema.TypeSet,
							Description: "The alarm method for log service triggers the configuration of the log service.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region_name_en": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The English region name.",
									},
									"region_name_cn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Chinese region name.",
									},
									"project_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The project name.",
									},
									"project_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The project id.",
									},
									"topic_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The topic id.",
									},
									"topic_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The topic name.",
									},
								},
							},
						},
						"message_queue": {
							Computed:    true,
							Type:        schema.TypeSet,
							Description: "The triggered message queue when the alarm notification method is Kafka message queue.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The kafka instance id.",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region.",
									},
									"topic": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The topic name.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The message queue type, only support kafka now.",
									},
									"vpc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The vpc id.",
									},
								},
							},
						},
						"contact_group_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "When the alarm notification method is phone, SMS, or email, the triggered alarm contact group ID.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"created_at": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The create time.",
						},
						"updated_at": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The updated time.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCloudMonitorEventRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCloudMonitorEventRuleService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCloudMonitorEventRules())
}
