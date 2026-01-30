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
			"tags": ve.TagsSchema(),
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
						"service_managed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the Acl is managed by service.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the Acl.",
						},
						"tags": ve.TagsSchemaComputed(),
						"acl_entries": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The acl entry list of the Acl.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the AclEntry.",
									},
									"entry": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The address range of the IP entry.",
									},
								},
							},
						},
						// 对应 DescribeAclAttributes API 的响应参数中的 Listeners 结构体
						"listener_details": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The listener details of the Acl.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"listener_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the listener.",
									},
									"listener_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the listener.",
									},
									"protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The protocol of the listener.",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The port receiving request of the listener.",
									},
									"acl_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The control method of the listener for this Acl. Valid values: `black`, `white`.",
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

func dataSourceVolcengineAclsRead(d *schema.ResourceData, meta interface{}) error {
	aclService := NewAclService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(aclService, d, DataSourceVolcengineAcls())
}
