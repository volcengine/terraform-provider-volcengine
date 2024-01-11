package cloud_monitor_contact

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudMonitorContacts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudMonitorContactsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Contact IDs.",
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
			"contacts": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of contact.",
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
	}
}

func dataSourceVolcengineCloudMonitorContactsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineCloudMonitorContacts())
}
