package kms_key_material

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

The KmsKeyMaterial is not support import.

*/

func ResourceVolcengineKmsKeyMaterial() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsKeyMaterialCreate,
		Read:   resourceVolcengineKmsKeyMaterialRead,
		Delete: resourceVolcengineKmsKeyMaterialDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			// Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"keyring_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The name of keyring.",
			},
			"key_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The name of key.",
			},
			"key_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The id of key. When key_id is not specified, both keyring_name and key_name must be specified.",
			},
			"encrypted_key_material": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The encrypted key material, Base64 encoded.",
			},
			"import_token": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The import token.",
			},
			"expiration_model": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "KEY_MATERIAL_DOES_NOT_EXPIRE",
				ValidateFunc: validation.StringInSlice([]string{"KEY_MATERIAL_DOES_NOT_EXPIRE", "KEY_MATERIAL_EXPIRES"}, false),
				Description:  "The expiration model of key material. Valid values: `KEY_MATERIAL_DOES_NOT_EXPIRE`, `KEY_MATERIAL_EXPIRES`. Default value: `KEY_MATERIAL_DOES_NOT_EXPIRE`.",
			},
			"valid_to": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The valid to timestamp of key material. Required when expiration_model is KEY_MATERIAL_EXPIRES. Unit: second.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsKeyMaterialCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyMaterialService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsKeyMaterial())
	if err != nil {
		return fmt.Errorf("error on creating kms_key_material %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyMaterialRead(d, meta)
}

func resourceVolcengineKmsKeyMaterialRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyMaterialService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsKeyMaterial())
	if err != nil {
		return fmt.Errorf("error on reading kms_key_material %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsKeyMaterialUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyMaterialService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKmsKeyMaterial())
	if err != nil {
		return fmt.Errorf("error on updating kms_key_material %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyMaterialRead(d, meta)
}

func resourceVolcengineKmsKeyMaterialDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyMaterialService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsKeyMaterial())
	if err != nil {
		return fmt.Errorf("error on deleting kms_key_material %q, %s", d.Id(), err)
	}
	return err
}
