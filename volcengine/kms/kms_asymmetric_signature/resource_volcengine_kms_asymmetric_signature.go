package kms_asymmetric_signature

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The KmsAsymmetricSignature is not support import.

*/

func ResourceVolcengineKmsAsymmetricSignature() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsAsymmetricSignatureCreate,
		Read:   resourceVolcengineKmsAsymmetricSignatureRead,
		Delete: resourceVolcengineKmsAsymmetricSignatureDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"keyring_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The name of the keyring.",
			},
			"key_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The name of the key.",
			},
			"key_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.",
			},
			"message": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The message to be signed, Base64 encoded.",
			},
			"message_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"RAW", "DIGEST"}, true),
				Default:      "RAW",
				Description:  "The type of message. Valid values: RAW or DIGEST. When message_type is DIGEST, KMS does not process the message digest of the original data source, it will sign directly with the private key.",
			},
			"algorithm": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"RSA_PSS_SHA_256", "RSA_PKCS1_SHA_256", "RSA_PSS_SHA_384", "RSA_PKCS1_SHA_384", "RSA_PSS_SHA_512", "RSA_PKCS1_SHA_512", "ECDSA_SHA_256", "ECDSA_SHA_384", "ECDSA_SHA_512", "SM2_DSA"}, true),
				Description:  "The signing algorithm. valid values: `RSA_PSS_SHA_256`, `RSA_PKCS1_SHA_256`, `RSA_PSS_SHA_384`, `RSA_PKCS1_SHA_384`, `RSA_PSS_SHA_512`, `RSA_PKCS1_SHA_512`.",
			},
			"signature": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The signature, Base64 encoded. The produced signature stays stable across applies. If the message should be re-signed on each apply use the `volcengine_kms_asymmetric_signatures` data source.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsAsymmetricSignatureCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsAsymmetricSignatureService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsAsymmetricSignature())
	if err != nil {
		return fmt.Errorf("error on creating kms_asymmetric_signature %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsAsymmetricSignatureRead(d, meta)
}

func resourceVolcengineKmsAsymmetricSignatureRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsAsymmetricSignatureService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsAsymmetricSignature())
	if err != nil {
		return fmt.Errorf("error on reading kms_asymmetric_signature %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsAsymmetricSignatureDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsAsymmetricSignatureService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsAsymmetricSignature())
	if err != nil {
		return fmt.Errorf("error on deleting kms_asymmetric_signature %q, %s", d.Id(), err)
	}
	return err
}
