package kms_secret_rotate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KmsSecretRotate can be imported using the secret_name, e.g.
```
$ terraform import volcengine_kms_secret_rotate.default ecs-secret-test
```

*/

func ResourceVolcengineKmsSecretRotate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsSecretRotateCreate,
		Read:   resourceVolcengineKmsSecretRotateRead,
		Update: resourceVolcengineKmsSecretRotateUpdate,
		Delete: resourceVolcengineKmsSecretRotateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the secret to manually rotate.",
			},
			"version_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version alias after rotation. Manual rotation can be triggered by modifying version_name.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsSecretRotateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretRotateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsSecretRotate())
	if err != nil {
		return fmt.Errorf("error on creating kms_secret_rotate %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsSecretRotateRead(d, meta)
}

func resourceVolcengineKmsSecretRotateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretRotateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsSecretRotate())
	if err != nil {
		return fmt.Errorf("error on reading kms_secret_rotate %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsSecretRotateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretRotateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKmsSecretRotate())
	if err != nil {
		return fmt.Errorf("error on updating kms_secret_rotate %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsSecretRotateRead(d, meta)
}

func resourceVolcengineKmsSecretRotateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretRotateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsSecretRotate())
	if err != nil {
		return fmt.Errorf("error on deleting kms_secret_rotate %q, %s", d.Id(), err)
	}
	return err
}
