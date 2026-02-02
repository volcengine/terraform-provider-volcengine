package kms_cancel_key_deletion

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KmsCancelKeyDeletion can be imported using the id, e.g.
```
$ terraform import volcengine_kms_cancel_key_deletion.default resource_id
or
$ terraform import volcengine_kms_cancel_key_deletion.default key_name:keyring_name
```

*/

func ResourceVolcengineKmsCancelKeyDeletion() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsCancelKeyDeletionCreate,
		Read:   resourceVolcengineKmsCancelKeyDeletionRead,
		Delete: resourceVolcengineKmsCancelKeyDeletionDelete,
		Importer: &schema.ResourceImporter{
			State: kmsCancelKeyDeletionImporter,
		},
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
				Computed:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The name of the key.",
			},
			"key_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.",
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

func resourceVolcengineKmsCancelKeyDeletionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsCancelKeyDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsCancelKeyDeletion())
	if err != nil {
		return fmt.Errorf("error on creating kms_cancel_key_deletion %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsCancelKeyDeletionRead(d, meta)
}

func resourceVolcengineKmsCancelKeyDeletionRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsCancelKeyDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsCancelKeyDeletion())
	if err != nil {
		return fmt.Errorf("error on reading kms_cancel_key_deletion %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsCancelKeyDeletionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsCancelKeyDeletionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsCancelKeyDeletion())
	if err != nil {
		return fmt.Errorf("error on deleting kms_cancel_key_deletion %q, %s", d.Id(), err)
	}
	return err
}
