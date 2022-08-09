package user

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CR User can be imported using the CR instance name, e.g.
```
$ terraform import volcenginec_cr_user.default cr-basic
```

*/

func ResourceVolcengineCrUser() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCrUserCreate,
		Read:   resourceVolcengineCrUserRead,
		Update: resourceVolcengineCrUserUpdate,
		Delete: resourceVolcengineCrUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"registry": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The registry name that will set user.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user password.",
			},
		},
	}

	dataSource := DataSourceVolcengineCrUsers().Schema["users"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineCrUserCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCrUser())
	if err != nil {
		return fmt.Errorf("Error on creating CR user %q,%s", d.Id(), err)
	}
	return resourceVolcengineCrUserRead(d, meta)
}

func resourceVolcengineCrUserUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCrUser())
	if err != nil {
		return fmt.Errorf("error on updating CR user  %q, %s", d.Id(), err)
	}
	return resourceVolcengineCrUserRead(d, meta)
}

func resourceVolcengineCrUserDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCrUser())
	if err != nil {
		return fmt.Errorf("error on deleting CR user %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCrUserRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCrUser())
	if err != nil {
		return fmt.Errorf("error on reading CR user %q,%s", d.Id(), err)
	}
	return err
}
