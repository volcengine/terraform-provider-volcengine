package cloud_monitor_contact_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudMonitorContactGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudMonitorContactGroupsRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search for keywords in contact group names, supports fuzzy search.",
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
			"groups": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the contact group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the contact group.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the account.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the contact group.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time.",
						},
						"contacts": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Contact information in the contact group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the contact.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of contact.",
									},
									"phone": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The phone of contact.",
									},
									"email": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The email of contact.",
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

func dataSourceVolcengineCloudMonitorContactGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCloudMonitorContactGroupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCloudMonitorContactGroups())
}
