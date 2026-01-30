package rule_bound_host_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsRuleBoundHostGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsRuleBoundHostGroupsRead,
		Schema: map[string]*schema.Schema{
			"rule_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the rule.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"host_groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of Host Group query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the host group.",
						},
						"host_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the host group.",
						},
						"host_group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the host group.",
						},
						"host_identifier": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The identifier of the host.",
						},
						"auto_update": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable auto update.",
						},
						"update_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The start time of auto update.",
						},
						"update_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The end time of auto update.",
						},
						"service_logging": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable service logging.",
						},
						"iam_project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the iam project.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the host group.",
						},
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modification time of the host group.",
						},
					},
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
		},
	}
}

func dataSourceVolcengineTlsRuleBoundHostGroupsRead(d *schema.ResourceData, meta interface{}) error {
	TlsRuleBoundHostGroupService := NewTlsRuleBoundHostGroupService(meta.(*ve.SdkClient))
	return (&ve.Dispatcher{}).Data(TlsRuleBoundHostGroupService, d, DataSourceVolcengineTlsRuleBoundHostGroups())
}
