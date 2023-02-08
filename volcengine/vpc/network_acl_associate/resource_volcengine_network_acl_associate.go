package network_acl_associate

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"time"
)

/*

Import
NetworkAcl associate can be imported using the network_acl_id:resource_id, e.g.
```
$ terraform import volcengine_network_acl_associate.default nacl-172leak37mi9s4d1w33pswqkh:subnet-637jxq81u5mon3gd6ivc7rj
```

*/

func ResourceVolcengineNetworkAclAssociate() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineAclAssociateCreate,
		Read:   resourceVolcengineAclAssociateRead,
		Delete: resourceVolcengineAclAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: aclAssociateImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"network_acl_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of Network Acl.",
			},
			"resource_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The resource id of Network Acl.",
			},
		},
	}
}

func resourceVolcengineAclAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	aclAssociateService := NewNetworkAclAssociateService(meta.(*ve.SdkClient))
	err = aclAssociateService.Dispatcher.Create(aclAssociateService, d, ResourceVolcengineNetworkAclAssociate())
	if err != nil {
		return fmt.Errorf("error on creating acl Associate %q, %w", d.Id(), err)
	}
	return resourceVolcengineAclAssociateRead(d, meta)
}

func resourceVolcengineAclAssociateRead(d *schema.ResourceData, meta interface{}) (err error) {
	aclAssociateService := NewNetworkAclAssociateService(meta.(*ve.SdkClient))
	err = aclAssociateService.Dispatcher.Read(aclAssociateService, d, ResourceVolcengineNetworkAclAssociate())
	if err != nil {
		return fmt.Errorf("error on reading acl Associate %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineAclAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	aclAssociateService := NewNetworkAclAssociateService(meta.(*ve.SdkClient))
	err = aclAssociateService.Dispatcher.Delete(aclAssociateService, d, ResourceVolcengineNetworkAclAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting acl Associate %q, %w", d.Id(), err)
	}
	return err
}
