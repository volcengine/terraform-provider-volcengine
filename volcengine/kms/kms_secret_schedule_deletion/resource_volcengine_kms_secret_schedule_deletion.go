package kms_secret_schedule_deletion

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KmsSecretScheduleDeletion can be imported using the id, e.g.
```
$ terraform import volcengine_kms_secret_schedule_deletion.default resource_id
```

*/

func ResourceVolcengineKmsSecretScheduleDeletion() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsSecretScheduleDeletionCreate,
		Read:   resourceVolcengineKmsSecretScheduleDeletionRead,
		Update: resourceVolcengineKmsSecretScheduleDeletionUpdate,
		Delete: resourceVolcengineKmsSecretScheduleDeletionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the secret.",
			},
			"pending_window_in_days": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of days after which the secret will be deleted.",
			},
			"secret_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of secret.",
			},
			"uuid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of secret.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsSecretScheduleDeletionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretScheduleDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsSecretScheduleDeletion())
	if err != nil {
		return fmt.Errorf("error on creating kms_secret_schedule_deletion %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsSecretScheduleDeletionRead(d, meta)
}

func resourceVolcengineKmsSecretScheduleDeletionRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretScheduleDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsSecretScheduleDeletion())
	if err != nil {
		return fmt.Errorf("error on reading kms_secret_schedule_deletion %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsSecretScheduleDeletionUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretScheduleDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKmsSecretScheduleDeletion())
	if err != nil {
		return fmt.Errorf("error on updating kms_secret_schedule_deletion %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsSecretScheduleDeletionRead(d, meta)
}

func resourceVolcengineKmsSecretScheduleDeletionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretScheduleDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsSecretScheduleDeletion())
	if err != nil {
		return fmt.Errorf("error on deleting kms_secret_schedule_deletion %q, %s", d.Id(), err)
	}
	return err
}
