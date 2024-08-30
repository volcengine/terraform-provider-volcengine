package vepfs_mount_service

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VepfsMountService can be imported using the id, e.g.
```
$ terraform import volcengine_vepfs_mount_service.default resource_id
```

*/

func ResourceVolcengineVepfsMountService() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVepfsMountServiceCreate,
		Read:   resourceVolcengineVepfsMountServiceRead,
		Update: resourceVolcengineVepfsMountServiceUpdate,
		Delete: resourceVolcengineVepfsMountServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"mount_service_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the mount service.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet id of the mount service.",
			},
			"node_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The node type of the mount service. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The node type of the mount service.",
			},

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the mount service.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account id of the mount service.",
			},
			"region_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The region id of the mount service.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The zone id of the mount service.",
			},
			"zone_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The zone name of the mount service.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The vpc id of the mount service.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The created time of the mount service.",
			},
			"nodes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The nodes info of the mount service.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of ecs instance.",
						},
						"default_password": {
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Description: "The default password of ecs instance.",
						},
					},
				},
			},
			"attach_file_systems": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The attached file system info of the mount service.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_system_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vepfs file system.",
						},
						"file_system_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the vepfs file system.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id of the vepfs file system.",
						},
						"customer_path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vepfs file system.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the vepfs file system.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineVepfsMountServiceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsMountServiceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVepfsMountService())
	if err != nil {
		return fmt.Errorf("error on creating vepfs_mount_service %q, %s", d.Id(), err)
	}
	return resourceVolcengineVepfsMountServiceRead(d, meta)
}

func resourceVolcengineVepfsMountServiceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsMountServiceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVepfsMountService())
	if err != nil {
		return fmt.Errorf("error on reading vepfs_mount_service %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVepfsMountServiceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsMountServiceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVepfsMountService())
	if err != nil {
		return fmt.Errorf("error on updating vepfs_mount_service %q, %s", d.Id(), err)
	}
	return resourceVolcengineVepfsMountServiceRead(d, meta)
}

func resourceVolcengineVepfsMountServiceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsMountServiceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVepfsMountService())
	if err != nil {
		return fmt.Errorf("error on deleting vepfs_mount_service %q, %s", d.Id(), err)
	}
	return err
}
