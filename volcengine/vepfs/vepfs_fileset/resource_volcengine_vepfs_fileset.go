package vepfs_fileset

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VepfsFileset can be imported using the file_system_id:fileset_id, e.g.
```
$ terraform import volcengine_vepfs_fileset.default file_system_id:fileset_id
```

*/

func ResourceVolcengineVepfsFileset() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVepfsFilesetCreate,
		Read:   resourceVolcengineVepfsFilesetRead,
		Update: resourceVolcengineVepfsFilesetUpdate,
		Delete: resourceVolcengineVepfsFilesetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"file_system_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the vepfs file system.",
			},
			"fileset_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the vepfs fileset.",
			},
			"fileset_path": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The path of the vepfs fileset.",
			},
			"max_iops": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The max IOPS qos limit of the vepfs fileset.",
			},
			"max_bandwidth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The max bandwidth qos limit of the vepfs fileset. Unit: MB/s.",
			},
			"file_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The file number limit of the vepfs fileset.",
			},
			"capacity_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The capacity limit of the vepfs fileset. Unit: Gib.",
			},

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the vepfs fileset.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the vepfs fileset.",
			},
			"file_used": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The used file number of the vepfs fileset.",
			},
			"capacity_used": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The used capacity of the vepfs fileset. Unit: GiB.",
			},
			"max_inode_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The max number of inode in the vepfs fileset.",
			},
		},
	}
	return resource
}

func resourceVolcengineVepfsFilesetCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsFilesetService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVepfsFileset())
	if err != nil {
		return fmt.Errorf("error on creating vepfs_fileset %q, %s", d.Id(), err)
	}
	return resourceVolcengineVepfsFilesetRead(d, meta)
}

func resourceVolcengineVepfsFilesetRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsFilesetService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVepfsFileset())
	if err != nil {
		return fmt.Errorf("error on reading vepfs_fileset %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVepfsFilesetUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsFilesetService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVepfsFileset())
	if err != nil {
		return fmt.Errorf("error on updating vepfs_fileset %q, %s", d.Id(), err)
	}
	return resourceVolcengineVepfsFilesetRead(d, meta)
}

func resourceVolcengineVepfsFilesetDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsFilesetService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVepfsFileset())
	if err != nil {
		return fmt.Errorf("error on deleting vepfs_fileset %q, %s", d.Id(), err)
	}
	return err
}

var filesetImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("file_system_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("fileset_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
