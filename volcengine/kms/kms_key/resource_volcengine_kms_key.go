package kms_key

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KmsKey can be imported using the id, e.g.
```
$ terraform import volcengine_kms_key.default resource_id
```

*/

func ResourceVolcengineKmsKey() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsKeyCreate,
		Read:   resourceVolcengineKmsKeyRead,
		Update: resourceVolcengineKmsKeyUpdate,
		Delete: resourceVolcengineKmsKeyDelete,
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
			"key_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the CMK.",
			},
			"key_spec": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The type of the keys.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the key.",
			},
			"key_usage": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The usage of the key.",
			},
			"protection_level": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The protection level of the key.",
			},
			"rotate_state": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The rotation state of the key.",
			},
			"origin": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The origin of the key.",
			},
			"multi_region": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "Whether it is the master key of the Multi-region type.",
			},
			"tags": ve.TagsSchema(),
			"pending_window_in_days": {
				Type:        schema.TypeInt,
				ForceNew:    true,
				Optional:    true,
				Description: "The pre-deletion cycle of the key.",
			},
			// computed
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
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the key.",
			},
			"schedule_delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time when the key will be deleted.",
			},
			"rotation_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The rotation configuration of the key.",
			},
			"last_rotation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The last time the key was rotated.",
			},
			"schedule_rotation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The next time the key will be rotated.",
			},
			"trn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource.",
			},
			"key_material_expire_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time when the key material will expire.",
			},
			"multi_region_configuration": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "The configuration of Multi-region key.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"multi_region_key_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the multi-region key.",
						},
						"primary_key": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "Trn and region id of the primary multi-region key.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The trn of multi-region key.",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region id of multi-region key.",
									},
								},
							},
						},
						"replica_keys": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Trn and region id of replica multi-region keys.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The trn of multi-region key.",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region id of multi-region key.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineKmsKeyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsKey())
	if err != nil {
		return fmt.Errorf("error on creating kms_key %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyRead(d, meta)
}

func resourceVolcengineKmsKeyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsKey())
	if err != nil {
		return fmt.Errorf("error on reading kms_key %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsKeyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKmsKey())
	if err != nil {
		return fmt.Errorf("error on updating kms_key %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsKeyRead(d, meta)
}

func resourceVolcengineKmsKeyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsKey())
	if err != nil {
		return fmt.Errorf("error on deleting kms_key %q, %s", d.Id(), err)
	}
	return err
}
