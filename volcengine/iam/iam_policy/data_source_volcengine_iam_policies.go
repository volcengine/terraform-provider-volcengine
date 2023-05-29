package iam_policy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamPolicies() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamPoliciesRead,
		Schema: map[string]*schema.Schema{
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Policy.",
			},
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
			"query": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Query policies, support policy name or description.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of policy.",
			},
			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the IAM user.",
			},
			"role_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the IAM role.",
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
						"user_attach_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user attach time of the Policy.The data show only query with user_name.",
						},
						"user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the IAM user.The data show only query with user_name.",
						},
						"role_attach_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The role attach time of the Policy.The data show only query with role_name.",
						},
						"role_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the IAM role.The data show only query with role_name.",
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
