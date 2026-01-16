package alb_acl

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAlbAcls() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAlbAclsRead,
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
			"acl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of acl.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of project.",
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
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of Acl.",
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
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of Acl.",
						},
						"tags": ve.TagsSchemaComputed(),
						"listeners": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The listeners of acl.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of acl.",
									},
									"listener_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of Listener.",
									},
									"listener_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Name of Listener.",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The port info of listener.",
									},
									"protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The protocol info of listener.",
									},
								},
							},
						},
						"acl_entries": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The entries info of acl.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of entry.",
									},
									"entry": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The info of entry.",
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

func dataSourceVolcengineAlbAclsRead(d *schema.ResourceData, meta interface{}) error {
	aclService := NewAclService(meta.(*ve.SdkClient))
	return aclService.Dispatcher.Data(aclService, d, DataSourceVolcengineAlbAcls())
}
