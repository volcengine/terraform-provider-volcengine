package cen_attach_instance

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Cen attach instance can be imported using the CenId:InstanceId:InstanceType:RegionId, e.g.
```
$ terraform import volcengine_cen_attach_instance.default cen-7qthudw0ll6jmc***:vpc-2fexiqjlgjif45oxruvso****:VPC:cn-beijing
```

*/

func ResourceVolcengineCenAttachInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCenAttachInstanceCreate,
		Read:   resourceVolcengineCenAttachInstanceRead,
		Delete: resourceVolcengineCenAttachInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: cenAttachInstanceImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"cen_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the cen.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the instance.",
			},
			"instance_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "The type of the instance.",
				ValidateFunc: validation.StringInSlice([]string{"VPC", "DCGW"}, false),
			},
			"instance_region_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The region ID of the instance.",
			},
			"instance_owner_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The owner ID of the instance.",
			},
		},
	}
	s := DataSourceVolcengineCenAttachInstances().Schema["attach_instances"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(s, &resource.Schema)
	return resource
}

func resourceVolcengineCenAttachInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	cenAttachInstanceService := NewCenAttachInstanceService(meta.(*ve.SdkClient))
	err = cenAttachInstanceService.Dispatcher.Create(cenAttachInstanceService, d, ResourceVolcengineCenAttachInstance())
	if err != nil {
		return fmt.Errorf("error on creating cen attach instance  %q, %s", d.Id(), err)
	}
	return resourceVolcengineCenAttachInstanceRead(d, meta)
}

func resourceVolcengineCenAttachInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	cenAttachInstanceService := NewCenAttachInstanceService(meta.(*ve.SdkClient))
	err = cenAttachInstanceService.Dispatcher.Read(cenAttachInstanceService, d, ResourceVolcengineCenAttachInstance())
	if err != nil {
		return fmt.Errorf("error on reading cen attach instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCenAttachInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	cenAttachInstanceService := NewCenAttachInstanceService(meta.(*ve.SdkClient))
	err = cenAttachInstanceService.Dispatcher.Delete(cenAttachInstanceService, d, ResourceVolcengineCenAttachInstance())
	if err != nil {
		return fmt.Errorf("error on deleting cen attach instance %q, %s", d.Id(), err)
	}
	return err
}
