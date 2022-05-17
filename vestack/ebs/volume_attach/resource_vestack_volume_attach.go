package volume_attach

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
VolumeAttach can be imported using the id, e.g.
```
$ terraform import vestack_volume_attach.default vol-abc12345:i-abc12345
```

*/

func ResourceVestackVolumeAttach() *schema.Resource {
	return &schema.Resource{
		Create: resourceVestackVolumeAttachCreate,
		Read:   resourceVestackVolumeAttachRead,
		Update: resourceVestackVolumeAttachUpdate,
		Delete: resourceVestackVolumeAttachDelete,
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
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Delete Volume with Attached Instance.",
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

func resourceVestackVolumeAttachCreate(d *schema.ResourceData, meta interface{}) (err error) {
	volumeAttachService := NewVolumeAttachService(meta.(*ve.SdkClient))
	err = volumeAttachService.Dispatcher.Create(volumeAttachService, d, ResourceVestackVolumeAttach())
	if err != nil {
		return fmt.Errorf("error on attach volume %q, %w", d.Id(), err)
	}
	return resourceVestackVolumeAttachRead(d, meta)
}

func resourceVestackVolumeAttachRead(d *schema.ResourceData, meta interface{}) (err error) {
	volumeAttachService := NewVolumeAttachService(meta.(*ve.SdkClient))
	err = volumeAttachService.Dispatcher.Read(volumeAttachService, d, ResourceVestackVolumeAttach())
	if err != nil {
		return fmt.Errorf("error on reading volume %q, %w", d.Id(), err)
	}
	return err
}

func resourceVestackVolumeAttachUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	return resourceVestackVolumeAttachRead(d, meta)
}

func resourceVestackVolumeAttachDelete(d *schema.ResourceData, meta interface{}) (err error) {
	volumeAttachService := NewVolumeAttachService(meta.(*ve.SdkClient))
	err = volumeAttachService.Dispatcher.Delete(volumeAttachService, d, ResourceVestackVolumeAttach())
	if err != nil {
		return fmt.Errorf("error on detach volume %q, %w", d.Id(), err)
	}
	return err
}
