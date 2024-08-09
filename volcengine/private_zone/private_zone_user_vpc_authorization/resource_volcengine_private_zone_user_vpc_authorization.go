package private_zone_user_vpc_authorization

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
PrivateZoneUserVpcAuthorization can be imported using the id, e.g.
```
$ terraform import volcengine_private_zone_user_vpc_authorization.default resource_id
```

*/

func ResourceVolcenginePrivateZoneUserVpcAuthorization() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivateZoneUserVpcAuthorizationCreate,
		Read:   resourceVolcenginePrivateZoneUserVpcAuthorizationRead,
		Delete: resourceVolcenginePrivateZoneUserVpcAuthorizationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The account Id which authorizes the private zone resource.",
			},
		},
	}
	return resource
}

func resourceVolcenginePrivateZoneUserVpcAuthorizationCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneUserVpcAuthorizationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcenginePrivateZoneUserVpcAuthorization())
	if err != nil {
		return fmt.Errorf("error on creating private_zone_user_vpc_authorization %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneUserVpcAuthorizationRead(d, meta)
}

func resourceVolcenginePrivateZoneUserVpcAuthorizationRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneUserVpcAuthorizationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcenginePrivateZoneUserVpcAuthorization())
	if err != nil {
		return fmt.Errorf("error on reading private_zone_user_vpc_authorization %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcenginePrivateZoneUserVpcAuthorizationUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneUserVpcAuthorizationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcenginePrivateZoneUserVpcAuthorization())
	if err != nil {
		return fmt.Errorf("error on updating private_zone_user_vpc_authorization %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneUserVpcAuthorizationRead(d, meta)
}

func resourceVolcenginePrivateZoneUserVpcAuthorizationDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneUserVpcAuthorizationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcenginePrivateZoneUserVpcAuthorization())
	if err != nil {
		return fmt.Errorf("error on deleting private_zone_user_vpc_authorization %q, %s", d.Id(), err)
	}
	return err
}
