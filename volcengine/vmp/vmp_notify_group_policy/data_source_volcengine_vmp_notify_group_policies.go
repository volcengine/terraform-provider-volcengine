package vmp_notify_group_policy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVmpNotifyGroupPolicies() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVmpNotifyGroupPoliciesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of notify group policy ids.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of notify group policy.",
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
			"notify_policies": {
				Type:        schema.TypeList,
				Description: "The list of notify group policies.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the notify group policy.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of notify group policy.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of notify group policy.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of notify group policy.",
						},
						"levels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The levels of the notify group policy.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"level": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The level of the policy.",
									},
									"group_by": {
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "The aggregate dimension.",
									},
									"group_wait": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The wait time.",
									},
									"group_interval": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The aggregation cycle.",
									},
									"repeat_interval": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The notification cycle.",
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

func dataSourceVolcengineVmpNotifyGroupPoliciesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineVmpNotifyGroupPolicies())
}
