package vke_permission

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVkePermissions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVkePermissionsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of RBAC Permission IDs.",
			},
			"role_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of RBAC Role Names.",
			},
			"cluster_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Cluster IDs.",
			},
			"namespaces": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Namespaces.",
			},
			"grantee_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:         schema.HashInt,
				Description: "A list of Grantee IDs.",
			},
			"grantee_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of Grantee. Valid values: `User`, `Role`.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of RBAC Permission.",
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
			"access_policies": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the RBAC Permission.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the RBAC Permission.",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The message of the RBAC Permission.",
						},
						"role_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the RBAC Role.",
						},
						"cluster_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Cluster.",
						},
						"namespace": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Namespace of the RBAC Permission.",
						},
						"grantee_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the Grantee.",
						},
						"granted_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The granted time of the RBAC Permission.",
						},
						"revoked_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The revoked time of the RBAC Permission.",
						},
						"grantee_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the Grantee.",
						},
						"authorizer_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the Authorizer.",
						},
						"authorizer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Authorizer.",
						},
						"authorizer_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the Authorizer.",
						},
						"authorized_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The authorized time of the RBAC Permission.",
						},
						"is_custom_role": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the RBAC Role is custom role.",
						},
						"kube_role_binding_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Kube Role Binding.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVkePermissionsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVkePermissionService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVkePermissions())
}
