package iam_user_policy_attachment

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamUserPolicyAttachments() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamUserPolicyAttachmentsRead,
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the user.",
			},
			"policies": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of policies.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the policy.",
						},
						"policy_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the policy.",
						},
						"policy_trn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The trn of the policy.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the policy.",
						},
						"attach_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The attach date of the policy.",
						},
						"policy_scope": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The scope of the policy.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"policy_scope_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the policy scope.",
									},
									"project_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the project.",
									},
									"project_display_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The display name of the project.",
									},
									"attach_date": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The attach date of the policy scope.",
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

func dataSourceVolcengineIamUserPolicyAttachmentsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamUserPolicyAttachmentService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineIamUserPolicyAttachments())
}
