package instance

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Instance can be imported using the id, e.g.
```
$ terraform import volcengine_veenedge_instance.default veenn769ewmjjqyqh5dv
```

*/

func ResourceVolcengineInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineInstanceCreate,
		Read:   resourceVolcengineInstanceRead,
		Delete: resourceVolcengineInstanceDelete,
		Update: resourceVolcengineInstanceUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Import an exist instance, usually for import a default instance generated with cloud server creating.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of instance, only effected in update scene.",
			},
			"secret_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"KeyPair", "Password"}, false),
				Description:  "The type of secret, only effected in update scene. The value can be `KeyPair` or `Password`.",
			},
			"secret_data": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The data of secret, only effected in update scene.",
			},
			"cloudserver_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The id of cloud server.",
			},
			"area_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The area name.",
			},
			"isp": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The isp info.",
			},
			"default_isp": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The default isp for multi line node.",
			},
			"cluster_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The name of cluster.",
			},
		},
	}

	return resource
}

func resourceVolcengineInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineInstance())
	if err != nil {
		return fmt.Errorf(" Error on creating instance %q,%s", d.Id(), err)
	}
	return resourceVolcengineInstanceRead(d, meta)
}

func resourceVolcengineInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineInstance())
	if err != nil {
		return fmt.Errorf("error on deleting instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineInstance())
	if err != nil {
		return fmt.Errorf("error on reading instance %q,%s", d.Id(), err)
	}
	return err
}

func resourceVolcengineInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewInstanceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineInstance())
	if err != nil {
		return fmt.Errorf("error on updating instance %q, %s", d.Id(), err)
	}
	return resourceVolcengineInstanceRead(d, meta)
}
