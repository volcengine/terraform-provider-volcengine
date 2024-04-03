package security_group

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
PrivateLink Security Group Service can be imported using the endpoint id and security group id, e.g.
```
$ terraform import volcengine_privatelink_security_group.default ep-2fe630gurkl37k5gfuy33****:sg-xxxxx
```

*/

func ResourceVolcenginePrivatelinkSecurityGroupService() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivatelinkSecurityGroupCreate,
		Read:   resourceVolcenginePrivatelinkSecurityGroupRead,
		Delete: resourceVolcenginePrivatelinkSecurityGroupDelete,
		Importer: &schema.ResourceImporter{
			State: sgImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"endpoint_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the endpoint.",
			},
			"security_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The id of the security group. " +
					"It is not recommended to use this resource for binding security groups, it is recommended to use the `security_group_id` field of `volcengine_privatelink_vpc_endpoint` for binding.\n" +
					"If using this resource and `volcengine_privatelink_vpc_endpoint` jointly for operations, use lifecycle ignore_changes to suppress changes to the `security_group_id` field in `volcengine_privatelink_vpc_endpoint`.",
			},
		},
	}
	return resource
}

func resourceVolcenginePrivatelinkSecurityGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateLinkSecurityGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcenginePrivatelinkSecurityGroupService())
	if err != nil {
		return fmt.Errorf("error on creating private link security group service %q, %w", d.Id(), err)
	}
	return resourceVolcenginePrivatelinkSecurityGroupRead(d, meta)
}

func resourceVolcenginePrivatelinkSecurityGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateLinkSecurityGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcenginePrivatelinkSecurityGroupService())
	if err != nil {
		return fmt.Errorf("error on reading private link security group service %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcenginePrivatelinkSecurityGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateLinkSecurityGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcenginePrivatelinkSecurityGroupService())
	if err != nil {
		return fmt.Errorf("error on deleting private link security group service %q, %w", d.Id(), err)
	}
	return nil
}

var sgImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("endpoint_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("security_group_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
