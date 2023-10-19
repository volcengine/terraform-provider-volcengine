package iam_user_group_policy_attachment

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamUserGroupPolicyAttachments() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamUserGroupPolicyAttachmentsRead,
		Schema: map[string]*schema.Schema{
			"user_group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A name of user group.",
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
			"policies": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_trn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource name of the strategy.",
						},
						"policy_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the policy.",
						},
						"policy_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the policy.",
						},
						"attach_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Attached time.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamUserGroupPolicyAttachmentsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamUserGroupPolicyAttachmentService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineIamUserGroupPolicyAttachments())
}
