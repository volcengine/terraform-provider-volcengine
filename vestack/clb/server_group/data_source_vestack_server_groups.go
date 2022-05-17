package server_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

func DataSourceVestackServerGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVestackServerGroupsRead,
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
					},
				},
			},
		},
	}
}

func dataSourceVestackServerGroupsRead(d *schema.ResourceData, meta interface{}) error {
	serverGroupService := NewServerGroupService(meta.(*ve.SdkClient))
	return serverGroupService.Dispatcher.Data(serverGroupService, d, DataSourceVestackServerGroups())
}
