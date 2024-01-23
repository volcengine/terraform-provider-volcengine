package organization

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineOrganizations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineOrganizationsRead,
		Schema: map[string]*schema.Schema{
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
			"organizations": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the organization.",
						},
						"owner": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The owner id of the organization.",
						},
						"type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The type of the organization.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the organization.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the organization.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The status of the organization.",
						},
						"delete_uk": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The delete uk of the organization.",
						},
						"account_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The account id of the organization owner.",
						},
						"account_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account name of the organization owner.",
						},
						"main_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The main name of the organization owner.",
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The created time of the organization.",
						},
						"updated_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The updated time of the organization.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deleted time of the organization.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineOrganizationsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewOrganizationService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineOrganizations())
}
