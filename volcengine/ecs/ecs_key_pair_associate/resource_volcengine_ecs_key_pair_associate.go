package ecs_key_pair_associate

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ECS key pair associate can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_key_pair_associate.default kp-ybti5tkpkv2udbfolrft:i-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVolcengineEcsKeyPairAssociate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEcsKeyPairAssociateCreate,
		Read:   resourceVolcengineEcsKeyPairAssociateRead,
		Delete: resourceVolcengineEcsKeyPairAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("key_pair_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("instance_id", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"key_pair_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of ECS KeyPair Associate.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of ECS Instance.",
			},
		},
	}
	return resource
}

func resourceVolcengineEcsKeyPairAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsKeyPairAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineEcsKeyPairAssociate())
	if err != nil {
		return fmt.Errorf("error on creating ecs key pair Associate %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsKeyPairAssociateRead(d, meta)
}

func resourceVolcengineEcsKeyPairAssociateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsKeyPairAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineEcsKeyPairAssociate())
	if err != nil {
		return fmt.Errorf("error on reading ecs key pair Associate %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEcsKeyPairAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsKeyPairAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineEcsKeyPairAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting ecs key pair Associate %q, %s", d.Id(), err)
	}
	return err
}