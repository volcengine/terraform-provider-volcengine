package apig_custom_domain

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineApigCustomDomains() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineApigCustomDomainsRead,
		Schema: map[string]*schema.Schema{
			"gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of api gateway.",
			},
			"service_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of api gateway service.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource type of domain. Valid values: `Console`, `Ingress`.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
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
			"custom_domains": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the custom domain.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The custom domain of the api gateway service.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the domain.",
						},
						"certificate_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the certificate.",
						},
						"service_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the api gateway service.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the custom domain.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type of domain.",
						},
						"comments": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The comments of the custom domain.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the custom domain.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the custom domain.",
						},
						"ssl_redirect": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to redirect https.",
						},
						"protocol": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The protocol of the custom domain.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineApigCustomDomainsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewApigCustomDomainService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineApigCustomDomains())
}
