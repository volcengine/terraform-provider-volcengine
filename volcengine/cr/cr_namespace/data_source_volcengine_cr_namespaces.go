package cr_namespace

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCrNamespaces() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCrNamespacesRead,
		Schema: map[string]*schema.Schema{
			"registry": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The target cr instance name.",
			},
			"names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list of instance IDs.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of instance query.",
			},
			"namespaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of namespaces query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of OCI repository.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when namespace created.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCrNamespacesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCrNamespaceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineCrNamespaces())
}
