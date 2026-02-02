package kms_secret_restore

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The KmsSecretRestore is not support import.

*/

func ResourceVolcengineKmsSecretRestore() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsSecretRestoreCreate,
		Read:   resourceVolcengineKmsSecretRestoreRead,
		Delete: resourceVolcengineKmsSecretRestoreDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"secret_data_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The data key ciphertext returned during backup. Base64 encoded.",
			},
			"backup_data": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The full secret data returned during backup. JSON format.",
			},
			"signature": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The signature of the backup data returned during backup. Base64 encoded.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsSecretRestoreCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretRestoreService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsSecretRestore())
	if err != nil {
		return fmt.Errorf("error on creating kms_secret_restore %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsSecretRestoreRead(d, meta)
}

func resourceVolcengineKmsSecretRestoreRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretRestoreService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsSecretRestore())
	if err != nil {
		return fmt.Errorf("error on reading kms_secret_restore %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsSecretRestoreDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretRestoreService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsSecretRestore())
	if err != nil {
		return fmt.Errorf("error on deleting kms_secret_restore %q, %s", d.Id(), err)
	}
	return err
}
