package volume_attach

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VolumeAttach can be imported using the id, e.g.
```
$ terraform import volcengine_volume_attach.default vol-abc12345:i-abc12345
```

*/

func ResourceVolcengineVolumeAttach() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineVolumeAttachCreate,
		Read:   resourceVolcengineVolumeAttachRead,
		Update: resourceVolcengineVolumeAttachUpdate,
		Delete: resourceVolcengineVolumeAttachDelete,
		Importer: &schema.ResourceImporter{
			State: importVolumeAttach,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Minute),
			Delete: schema.DefaultTimeout(2 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"volume_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Id of Volume.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Id of Instance.",
			},
			"delete_with_instance": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				Description: "Delete Volume with Attached Instance." +
					"It is not recommended to use this field. If used, please ensure that the value of this field is consistent with the value of `delete_with_instance` in volcengine_volume.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of Volume.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of Volume.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Update time of Volume.",
			},
		},
	}
}

func resourceVolcengineVolumeAttachCreate(d *schema.ResourceData, meta interface{}) (err error) {
	volumeAttachService := NewVolumeAttachService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(volumeAttachService, d, ResourceVolcengineVolumeAttach())
	if err != nil {
		return fmt.Errorf("error on attach volume %q, %w", d.Id(), err)
	}
	return resourceVolcengineVolumeAttachRead(d, meta)
}

func resourceVolcengineVolumeAttachRead(d *schema.ResourceData, meta interface{}) (err error) {
	volumeAttachService := NewVolumeAttachService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(volumeAttachService, d, ResourceVolcengineVolumeAttach())
	if err != nil {
		return fmt.Errorf("error on reading volume %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineVolumeAttachUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	return resourceVolcengineVolumeAttachRead(d, meta)
}

func resourceVolcengineVolumeAttachDelete(d *schema.ResourceData, meta interface{}) (err error) {
	volumeAttachService := NewVolumeAttachService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(volumeAttachService, d, ResourceVolcengineVolumeAttach())
	if err != nil {
		return fmt.Errorf("error on detach volume %q, %w", d.Id(), err)
	}
	return err
}
