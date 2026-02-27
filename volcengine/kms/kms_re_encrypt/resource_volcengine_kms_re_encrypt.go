package kms_re_encrypt

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The KmsReEncrypt is not support import.

*/

func ResourceVolcengineKmsReEncrypt() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsReEncryptCreate,
		Read:   resourceVolcengineKmsReEncryptRead,
		Delete: resourceVolcengineKmsReEncryptDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"new_keyring_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The new keyring name.",
			},
			"new_key_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{"new_key_name", "new_key_id"},
				Description:  "The new key name.",
			},
			"new_key_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{"new_key_name", "new_key_id"},
				Description:  "The new key id. When new_key_id is not specified, both new_keyring_name and new_key_name must be specified.",
			},
			"source_ciphertext_blob": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The source ciphertext, Base64 encoded.",
			},
			"old_encryption_context": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The old encryption context JSON string.",
			},
			"new_encryption_context": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The new encryption context JSON string.",
			},
			"ciphertext_blob": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The re-encrypted ciphertext, Base64 encoded. The data stays stable across applies. If a changing ciphertext is needed use the `volcengine_kms_re_encrypts` data source.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsReEncryptCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsReEncryptService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsReEncrypt())
	if err != nil {
		return fmt.Errorf("error on creating kms_re_encrypt %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsReEncryptRead(d, meta)
}

func resourceVolcengineKmsReEncryptRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsReEncryptService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsReEncrypt())
	if err != nil {
		return fmt.Errorf("error on reading kms_re_encrypt %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsReEncryptDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsReEncryptService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsReEncrypt())
	if err != nil {
		return fmt.Errorf("error on deleting kms_re_encrypt %q, %s", d.Id(), err)
	}
	return err
}
