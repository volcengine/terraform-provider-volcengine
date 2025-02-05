package account

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineMongoDBAccounts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAccountsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Target query mongodb instance id.",
			},
			"account_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of account. This field support fuzzy query.",
			},
			"auth_db": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The database of account. This field support fuzzy query.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of accounts query.",
			},
			"accounts": {
				Description: "The collection of accounts query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of account.",
						},
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of account.",
						},
						"auth_db": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The database of account.",
						},
						"account_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of account.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of account.",
						},
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modify time of account.",
						},
						"account_privileges": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The privilege info of mongo instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"db_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Name of DB.",
									},
									"role_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Name of role.",
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

func dataSourceVolcengineAccountsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewMongoDBAccountService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineMongoDBAccounts())
}
