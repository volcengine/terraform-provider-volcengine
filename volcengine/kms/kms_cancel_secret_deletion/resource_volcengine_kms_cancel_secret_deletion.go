package kms_cancel_secret_deletion

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KmsCancelSecretDeletion can be imported using the id, e.g.
```
$ terraform import volcengine_kms_cancel_secret_deletion.default secret_name
```

*/

func ResourceVolcengineKmsCancelSecretDeletion() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsCancelSecretDeletionCreate,
		Read:   resourceVolcengineKmsCancelSecretDeletionRead,
		Delete: resourceVolcengineKmsCancelSecretDeletionDelete,
		Importer: &schema.ResourceImporter{
			State: kmsCancelSecretDeletionImporter,
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
				Description: "The name of the secret.",
			},
			"secret_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the secret.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsCancelSecretDeletionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsCancelSecretDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsCancelSecretDeletion())
	if err != nil {
		return fmt.Errorf("error on creating kms_cancel_secret_deletion %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsCancelSecretDeletionRead(d, meta)
}

func resourceVolcengineKmsCancelSecretDeletionRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsCancelSecretDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsCancelSecretDeletion())
	if err != nil {
		return fmt.Errorf("error on reading kms_cancel_secret_deletion %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsCancelSecretDeletionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsCancelSecretDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsCancelSecretDeletion())
	if err != nil {
		return fmt.Errorf("error on deleting kms_cancel_secret_deletion %q, %s", d.Id(), err)
	}
	return err
}
