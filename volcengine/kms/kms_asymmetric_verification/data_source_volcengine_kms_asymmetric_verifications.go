package kms_asymmetric_verification

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsAsymmetricVerifications() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsAsymmetricVerificationsRead,
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
			"message": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The message to be verified, Base64 encoded.",
			},
			"message_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"RAW", "DIGEST"}, true),
				Default:      "RAW",
				Description:  "The type of message. Valid values: RAW or DIGEST. When message_type is DIGEST, KMS does not process the message digest of the original data source, it will sign directly with the private key.",
			},
			"algorithm": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"RSA_PSS_SHA_256", "RSA_PKCS1_SHA_256", "RSA_PSS_SHA_384", "RSA_PKCS1_SHA_384", "RSA_PSS_SHA_512", "RSA_PKCS1_SHA_512", "ECDSA_SHA_256", "ECDSA_SHA_384", "ECDSA_SHA_512", "SM2_DSA"}, true),
				Description:  "The signing algorithm. valid values: `RSA_PSS_SHA_256`, `RSA_PKCS1_SHA_256`, `RSA_PSS_SHA_384`, `RSA_PKCS1_SHA_384`, `RSA_PSS_SHA_512`, `RSA_PKCS1_SHA_512`.",
			},
			"signature": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The signature to be verified, Base64 encoded.",
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
			"verification_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The verification result.",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"signature_valid": {
						Type:        schema.TypeBool,
						Computed:    true,
						Description: "Whether the signature is valid.",
					},
				}},
			},
		},
	}
}

func dataSourceVolcengineKmsAsymmetricVerificationsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsAsymmetricVerificationService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsAsymmetricVerifications())
}
