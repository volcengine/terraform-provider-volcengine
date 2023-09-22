package nas_mount_point

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Nas Mount Point can be imported using the file system id and mount point id, e.g.
```
$ terraform import volcengine_nas_mount_point.default enas-cnbj18bcb923****:mount-a6ee****
```

*/

func ResourceVolcengineNasMountPoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineNasMountPointCreate,
		Read:   resourceVolcengineNasMountPointRead,
		Update: resourceVolcengineNasMountPointUpdate,
		Delete: resourceVolcengineNasMountPointDelete,
		Importer: &schema.ResourceImporter{
			State: importNasMountPoint,
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
				Description: "The file system id.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet id.",
			},
			"permission_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The permission group id.",
			},
			"mount_point_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The mount point name.",
			},
			"mount_point_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The mount point id.",
			},
		},
	}
}

func resourceVolcengineNasMountPointCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineNasMountPointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineNasMountPoint())
	if err != nil {
		return fmt.Errorf("error on creating volcengine Nas Mount Point: %q, %w", d.Id(), err)
	}
	return resourceVolcengineNasMountPointRead(d, meta)
}

func resourceVolcengineNasMountPointRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineNasMountPointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineNasMountPoint())
	if err != nil {
		return fmt.Errorf("error on reading volcengine Nas Mount Point: %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineNasMountPointUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineNasMountPointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineNasMountPoint())
	if err != nil {
		return fmt.Errorf("error on updating volcengine Nas Mount Point: %q, %w", d.Id(), err)
	}
	return resourceVolcengineNasMountPointRead(d, meta)
}

func resourceVolcengineNasMountPointDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineNasMountPointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineNasMountPoint())
	if err != nil {
		return fmt.Errorf("error on deleting volcengine Nas Mount Point: %q, %w", d.Id(), err)
	}
	return nil
}

func importNasMountPoint(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form fileSystemId:mountPointId")
	}
	err = data.Set("file_system_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("mount_point_id", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
