package veecp_kubeconfig

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVeecpKubeconfigs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVeecpKubeconfigsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Kubeconfig IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Kubeconfig.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Kubeconfig query.",
			},
			"page_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The page number of Kubeconfigs query.",
			},
			"page_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The page size of Kubeconfigs query.",
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
			"role_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:         schema.HashInt,
				Description: "A list of Role IDs.",
			},
			"types": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The type of Kubeconfigs query.",
			},

			"kubeconfigs": {
				Description: "The collection of VkeKubeconfig query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Kubeconfig.",
						},
						"kubeconfig_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Kubeconfig.",
						},
						"user_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The account ID of the Kubeconfig.",
						},
						"cluster_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Cluster ID of the Kubeconfig.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the Kubeconfig.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the Kubeconfig.",
						},
						"expire_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expire time of the Kubeconfig.",
						},
						"kubeconfig": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Kubeconfig data with public/private network access, returned in BASE64 encoding.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVeecpKubeconfigsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVeecpKubeconfigService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVeecpKubeconfigs())
}
