package cloudfs_access

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudFs Access can be imported using the FsName:AccessId, e.g.
```
$ terraform import volcengine_cloudfs_file_system.default tfname:access-**rdgmedx3fow
```

*/

func ResourceVolcengineCloudfsAccess() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudfsAccessCreate,
		Read:   resourceVolcengineCloudfsAccessRead,
		Update: resourceVolcengineCloudfsAccessUpdate,
		Delete: resourceVolcengineCloudfsAccessDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("fs_name", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("access_id", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"fs_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of file system.",
			},
			"access_account_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The account id of access.",
			},
			"access_iam_role": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "The iam role of access. If the VPC of another account is attached, " +
					"the other account needs to create a role with CFSCacheAccess permission, " +
					"and enter the role name as a parameter.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of subnet.",
			},
			"security_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of security group.",
			},
			"vpc_route_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether enable all vpc route.",
			},

			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of access.",
			},
			"created_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time.",
			},
			"is_default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether is default access.",
			},
			"access_service_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The service name of access.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of vpc.",
			},
			"access_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of access.",
			},
		},
	}
	return resource
}

func resourceVolcengineCloudfsAccessCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineCloudfsAccess())
	if err != nil {
		return fmt.Errorf("error on creating access %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudfsAccessRead(d, meta)
}

func resourceVolcengineCloudfsAccessRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineCloudfsAccess())
	if err != nil {
		return fmt.Errorf("error on reading access %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudfsAccessUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineCloudfsAccess())
	if err != nil {
		return fmt.Errorf("error on updating access %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudfsAccessRead(d, meta)
}

func resourceVolcengineCloudfsAccessDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineCloudfsAccess())
	if err != nil {
		return fmt.Errorf("error on deleting access %q, %s", d.Id(), err)
	}
	return err
}
