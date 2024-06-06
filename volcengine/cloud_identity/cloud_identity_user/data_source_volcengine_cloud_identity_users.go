package cloud_identity_user

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudIdentityUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudIdentityUsersRead,
		Schema: map[string]*schema.Schema{
			"department_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The department id.",
			},
			"user_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"display_name"},
				Description:   "The name of cloud identity user.",
			},
			"display_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"user_name"},
				Description:   "The display name of cloud identity user.",
			},
			"source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The source of cloud identity user. Valid values: `Sync`, `Manual`.",
			},
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
			"users": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the cloud identity user.",
						},
						"user_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the cloud identity user.",
						},
						"user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the cloud identity user.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The display name of the cloud identity user.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the cloud identity user.",
						},
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source of the cloud identity user.",
						},
						"email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The email of the cloud identity user.",
						},
						"phone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The phone of the cloud identity user.",
						},
						"identity_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The identity type of the cloud identity user.",
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The created time of the cloud identity user.",
						},
						"updated_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The updated time of the cloud identity user.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCloudIdentityUsersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCloudIdentityUserService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCloudIdentityUsers())
}
