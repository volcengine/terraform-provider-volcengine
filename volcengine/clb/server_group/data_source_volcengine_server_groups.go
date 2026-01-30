package server_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineServerGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineServerGroupsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of ServerGroup IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of ServerGroup.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of ServerGroup query.",
			},
			"load_balancer_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the Clb.",
			},
			"server_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the ServerGroup.",
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"instance", "ip"}, false),
				Description:  "The type of ServerGroup. Valid values: `instance`, `ip`.",
			},
			"tags": ve.TagsSchema(),

			"groups": {
				Description: "The collection of ServerGroup query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true, // tf中不支持写值
							Description: "The ID of the ServerGroup.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the ServerGroup.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the ServerGroup.",
						},
						"server_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the ServerGroup.",
						},
						"server_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the ServerGroup.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the ServerGroup.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the ServerGroup.",
						},
						"address_ip_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The address IP version of the ServerGroup.",
						},
						"any_port_enabled": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether full port forwarding is enabled.",
						},
						"tags": ve.TagsSchemaComputed(),
						// DescribeServerGroupAttributes 详情API返回
						"load_balancer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the LoadBalancer.",
						},
						"listeners": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The listeners of the ServerGroup.",
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

func dataSourceVolcengineServerGroupsRead(d *schema.ResourceData, meta interface{}) error {
	serverGroupService := NewServerGroupService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(serverGroupService, d, DataSourceVolcengineServerGroups())
}
