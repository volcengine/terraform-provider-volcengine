package kms_secret_backup

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KmsSecretBackup can be imported using the secret_name, e.g.
```
$ terraform import volcengine_kms_secret_backup.default ecs-secret-test
```

*/

func ResourceVolcengineKmsSecretBackup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsSecretBackupCreate,
		Read:   resourceVolcengineKmsSecretBackupRead,
		Delete: resourceVolcengineKmsSecretBackupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the secret to backup.",
			},
			"secret_data_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ciphertext of the data key used to encrypt the secret value, Base64 encoded.",
			},
			"backup_data": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The full backup data of the secret. JSON format.",
			},
			"signature": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The signature of the backup_data. Base64 encoded.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsSecretBackupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsSecretBackup())
	if err != nil {
		return fmt.Errorf("error on creating kms_secret_backup %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsSecretBackupRead(d, meta)
}

func resourceVolcengineKmsSecretBackupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsSecretBackup())
	if err != nil {
		return fmt.Errorf("error on reading kms_secret_backup %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsSecretBackupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsSecretBackup())
	if err != nil {
		return fmt.Errorf("error on deleting kms_secret_backup %q, %s", d.Id(), err)
	}
	return err
}
