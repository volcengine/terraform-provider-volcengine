package kms_key_rotation

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KmsKeyRotation can be imported using the id, e.g.
```
$ terraform import volcengine_kms_key_rotation.default resource_id
or
$ terraform import volcengine_kms_key_rotation.default key_name:keyring_name
```

*/

func ResourceVolcengineKmsKeyRotation() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsKeyRotationCreate,
		Read:   resourceVolcengineKmsKeyRotationRead,
		Update: resourceVolcengineKmsKeyRotationUpdate,
		Delete: resourceVolcengineKmsKeyRotationDelete,
		Importer: &schema.ResourceImporter{
			State: kmsKeyRotationImporter,
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
			// 支持修改，通过调用 EnableKeyRotation API实现
			"rotate_interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(90, 2560),
				Description:  "Key rotation period, unit: days; value range: [90, 2560].",
			},
			"rotation_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the key rotation.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsKeyRotationCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyRotationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsKeyRotation())
	if err != nil {
		return fmt.Errorf("error on creating kms_key_rotation %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyRotationRead(d, meta)
}

func resourceVolcengineKmsKeyRotationRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyRotationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsKeyRotation())
	if err != nil {
		return fmt.Errorf("error on reading kms_key_rotation %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsKeyRotationUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyRotationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKmsKeyRotation())
	if err != nil {
		return fmt.Errorf("error on updating kms_key_rotation %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyRotationRead(d, meta)
}

func resourceVolcengineKmsKeyRotationDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyRotationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsKeyRotation())
	if err != nil {
		return fmt.Errorf("error on deleting kms_key_rotation %q, %s", d.Id(), err)
	}
	return err
}
