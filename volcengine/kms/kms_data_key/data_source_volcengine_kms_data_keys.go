package kms_data_key

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsDataKeys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsDataKeysRead,
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
				Description:  "The name of the key. Only symmetric key is supported.",
			},
			"key_id": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.",
			},
			"encryption_context": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The JSON string of key/value pairs.",
			},
			"number_of_bytes": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(1, 1024),
				Description:  "The length of data key to generate.",
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
			"data_key_info": {
				Description: "The data key info.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"plaintext": {
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Description: "The generated plaintext, Base64 encoded.",
						},
						"ciphertext_blob": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The generated ciphertext, Base64 encoded.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineKmsDataKeysRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsDataKeyService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsDataKeys())
}
