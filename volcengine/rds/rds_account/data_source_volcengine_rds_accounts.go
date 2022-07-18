package rds_account

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsAccounts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsAccountsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of database account IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of database account.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of database account query.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the RDS instance.",
			},
			"account_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the database account.",
			},
			"rds_accounts": {
				Description: "The collection of RDS instance account query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true, // tf中不支持写值
							Description: "The ID of the RDS instance account.",
						},
						"account_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the database account.",
						},
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the database account.",
						},
						"account_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the database account.",
						},
						"db_privileges": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The privilege detail list of RDS instance account.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"db_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of database.",
									},
									"account_privilege": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The privilege type of the account.",
									},
									"account_privilege_str": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The privilege string of the account.",
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

func dataSourceVolcengineRdsAccountsRead(d *schema.ResourceData, meta interface{}) error {
	rdsAccountService := NewRdsAccountService(meta.(*volc.SdkClient))
	return rdsAccountService.Dispatcher.Data(rdsAccountService, d, DataSourceVolcengineRdsAccounts())
}
