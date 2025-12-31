package alb_ca_certificate

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAlbCaCertificates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAlbCaCertificatesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of CA certificate IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
			},
			"ca_certificate_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the CA certificate.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of the CA certificate.",
			},
			// "tags": ve.TagsSchema(),
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
			"certificates": {
				Description: "The collection of CA certificates query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_certificate_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the CA certificate.",
						},
						"ca_certificate_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the CA certificate.",
						},
						"certificate_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the CA certificate.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the CA certificate.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the CA Certificate.",
						},
						"expired_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expire time of the CA Certificate.",
						},
						"domain_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain name of the CA Certificate.",
						},
						"san": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The san extension of the CA Certificate.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the CA Certificate.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ProjectName of the CA Certificate.",
						},
						// "tags": ve.TagsSchemaComputed(),
						"listeners": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "The ID list of the CA Listener.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineAlbCaCertificatesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewAlbCaCertificateService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineAlbCaCertificates())
}
