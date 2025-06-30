package kms_keyring

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KmsKeyring can be imported using the id, e.g.
```
$ terraform import volcengine_kms_keyring.default resource_id
```

*/

func ResourceVolcengineKmsKeyring() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsKeyringCreate,
		Read:   resourceVolcengineKmsKeyringRead,
		Update: resourceVolcengineKmsKeyringUpdate,
		Delete: resourceVolcengineKmsKeyringDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"keyring_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the keyring.",
			},
			"keyring_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The type of the keyring.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the keyring.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the project.",
			},
			"creation_date": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The date when the keyring was created.",
			},
			"update_date": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The date when the keyring was updated.",
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The tenant ID of the keyring.",
			},
			"trn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The information about the tenant resource name (TRN).",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsKeyringCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyringService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsKeyring())
	if err != nil {
		return fmt.Errorf("error on creating kms_keyring %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyringRead(d, meta)
}

func resourceVolcengineKmsKeyringRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyringService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsKeyring())
	if err != nil {
		return fmt.Errorf("error on reading kms_keyring %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsKeyringUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyringService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKmsKeyring())
	if err != nil {
		return fmt.Errorf("error on updating kms_keyring %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyringRead(d, meta)
}

func resourceVolcengineKmsKeyringDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyringService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsKeyring())
	if err != nil {
		return fmt.Errorf("error on deleting kms_keyring %q, %s", d.Id(), err)
	}
	return err
}
