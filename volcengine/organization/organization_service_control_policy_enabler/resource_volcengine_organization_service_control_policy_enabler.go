package organization_service_control_policy_enabler

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ServiceControlPolicy enabler can be imported using the default_id (organization:service_control_policy_enable) , e.g.
```
$ terraform import volcengine_organization_service_control_policy_enabler.default organization:service_control_policy_enable
```

*/

func ResourceVolcengineOrganizationServiceControlPolicyEnabler() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineOrganizationServiceControlPolicyEnablerCreate,
		Read:   resourceVolcengineOrganizationServiceControlPolicyEnablerRead,
		Delete: resourceVolcengineOrganizationServiceControlPolicyEnablerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{},
	}
	return resource
}

func resourceVolcengineOrganizationServiceControlPolicyEnablerCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineOrganizationServiceControlPolicyEnabler())
	if err != nil {
		return fmt.Errorf("error on creating organization_service_control_policy_enabler: %q, %s", d.Id(), err)
	}
	return resourceVolcengineOrganizationServiceControlPolicyEnablerRead(d, meta)
}

func resourceVolcengineOrganizationServiceControlPolicyEnablerRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineOrganizationServiceControlPolicyEnabler())
	if err != nil {
		return fmt.Errorf("error on reading organization_service_control_policy_enabler: %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineOrganizationServiceControlPolicyEnablerDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineOrganizationServiceControlPolicyEnabler())
	if err != nil {
		return fmt.Errorf("erron on deleting organization_service_control_policy_enabler: %q, %s", d.Id(), err)
	}
	return err
}
