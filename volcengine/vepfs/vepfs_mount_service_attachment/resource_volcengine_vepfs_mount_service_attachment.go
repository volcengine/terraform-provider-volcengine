package vepfs_mount_service_attachment

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VepfsMountServiceAttachment can be imported using the mount_service_id:file_system_id, e.g.
```
$ terraform import volcengine_vepfs_mount_service_attachment.default mount_service_id:file_system_id
```

*/

func ResourceVolcengineVepfsMountServiceAttachment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVepfsMountServiceAttachmentCreate,
		Read:   resourceVolcengineVepfsMountServiceAttachmentRead,
		Delete: resourceVolcengineVepfsMountServiceAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: attachmentImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"mount_service_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the mount service.",
			},
			"file_system_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the vepfs file system.",
			},
			"customer_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The custom mount directory, the default value is file system id.",
			},

			// computed fields
			"attach_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The attach status of the vepfs file system.",
			},
		},
	}
	return resource
}

func resourceVolcengineVepfsMountServiceAttachmentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsMountServiceAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVepfsMountServiceAttachment())
	if err != nil {
		return fmt.Errorf("error on creating vepfs_mount_service_attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineVepfsMountServiceAttachmentRead(d, meta)
}

func resourceVolcengineVepfsMountServiceAttachmentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsMountServiceAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVepfsMountServiceAttachment())
	if err != nil {
		return fmt.Errorf("error on reading vepfs_mount_service_attachment %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVepfsMountServiceAttachmentUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsMountServiceAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVepfsMountServiceAttachment())
	if err != nil {
		return fmt.Errorf("error on updating vepfs_mount_service_attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineVepfsMountServiceAttachmentRead(d, meta)
}

func resourceVolcengineVepfsMountServiceAttachmentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsMountServiceAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVepfsMountServiceAttachment())
	if err != nil {
		return fmt.Errorf("error on deleting vepfs_mount_service_attachment %q, %s", d.Id(), err)
	}
	return err
}

var attachmentImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("mount_service_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("file_system_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
