package iam_caller_identity

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineIamCallerIdentities() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineIamCallerIdentitiesRead,
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
			"caller_identities": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of caller identities.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id.",
						},
						"trn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The trn.",
						},
						"identity_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The identity type.",
						},
						"identity_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The identity id.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineIamCallerIdentitiesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamCallerIdentityService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineIamCallerIdentities())
}
