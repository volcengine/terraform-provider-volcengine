package organization_unit

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineOrganizationUnits() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineOrganizationUnitsRead,
		Schema: map[string]*schema.Schema{
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
			"units": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the organization unit.",
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The created time of the organization unit.",
						},
						"updated_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The updated time of the organization unit.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deleted time of the organization unit.",
						},
						"owner": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The owner of the organization unit.",
						},
						"org_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the organization.",
						},
						"org_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The organization type.",
						},
						"parent_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parent Unit ID.",
						},
						"depth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The depth of the organization unit.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the organization unit.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the organization unit.",
						},
						"delete_uk": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Delete marker.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineOrganizationUnitsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewOrganizationUnitService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineOrganizationUnits())
}
