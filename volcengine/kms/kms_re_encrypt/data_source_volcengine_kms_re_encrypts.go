package kms_re_encrypt

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsReEncrypts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsReEncryptsRead,
		Schema: map[string]*schema.Schema{
			"new_keyring_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The new keyring name.",
			},
			"new_key_name": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"new_key_name", "new_key_id"},
				Description:  "The new key name.",
			},
			"new_key_id": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"new_key_name", "new_key_id"},
				Description:  "The new key id. When new_key_id is not specified, both new_keyring_name and new_key_name must be specified.",
			},
			"old_encryption_context": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The old encryption context JSON string of key/value pairs.",
			},
			"new_encryption_context": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The new encryption context JSON string of key/value pairs.",
			},
			"source_ciphertext_blob": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ciphertext data to be re-encrypted, Base64 encoded.",
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
							Description: "The re-encrypted ciphertext, Base64 encoded. The data gets re-encrypted on each apply, resulting in a changed ciphertext_blob. If a stable ciphertext is needed use the `volcengine_kms_re_encrypt` resource.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineKmsReEncryptsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsReEncryptService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsReEncrypts())
}
