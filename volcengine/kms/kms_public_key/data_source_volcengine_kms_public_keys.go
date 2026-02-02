package kms_public_key

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsPublicKeys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsPublicKeysRead,
		Schema: map[string]*schema.Schema{
			"keyring_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of keyring.",
			},
			"key_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of key.",
			},
			"key_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of key. When key_id is not specified, both keyring_name and key_name must be specified.",
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
			"public_key": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The public key info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of key.",
						},
						"public_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public key in PEM format.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineKmsPublicKeysRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsPublicKeyService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsPublicKeys())
}
