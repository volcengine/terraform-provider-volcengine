package kms_secret_version

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsSecretVersions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsSecretVersionsRead,
		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the secret.",
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
			"secret_versions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The version info of secret.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version ID of secret value.",
						},
						"version_stage": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version stage of secret value.",
						},
						"creation_date": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The creation time of secret version.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineKmsSecretVersionsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsSecretVersionService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsSecretVersions())
}
