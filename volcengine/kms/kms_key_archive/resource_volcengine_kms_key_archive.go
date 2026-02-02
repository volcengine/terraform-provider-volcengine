package kms_key_archive

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KmsKeyArchive can be imported using the id, e.g.
```
$ terraform import volcengine_kms_key_archive.default resource_id
or
$ terraform import volcengine_kms_key_archive.default key_name:keyring_name
```

*/

func ResourceVolcengineKmsKeyArchive() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsKeyArchiveCreate,
		Read:   resourceVolcengineKmsKeyArchiveRead,
		Delete: resourceVolcengineKmsKeyArchiveDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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

func resourceVolcengineKmsKeyArchiveCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyArchiveService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsKeyArchive())
	if err != nil {
		return fmt.Errorf("error on creating kms_key_archive %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyArchiveRead(d, meta)
}

func resourceVolcengineKmsKeyArchiveRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyArchiveService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsKeyArchive())
	if err != nil {
		return fmt.Errorf("error on reading kms_key_archive %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsKeyArchiveUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyArchiveService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKmsKeyArchive())
	if err != nil {
		return fmt.Errorf("error on updating kms_key_archive %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyArchiveRead(d, meta)
}

func resourceVolcengineKmsKeyArchiveDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyArchiveService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsKeyArchive())
	if err != nil {
		return fmt.Errorf("error on deleting kms_key_archive %q, %s", d.Id(), err)
	}
	return err
}
