package consumer_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineConsumerGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineConsumerGroupsRead,
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
			"consumer_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the consumer group.",
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The log project ID to which the consumption group belongs.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the log item to which the consumption group belongs.",
			},
			"topic_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The log topic ID to which the consumer belongs.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the log topic to which the consumption group belongs.",
			},
			"iam_project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IAM log project name.",
			},
			"consumer_groups": {
				Description: "List of log service consumption groups.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic_id": {
							Computed: true,
							Type:     schema.TypeSet,
							Set:      schema.HashString,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The list of log topic ids to be consumed by the consumer group.",
						},
						"project_id": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The log project ID to which the consumption group belongs.",
						},
						"project_name": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The name of the log item to which the consumption group belongs.",
						},
						"heartbeat_ttl": {
							Computed:    true,
							Type:        schema.TypeInt,
							Description: "The time of heart rate expiration, measured in seconds, has a value range of 1 to 300.",
						},
						"ordered_consume": {
							Computed:    true,
							Type:        schema.TypeBool,
							Description: "Whether to consume in sequence.",
						},
						"consumer_group_name": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The name of the consumer group.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineConsumerGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewConsumerGroupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineConsumerGroups())
}
