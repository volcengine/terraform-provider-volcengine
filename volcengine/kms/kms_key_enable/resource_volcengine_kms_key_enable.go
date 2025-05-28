package kms_key_enable

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KmsKeyEnable can be imported using the id, e.g.
```
$ terraform import volcengine_kms_key_enable.default resource_id
or
$ terraform import volcengine_kms_key_enable.default key_name:keyring_name
```

*/

func ResourceVolcengineKmsKeyEnable() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsKeyEnableCreate,
		Read:   resourceVolcengineKmsKeyEnableRead,
		Update: resourceVolcengineKmsKeyEnableUpdate,
		Delete: resourceVolcengineKmsKeyEnableDelete,
		Importer: &schema.ResourceImporter{
			State: kmsKeyEnableImporter,
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
			"key_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the key.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsKeyEnableCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyEnableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsKeyEnable())
	if err != nil {
		return fmt.Errorf("error on creating kms_key_enable %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyEnableRead(d, meta)
}

func resourceVolcengineKmsKeyEnableRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyEnableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsKeyEnable())
	if err != nil {
		return fmt.Errorf("error on reading kms_key_enable %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsKeyEnableUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyEnableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKmsKeyEnable())
	if err != nil {
		return fmt.Errorf("error on updating kms_key_enable %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyEnableRead(d, meta)
}

func resourceVolcengineKmsKeyEnableDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyEnableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsKeyEnable())
	if err != nil {
		return fmt.Errorf("error on deleting kms_key_enable %q, %s", d.Id(), err)
	}
	return err
}
