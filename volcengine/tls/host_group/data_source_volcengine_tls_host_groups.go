package host_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsHostGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsHostGroupRead,
		Schema: map[string]*schema.Schema{
			"host_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of host group.",
			},
			"host_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of host group.",
			},
			"host_identifier": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The identifier of host.",
			},
			"iam_project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of iam.",
			},
			"auto_update": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether enable auto update.",
			},
			"service_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether enable service logging.",
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
			"infos": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of host group.",
						},
						"host_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of host group.",
						},
						"host_group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of host group.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of host group.",
						},
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modify time of host group.",
						},
						"host_identifier": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The identifier of host.",
						},
						"host_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of host.",
						},
						"rule_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The rule count of host.",
						},
						"normal_heartbeat_status_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The normal heartbeat status count of host.",
						},
						"abnormal_heartbeat_status_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The abnormal heartbeat status count of host.",
						},
						"iam_project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of iam.",
						},
						"update_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update start time of log collector.",
						},
						"update_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update end time of log collector.",
						},
						"agent_latest_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The latest version of log collector.",
						},
						"auto_update": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether enable auto update.",
						},
						"service_logging": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether enable service logging.",
						},
						"host_ip_list": {
							Type:        schema.TypeSet,
							Computed:    true,
							Set:         schema.HashString,
							Description: "The ip list of host group.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsHostGroupRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsHostGroups())
}
