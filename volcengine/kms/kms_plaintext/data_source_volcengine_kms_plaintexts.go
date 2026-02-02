package kms_plaintext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsPlaintexts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsPlaintextsRead,
		Schema: map[string]*schema.Schema{
			"ciphertext_blob": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ciphertext to be decrypted.",
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
			"plaintext_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The decrypted plaintext.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"plaintext": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The decrypted plaintext, Base64 encoded.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineKmsPlaintextsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsPlaintextService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsPlaintexts())
}
