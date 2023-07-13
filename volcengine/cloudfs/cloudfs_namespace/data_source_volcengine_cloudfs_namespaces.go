package cloudfs_namespace

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudfsNamespaces() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudfsNamespacesRead,
		Schema: map[string]*schema.Schema{
			"fs_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of file system.",
			},
			"ns_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of namespace.",
			},
			"tos_bucket": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of tos bucket.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of cloudfs.",
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
			"namespaces": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the namespace.",
						},
						"tos_bucket": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the tos bucket.",
						},
						"tos_prefix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The tos prefix.",
						},
						"read_only": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the namespace is read-only.",
						},
						"is_my_bucket": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the tos bucket is your own bucket.",
						},
						"service_managed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the namespace is the official service for volcengine.",
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the namespace.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the namespace.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCloudfsNamespacesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCloudfsNamespaceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineCloudfsNamespaces())
}
