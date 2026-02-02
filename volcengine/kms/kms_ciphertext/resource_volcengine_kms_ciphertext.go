package kms_ciphertext

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The KmsCiphertext is not support import.

*/

func ResourceVolcengineKmsCiphertext() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsCiphertextCreate,
		Read:   resourceVolcengineKmsCiphertextRead,
		Delete: resourceVolcengineKmsCiphertextDelete,
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
				Description: "The plaintext to be symmetrically encrypted, Base64 encoded.",
			},
			"encryption_context": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The JSON string of key/value pairs.",
			},
			"ciphertext_blob": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ciphertext, Base64 encoded. The produced ciphertext_blob stays stable across applies. If the plaintext should be re-encrypted on each apply use the `volcengine_kms_ciphertexts` data source.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsCiphertextCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsCiphertextService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsCiphertext())
	if err != nil {
		return fmt.Errorf("error on creating kms_ciphertext %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsCiphertextRead(d, meta)
}

func resourceVolcengineKmsCiphertextRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsCiphertextService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsCiphertext())
	if err != nil {
		return fmt.Errorf("error on reading kms_ciphertext %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsCiphertextDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsCiphertextService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsCiphertext())
	if err != nil {
		return fmt.Errorf("error on deleting kms_ciphertext %q, %s", d.Id(), err)
	}
	return err
}
