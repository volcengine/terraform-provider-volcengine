package cen_grant_instance

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Cen grant instance can be imported using the CenId:CenOwnerId:InstanceId:InstanceType:RegionId, e.g.
```
$ terraform import vestack_cen_grant_instance.default cen-7qthudw0ll6jmc***:210000****:vpc-2fexiqjlgjif45oxruvso****:VPC:cn-beijing
```

*/

func ResourceVestackCenGrantInstance() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackCenGrantInstanceCreate,
		Read:   resourceVestackCenGrantInstanceRead,
		Delete: resourceVestackCenGrantInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: cenGrantInstanceImporter,
		},
		Schema: map[string]*schema.Schema{
			"cen_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the cen.",
			},
			"cen_owner_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The owner ID of the cen.",
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
		},
	}
	return resource
}

func resourceVestackCenGrantInstanceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	grantInstanceService := NewCenGrantInstanceService(meta.(*ve.SdkClient))
	err = grantInstanceService.Dispatcher.Create(grantInstanceService, d, ResourceVestackCenGrantInstance())
	if err != nil {
		return fmt.Errorf("error on creating cen grant instance  %q, %s", d.Id(), err)
	}
	return resourceVestackCenGrantInstanceRead(d, meta)
}

func resourceVestackCenGrantInstanceRead(d *schema.ResourceData, meta interface{}) (err error) {
	grantInstanceService := NewCenGrantInstanceService(meta.(*ve.SdkClient))
	err = grantInstanceService.Dispatcher.Read(grantInstanceService, d, ResourceVestackCenGrantInstance())
	if err != nil {
		return fmt.Errorf("error on reading cen grant instance %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackCenGrantInstanceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	grantInstanceService := NewCenGrantInstanceService(meta.(*ve.SdkClient))
	err = grantInstanceService.Dispatcher.Delete(grantInstanceService, d, ResourceVestackCenGrantInstance())
	if err != nil {
		return fmt.Errorf("error on deleting cen grant instance %q, %s", d.Id(), err)
	}
	return err
}
