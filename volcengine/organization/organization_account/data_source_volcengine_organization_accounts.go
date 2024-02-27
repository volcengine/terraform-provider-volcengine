package organization_account

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineOrganizationAccounts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineOrganizationAccountsRead,
		Schema: map[string]*schema.Schema{
			"search": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id or the show name of the account. This field supports fuzzy query.",
			},
			"org_unit_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the organization unit.",
			},
			"verification_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the verification.",
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
			"accounts": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the account.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the account.",
						},
						"account_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the account.",
						},
						"show_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The show name of the account.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the account.",
						},
						"owner": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The owner id of the account.",
						},
						"org_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the organization.",
						},
						"org_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The type of the organization. `1` means business organization.",
						},
						"org_unit_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the organization unit.",
						},
						"org_unit_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the organization unit.",
						},
						"org_verification_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the organization verification.",
						},
						"iam_role": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the iam role.",
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The created time of the account.",
						},
						"updated_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The updated time of the account.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deleted time of the account.",
						},
						"delete_uk": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The delete uk of the account.",
						},
						"join_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The join type of the account. `0` means create, `1` means invitation.",
						},
						"allow_exit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether to allow exit the organization. `0` means allowed, `1` means not allowed.",
						},
						"allow_console": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether to allow the account enable console. `0` means allowed, `1` means not allowed.",
						},
						"is_owner": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether the account is owner. `0` means not owner, `1` means owner.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineOrganizationAccountsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewOrganizationAccountService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineOrganizationAccounts())
}
