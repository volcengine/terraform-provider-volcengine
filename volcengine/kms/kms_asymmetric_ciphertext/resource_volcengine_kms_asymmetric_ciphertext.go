package kms_asymmetric_ciphertext

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The KmsAsymmetricCiphertext is not support import.

*/

func ResourceVolcengineKmsAsymmetricCiphertext() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsAsymmetricCiphertextCreate,
		Read:   resourceVolcengineKmsAsymmetricCiphertextRead,
		Delete: resourceVolcengineKmsAsymmetricCiphertextDelete,
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
			"plaintext": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The plaintext to be encrypted, Base64 encoded.",
			},
			"algorithm": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"RSAES_OAEP_SHA_256", "SM2PKE"}, true),
				Description:  "The encryption algorithm. valid values: `RSAES_OAEP_SHA_256`, `SM2PKE`.",
			},
			"ciphertext_blob": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ciphertext, Base64 encoded. The produced ciphertext_blob stays stable across applies. If the plaintext should be re-encrypted on each apply use the `volcengine_kms_asymmetric_ciphertexts` data source.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsAsymmetricCiphertextCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsAsymmetricCiphertextService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsAsymmetricCiphertext())
	if err != nil {
		return fmt.Errorf("error on creating kms_asymmetric_ciphertext %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsAsymmetricCiphertextRead(d, meta)
}

func resourceVolcengineKmsAsymmetricCiphertextRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsAsymmetricCiphertextService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsAsymmetricCiphertext())
	if err != nil {
		return fmt.Errorf("error on reading kms_asymmetric_ciphertext %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsAsymmetricCiphertextDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsAsymmetricCiphertextService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsAsymmetricCiphertext())
	if err != nil {
		return fmt.Errorf("error on deleting kms_asymmetric_ciphertext %q, %s", d.Id(), err)
	}
	return err
}
