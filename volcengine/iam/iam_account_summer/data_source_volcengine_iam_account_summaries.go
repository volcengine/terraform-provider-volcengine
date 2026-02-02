package iam_account_summer

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamAccountSummaries() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamAccountSummariesRead,
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
			"account_summaries": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of account summaries.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_keys_per_user_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of access keys per user.",
						},
						"access_keys_per_account_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of access keys per account.",
						},
						"attached_policies_per_group_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of attached policies per group.",
						},
						"attached_system_policies_per_group_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of attached system policies per group.",
						},
						"attached_policies_per_role_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of attached policies per role.",
						},
						"attached_system_policies_per_role_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of attached system policies per role.",
						},
						"attached_system_policies_per_user_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of attached system policies per user.",
						},
						"attached_policies_per_user_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of attached policies per user.",
						},
						"groups_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of groups.",
						},
						"policies_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of policies.",
						},
						"roles_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of roles.",
						},
						"users_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of users.",
						},
						"policy_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of policy.",
						},
						"roles_usage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The usage of roles.",
						},
						"users_usage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The usage of users.",
						},
						"groups_usage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The usage of groups.",
						},
						"policies_usage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The usage of policies.",
						},
						"groups_per_user_quota": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The quota of groups per user.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamAccountSummariesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamAccountSummaryService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineIamAccountSummaries())
}
