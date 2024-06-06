package cloud_identity_permission_set

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudIdentityPermissionSets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudIdentityPermissionSetsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of cloud identity permission set IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of cloud identity permission set.",
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
			"permission_sets": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the cloud identity permission set.",
						},
						"permission_set_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the cloud identity permission set.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the cloud identity permission set.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the cloud identity permission set.",
						},
						"relay_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The relay state of the cloud identity permission set.",
						},
						"session_duration": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The session duration of the cloud identity permission set.",
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the cloud identity permission set.",
						},
						"updated_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The updated time of the cloud identity permission set.",
						},
						"permission_policies": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The policies of the cloud identity permission set.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"permission_policy_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the cloud identity permission set policy.",
									},
									"permission_policy_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the cloud identity permission set policy.",
									},
									"permission_policy_document": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The document of the cloud identity permission set policy.",
									},
									"create_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The create time of the cloud identity permission set policy.",
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

func dataSourceVolcengineCloudIdentityPermissionSetsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCloudIdentityPermissionSetService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCloudIdentityPermissionSets())
}
