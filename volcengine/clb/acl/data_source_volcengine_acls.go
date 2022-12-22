package acl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAcls() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAclsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Acl IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Acl.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ProjectName of Acl.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Acl query.",
			},
			"acl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of acl.",
			},
			"acls": {
				Description: "The collection of Acl query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Acl.",
						},
						"acl_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Acl.",
						},
						"acl_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Name of Acl.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of Acl.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time of Acl.",
						},
						"acl_entry_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of acl entry.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of Acl.",
						},
						"listeners": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "The listeners of Acl.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ProjectName of Acl.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineAclsRead(d *schema.ResourceData, meta interface{}) error {
	aclService := NewAclService(meta.(*ve.SdkClient))
	return aclService.Dispatcher.Data(aclService, d, DataSourceVolcengineAcls())
}
