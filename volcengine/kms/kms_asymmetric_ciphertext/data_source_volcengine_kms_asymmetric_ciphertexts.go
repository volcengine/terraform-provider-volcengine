package kms_asymmetric_ciphertext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsAsymmetricCiphertexts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsAsymmetricCiphertextsRead,
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
			"algorithm": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"RSAES_OAEP_SHA_256", "SM2PKE"}, true),
				Description:  "The encryption algorithm. valid values: `RSAES_OAEP_SHA_256`, `SM2PKE`.",
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
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"ciphertext_blob": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The ciphertext, Base64 encoded. The plaintext gets re-encrypted on each apply, resulting in a changed ciphertext_blob. If a stable ciphertext is needed use the `volcengine_kms_asymmetric_ciphertext` resource.",
					},
				}},
			},
		},
	}
}

func dataSourceVolcengineKmsAsymmetricCiphertextsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsAsymmetricCiphertextService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsAsymmetricCiphertexts())
}
