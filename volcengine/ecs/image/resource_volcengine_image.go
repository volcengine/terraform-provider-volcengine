package image

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Image can be imported using the id, e.g.
```
$ terraform import volcengine_image.default resource_id
```

*/

func ResourceVolcengineImage() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineImageCreate,
		Read:   resourceVolcengineImageRead,
		Update: resourceVolcengineImageUpdate,
		Delete: resourceVolcengineImageDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"image_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the custom image.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the custom image.",
			},
			"instance_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"instance_id", "snapshot_id", "snapshot_group_id"},
				Description: "The instance id of the custom image. Only one of `instance_id, snapshot_id, snapshot_group_id` can be specified." +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"snapshot_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"instance_id", "snapshot_id", "snapshot_group_id"},
				Description: "The snapshot id of the custom image. Only one of `instance_id, snapshot_id, snapshot_group_id` can be specified." +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"snapshot_group_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"instance_id", "snapshot_id", "snapshot_group_id"},
				Description: "The snapshot group id of the custom image. Only one of `instance_id, snapshot_id, snapshot_group_id` can be specified." +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"create_whole_image": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() != ""
				},
				Description: "Whether to create whole image. Default is false. This field is only effective when creating a new custom image.",
			},
			"boot_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				Description: "The boot mode of the custom image. Valid values: `BIOS`, `UEFI`. This field is only effective when modifying the image.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the custom image.",
			},
			"tags": ve.TagsSchema(),

			// computed fields
			"os_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The operating system type of Image.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of Image.",
			},
			"visibility": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The visibility of Image.",
			},
			"architecture": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The architecture of Image.",
			},
			"platform": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The platform of Image.",
			},
			"platform_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The platform version of Image.",
			},
			"os_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of Image operating system.",
			},
			"is_support_cloud_init": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the Image support cloud-init.",
			},
			"share_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The share mode of Image.",
			},
			"size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The size(GiB) of Image.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of Image.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of Image.",
			},
		},
	}
	return resource
}

func resourceVolcengineImageCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewImageService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineImage())
	if err != nil {
		return fmt.Errorf("error on creating image %q, %s", d.Id(), err)
	}
	return resourceVolcengineImageRead(d, meta)
}

func resourceVolcengineImageRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewImageService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineImage())
	if err != nil {
		return fmt.Errorf("error on reading image %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineImageUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewImageService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineImage())
	if err != nil {
		return fmt.Errorf("error on updating image %q, %s", d.Id(), err)
	}
	return resourceVolcengineImageRead(d, meta)
}

func resourceVolcengineImageDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewImageService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineImage())
	if err != nil {
		return fmt.Errorf("error on deleting image %q, %s", d.Id(), err)
	}
	return err
}
