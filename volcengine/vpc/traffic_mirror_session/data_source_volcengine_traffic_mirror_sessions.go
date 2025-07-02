package traffic_mirror_session

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTrafficMirrorSessions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTrafficMirrorSessionsRead,
		Schema: map[string]*schema.Schema{
			"traffic_mirror_session_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of traffic mirror session IDs.",
			},
			"traffic_mirror_session_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of traffic mirror session names.",
			},
			"network_interface_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of network interface.",
			},
			"traffic_mirror_target_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of traffic mirror target.",
			},
			"traffic_mirror_filter_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of traffic mirror filter.",
			},
			"virtual_network_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The ID of virtual network.",
			},
			"packet_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The packet length of traffic mirror session.",
			},
			"priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The priority of traffic mirror session.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of traffic mirror session.",
			},
			"tags": ve.TagsSchema(),
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
			"traffic_mirror_sessions": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of traffic mirror session.",
						},
						"traffic_mirror_session_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of traffic mirror session.",
						},
						"traffic_mirror_session_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of traffic mirror session.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of traffic mirror session.",
						},
						"traffic_mirror_target_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of traffic mirror target.",
						},
						"traffic_mirror_filter_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of traffic mirror filter.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of traffic mirror session.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The priority of traffic mirror session.",
						},
						"packet_length": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The packet length of traffic mirror session.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The business status of traffic mirror session.",
						},
						"lock_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The lock reason of traffic mirror session.",
						},
						"virtual_network_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of virtual network.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of traffic mirror session.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of traffic mirror session.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of traffic mirror session.",
						},
						"tags": ve.TagsSchemaComputed(),
						"traffic_mirror_source_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The IDs of traffic mirror source.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTrafficMirrorSessionsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTrafficMirrorSessionService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTrafficMirrorSessions())
}
