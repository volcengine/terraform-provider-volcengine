package volume

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Volume can be imported using the id, e.g.
```
$ terraform import volcengine_volume.default vol-mizl7m1kqccg5smt1bdpijuj
```
*/

func ResourceVolcengineVolume() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineVolumeCreate,
		Read:   resourceVolcengineVolumeRead,
		Update: resourceVolcengineVolumeUpdate,
		Delete: resourceVolcengineVolumeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the Zone.",
			},
			"volume_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of Volume.",
			},
			"volume_type": {
				Type:     schema.TypeString,
				Required: true,
				//ForceNew:    true,
				Description: "The type of Volume. Valid values: `ESSD_PL0`, `ESSD_FlexPL`, `TSSD_TL0`.",
			},
			"kind": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The kind of Volume, the value is `data`.",
			},
			"size": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The size of Volume.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the Volume.",
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Description: "The ID of the instance to which the created volume is automatically attached. It is recommended to attach the PostPaid volume to instance through resource `volume_attach`." +
					"When use this field to attach ecs instance, the attached volume cannot be deleted by terraform, please use `terraform state rm volcengine_volume.resource_name` command to remove it from terraform state file and management.",
			},
			"snapshot_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "The id of the snapshot. When creating a volume using snapshots, this field is required.\n" +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"volume_charge_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "PostPaid",
				Description: "The charge type of the Volume, the value is `PostPaid` or `PrePaid`. " +
					"The `PrePaid` volume cannot be detached.",
			},
			"extra_performance_type_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of extra performance for volume. The valid values for ESSD FlexPL volume are `Throughput`, `Balance`, `IOPS`. The valid value for TSSD_TL0 volume is `Throughput`.",
			},
			"extra_performance_iops": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					_, ok := d.GetOk("extra_performance_type_id")
					return !ok
				},
				Description: "The extra IOPS performance size for volume. Unit: times per second. The valid values for `Balance` and `IOPS` is 0~50000.",
			},
			"extra_performance_throughput_mb": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					_, ok := d.GetOk("extra_performance_type_id")
					return !ok
				},
				Description: "The extra Throughput performance size for volume. Unit: MB/s. The valid values for ESSD FlexPL volume is 0~650.",
			},
			"delete_with_instance": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Delete Volume with Attached Instance.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ProjectName of the Volume.",
			},
			"tags": ve.TagsSchema(),

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of Volume.",
			},
			"trade_status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Status of Trade.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of Volume.",
			},
		},
	}
}

func resourceVolcengineVolumeCreate(d *schema.ResourceData, meta interface{}) (err error) {
	volumeService := NewVolumeService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(volumeService, d, ResourceVolcengineVolume())
	if err != nil {
		return fmt.Errorf("error on creating volume %q, %w", d.Id(), err)
	}
	return resourceVolcengineVolumeRead(d, meta)
}

func resourceVolcengineVolumeRead(d *schema.ResourceData, meta interface{}) (err error) {
	volumeService := NewVolumeService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(volumeService, d, ResourceVolcengineVolume())
	if err != nil {
		return fmt.Errorf("error on reading volume %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineVolumeUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	volumeService := NewVolumeService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(volumeService, d, ResourceVolcengineVolume())
	if err != nil {
		return fmt.Errorf("error on updating volume %q, %w", d.Id(), err)
	}
	return resourceVolcengineVolumeRead(d, meta)
}

func resourceVolcengineVolumeDelete(d *schema.ResourceData, meta interface{}) (err error) {
	volumeService := NewVolumeService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(volumeService, d, ResourceVolcengineVolume())
	if err != nil {
		return fmt.Errorf("error on deleting volume %q, %w", d.Id(), err)
	}
	return err
}
