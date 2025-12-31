package alb_all_certificate

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAlbAllCertificates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAlbAllCertificatesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
			},
			"certificate_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Name of Certificate.",
			},
			"certificate_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"CA", "Server"}, false),
				Description:  "The type of Certificate. Valid values: `CA`, `Server`.",
			},
			"tags": ve.TagsSchema(),
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of Certificate.",
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
			"certificates": {
				Description: "The collection of Certificate query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Certificate.",
						},
						"certificate_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Certificate.",
						},
						"certificate_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Certificate.",
						},
						"certificate_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the Certificate.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the Certificate.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the Certificate.",
						},
						"expired_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expire time of the Certificate.",
						},
						"domain_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain name of the Certificate.",
						},
						"san": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The list of extended domain names for the certificate, separated by English commas ',', including (commonName, DnsName, IP).",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the Certificate.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ProjectName of the Certificate.",
						},
						"listeners": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "The ID list of the Listener.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineAlbAllCertificatesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewAlbAllCertificateService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineAlbAllCertificates())
}
