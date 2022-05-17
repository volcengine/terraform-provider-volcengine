package certificate

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

func DataSourceVestackCertificates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVestackCertificatesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list of Certificate IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "The Name Regex of Certificate.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Certificate query.",
			},
			"certificate_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the Certificate.",
			},
			"certificates": {
				Description: "The collection of Certificate query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true, // tf中不支持写值
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
						"listeners": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Computed:    true,
							Description: "The ID list of the Listener.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVestackCertificatesRead(d *schema.ResourceData, meta interface{}) error {
	certificateService := NewCertificateService(meta.(*ve.SdkClient))
	return certificateService.Dispatcher.Data(certificateService, d, DataSourceVestackCertificates())
}
