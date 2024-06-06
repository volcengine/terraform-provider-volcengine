package cloud_identity_user_provisioning

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudIdentityUserProvisionings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudIdentityUserProvisioningsRead,
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The account id.",
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
			"user_provisionings": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the cloud identity user provisioning.",
						},
						"user_provisioning_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the cloud identity user provisioning.",
						},
						"principal_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The principal type of the cloud identity user provisioning.",
						},
						"principal_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The principal id of the cloud identity user provisioning.",
						},
						"principal_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The principal name of the cloud identity user provisioning.",
						},
						"target_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The target account id of the cloud identity user provisioning.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the cloud identity user provisioning.",
						},
						"identity_source_strategy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The identity source strategy of the cloud identity user provisioning.",
						},
						"duplication_strategy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The duplication strategy of the cloud identity user provisioning.",
						},
						"duplication_suffix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The duplication suffix of the cloud identity user provisioning.",
						},
						"deletion_strategy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deletion strategy of the cloud identity user provisioning.",
						},
						"provision_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the cloud identity user provisioning.",
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The created time of the cloud identity user provisioning.",
						},
						"updated_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The updated time of the cloud identity user provisioning.",
						},
						"department_names": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The department names of the cloud identity user provisioning.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCloudIdentityUserProvisioningsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCloudIdentityUserProvisioningService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCloudIdentityUserProvisionings())
}
