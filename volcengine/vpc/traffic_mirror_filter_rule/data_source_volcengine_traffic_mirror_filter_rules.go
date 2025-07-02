package traffic_mirror_filter_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTrafficMirrorFilterRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTrafficMirrorFilterRulesRead,
		Schema: map[string]*schema.Schema{
			"traffic_mirror_filter_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of traffic mirror filter IDs.",
			},
			"traffic_mirror_filter_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of traffic mirror filter names.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of traffic mirror filter.",
			},
			"tags": ve.TagsSchema(),
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
			"traffic_mirror_filter_rules": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of traffic mirror filter rule.",
						},
						"traffic_mirror_filter_rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of traffic mirror filter rule.",
						},
						"traffic_mirror_filter_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of traffic mirror filter.",
						},
						"traffic_direction": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The traffic direction of traffic mirror filter rule.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The priority of traffic mirror filter rule.",
						},
						"policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The policy of traffic mirror filter rule.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol of traffic mirror filter rule.",
						},
						"source_cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source cidr block of traffic mirror filter rule.",
						},
						"source_port_range": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source port range of traffic mirror filter rule.",
						},
						"destination_cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination cidr block of traffic mirror filter rule.",
						},
						"destination_port_range": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination port range of traffic mirror filter rule.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of traffic mirror filter rule.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of traffic mirror filter rule.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of traffic mirror filter rule.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last update time of traffic mirror filter rule.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTrafficMirrorFilterRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTrafficMirrorFilterRuleService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTrafficMirrorFilterRules())
}
