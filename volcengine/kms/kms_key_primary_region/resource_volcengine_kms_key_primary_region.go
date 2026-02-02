package kms_key_primary_region

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*
Import
KmsKeyPrimaryRegion can be imported using the id, e.g.
```
$ terraform import volcengine_kms_key_primary_region.default key_id
or
$ terraform import volcengine_kms_key_primary_region.default key_name:keyring_name
```
*/
func ResourceVolcengineKmsKeyPrimaryRegion() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsKeyPrimaryRegionCreate,
		Read:   resourceVolcengineKmsKeyPrimaryRegionRead,
		// Update: resourceVolcengineKmsKeyPrimaryRegionUpdate,
		Delete: resourceVolcengineKmsKeyPrimaryRegionDelete,
		Importer: &schema.ResourceImporter{
			State: kmsKeyPrimaryRegionImporter,
		},
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
				Description: "The name of the keyring.",
			},
			"key_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The name of the key. Note: Only multi-region keys support updating primary region.",
			},
			"key_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Computed:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.",
			},
			// 注意：在切换密钥的主地域成功之后，例如主地域从 cn-beijing 切换到 cn-shanghai，那么当前地域会变为副本密钥，不再支持密钥轮转、切换主地域等操作
			// 因此继续对密钥进行操作，需要把地域切换为 cn-shanghai 地域
			"primary_region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The new primary region.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsKeyPrimaryRegionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyPrimaryRegionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsKeyPrimaryRegion())
	if err != nil {
		return fmt.Errorf("error on creating kms_primary_region %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyPrimaryRegionRead(d, meta)
}

func resourceVolcengineKmsKeyPrimaryRegionRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyPrimaryRegionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsKeyPrimaryRegion())
	if err != nil {
		return fmt.Errorf("error on reading kms_primary_region %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsKeyPrimaryRegionUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyPrimaryRegionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKmsKeyPrimaryRegion())
	if err != nil {
		return fmt.Errorf("error on updating kms_primary_region %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyPrimaryRegionRead(d, meta)
}

func resourceVolcengineKmsKeyPrimaryRegionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyPrimaryRegionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsKeyPrimaryRegion())
	if err != nil {
		return fmt.Errorf("error on deleting kms_primary_region %q, %s", d.Id(), err)
	}
	return err
}
