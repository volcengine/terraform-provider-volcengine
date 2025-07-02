package traffic_mirror_target

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTrafficMirrorTargets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTrafficMirrorTargetsRead,
		Schema: map[string]*schema.Schema{
			"traffic_mirror_target_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of traffic mirror target IDs.",
			},
			"traffic_mirror_target_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of traffic mirror target.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of traffic mirror target.",
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
			"traffic_mirror_targets": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of traffic mirror target.",
						},
						"traffic_mirror_target_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of traffic mirror target.",
						},
						"traffic_mirror_target_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of traffic mirror target.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of traffic mirror target.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance type of traffic mirror target.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance id of traffic mirror target.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of traffic mirror target.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of traffic mirror target.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of traffic mirror target.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of traffic mirror target.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTrafficMirrorTargetsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTrafficMirrorTargetService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTrafficMirrorTargets())
}
