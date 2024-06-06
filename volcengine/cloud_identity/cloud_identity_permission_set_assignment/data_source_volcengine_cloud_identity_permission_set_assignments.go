package cloud_identity_permission_set_assignment

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudIdentityPermissionSetAssignments() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudIdentityPermissionSetAssignmentsRead,
		Schema: map[string]*schema.Schema{
			"permission_set_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of cloud identity permission set.",
			},
			"target_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The target account id of cloud identity permission set assignment.",
			},
			"principal_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The principal type of cloud identity permission set. Valid values: `User`, `Group`.",
			},
			"principal_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The principal id of cloud identity permission set. When the `principal_type` is `User`, this field is specified to `UserId`. When the `principal_type` is `Group`, this field is specified to `GroupId`.",
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
			"assignments": {
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
						"permission_set_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the cloud identity permission set.",
						},
						"target_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The target account id of the cloud identity permission set assignment.",
						},
						"principal_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The principal type of the cloud identity permission set assignment.",
						},
						"principal_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The principal id of the cloud identity permission set assignment.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the cloud identity permission set assignment.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCloudIdentityPermissionSetAssignmentsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCloudIdentityPermissionSetAssignmentService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCloudIdentityPermissionSetAssignments())
}
