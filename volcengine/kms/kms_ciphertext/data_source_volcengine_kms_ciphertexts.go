package kms_ciphertext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsCiphertexts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsCiphertextsRead,
		Schema: map[string]*schema.Schema{
			"keyring_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the keyring.",
			},
			"key_name": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The name of the key.",
			},
			"key_id": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.",
			},
			"plaintext": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The plaintext to be encrypted, Base64 encoded.",
			},
			"encryption_context": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The JSON string of key/value pairs.",
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
			"ciphertext_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The information about the ciphertext.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ciphertext_blob": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ciphertext, Base64 encoded. The plaintext gets re-encrypted on each apply, resulting in a changed ciphertext_blob. If a stable ciphertext is needed use the `volcengine_kms_ciphertext` resource.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineKmsCiphertextsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsCiphertextService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsCiphertexts())
}
