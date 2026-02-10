package kms_asymmetric_plaintext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsAsymmetricPlaintexts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsAsymmetricPlaintextsRead,
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
			"ciphertext_blob": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ciphertext to be decrypted, Base64 encoded.",
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
			"plaintext_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The decrypted plaintext.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"plaintext": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The decrypted plaintext, Base64 encoded.",
					},
				}},
			},
		},
	}
}

func dataSourceVolcengineKmsAsymmetricPlaintextsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsAsymmetricPlaintextService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsAsymmetricPlaintexts())
}
