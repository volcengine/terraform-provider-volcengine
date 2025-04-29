package cr_endpoint_acl_policy

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CrEndpointAclPolicy can be imported using the registry:entry, e.g.
```
$ terraform import volcengine_cr_endpoint_acl_policy.default resource_id
```

*/

func ResourceVolcengineCrEndpointAclPolicy() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCrEndpointAclPolicyCreate,
		Read:   resourceVolcengineCrEndpointAclPolicyRead,
		Delete: resourceVolcengineCrEndpointAclPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: crEndpointAclPolicyImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"registry": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The registry name.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the acl policy. Valid values: `Public`.",
			},
			"entry": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ip list of the acl policy.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The description of the acl policy.",
			},
		},
	}
	return resource
}

func resourceVolcengineCrEndpointAclPolicyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrEndpointAclPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCrEndpointAclPolicy())
	if err != nil {
		return fmt.Errorf("error on creating cr_endpoint_acl_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineCrEndpointAclPolicyRead(d, meta)
}

func resourceVolcengineCrEndpointAclPolicyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrEndpointAclPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCrEndpointAclPolicy())
	if err != nil {
		return fmt.Errorf("error on reading cr_endpoint_acl_policy %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCrEndpointAclPolicyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrEndpointAclPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCrEndpointAclPolicy())
	if err != nil {
		return fmt.Errorf("error on updating cr_endpoint_acl_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineCrEndpointAclPolicyRead(d, meta)
}

func resourceVolcengineCrEndpointAclPolicyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrEndpointAclPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCrEndpointAclPolicy())
	if err != nil {
		return fmt.Errorf("error on deleting cr_endpoint_acl_policy %q, %s", d.Id(), err)
	}
	return err
}

func crEndpointAclPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'registry:entry'")
	}
	if err := d.Set("registry", items[0]); err != nil {
		return []*schema.ResourceData{d}, err
	}
	if err := d.Set("entry", items[1]); err != nil {
		return []*schema.ResourceData{d}, err
	}
	return []*schema.ResourceData{d}, nil
}
