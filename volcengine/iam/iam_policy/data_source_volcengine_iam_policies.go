package iam_policy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamPolicies() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamPoliciesRead,
		Schema: map[string]*schema.Schema{
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Policy query.",
			},
			"scope": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The scope of the Policy.",
			},
			"with_service_role_policy": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Whether to return the service role policy.",
			},
			"policies": {
				Description: "The collection of Policy query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Policy.",
						},
						"policy_trn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource name of the Policy.",
						},
						"policy_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Policy.",
						},
						"policy_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the Policy.",
						},
						"create_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the Policy.",
						},
						"update_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the Policy.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the Policy.",
						},
						"policy_document": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The document of the Policy.",
						},
						"category": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The category of the Policy.",
						},
						"attachment_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The attachment count of the Policy.",
						},
						"is_service_role_policy": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether the Policy is a service role policy.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamPoliciesRead(d *schema.ResourceData, meta interface{}) error {
	iamPolicyService := NewIamPolicyService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(iamPolicyService, d, DataSourceVolcengineIamPolicies())
}
