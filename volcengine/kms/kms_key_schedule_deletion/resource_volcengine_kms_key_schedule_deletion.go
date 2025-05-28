package kms_key_schedule_deletion

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KmsKeyScheduleDeletion can be imported using the id, e.g.
```
$ terraform import volcengine_kms_key_schedule_deletion.default resource_id
or
$ terraform import volcengine_kms_key_schedule_deletion.default key_name:keyring_name
```

*/

func ResourceVolcengineKmsKeyScheduleDeletion() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsKeyScheduleDeletionCreate,
		Read:   resourceVolcengineKmsKeyScheduleDeletionRead,
		Update: resourceVolcengineKmsKeyScheduleDeletionUpdate,
		Delete: resourceVolcengineKmsKeyScheduleDeletionDelete,
		Importer: &schema.ResourceImporter{
			State: kmsKeyScheduleDeletionImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"keyring_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the keyring.",
			},
			"key_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The name of the CMK.",
			},
			"key_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The id of the CMK.",
			},
			"pending_window_in_days": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of days after which the CMK will be deleted.",
			},
			"key_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the key.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsKeyScheduleDeletionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyScheduleDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsKeyScheduleDeletion())
	if err != nil {
		return fmt.Errorf("error on creating kms_key_schedule_deletion %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyScheduleDeletionRead(d, meta)
}

func resourceVolcengineKmsKeyScheduleDeletionRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyScheduleDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsKeyScheduleDeletion())
	if err != nil {
		return fmt.Errorf("error on reading kms_key_schedule_deletion %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsKeyScheduleDeletionUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyScheduleDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKmsKeyScheduleDeletion())
	if err != nil {
		return fmt.Errorf("error on updating kms_key_schedule_deletion %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyScheduleDeletionRead(d, meta)
}

func resourceVolcengineKmsKeyScheduleDeletionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyScheduleDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsKeyScheduleDeletion())
	if err != nil {
		return fmt.Errorf("error on deleting kms_key_schedule_deletion %q, %s", d.Id(), err)
	}
	return err
}
