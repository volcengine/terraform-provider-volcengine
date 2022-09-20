package ecs_key_pair

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ECS key pair can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_key_pair.default kp-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVolcengineEcsKeyPair() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEcsKeyPairCreate,
		Read:   resourceVolcengineEcsKeyPairRead,
		Update: resourceVolcengineEcsKeyPairUpdate,
		Delete: resourceVolcengineEcsKeyPairDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"key_pair_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(2, 64),
				Description:  "The name of key pair.",
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				StateFunc: func(v interface{}) string {
					switch ele := v.(type) {
					case string:
						return strings.TrimSpace(ele)
					default:
						return ""
					}
				},
				Description: "Public key string.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of key pair.",
			},
			"key_file": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "Target file to save private key. It is recommended that the value not be empty. " +
					"You only have one chance to download the private key, the volcengine will not save your private key, please keep it safe. " +
					"In the TF import scenario, this field will not write the private key locally.",
			},
			"finger_print": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The finger print info.",
			},
			"key_pair_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of key pair.",
			},
		},
	}
	return resource
}

func resourceVolcengineEcsKeyPairCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsKeyPairService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineEcsKeyPair())
	if err != nil {
		return fmt.Errorf("error on creating ecs key pair  %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsKeyPairRead(d, meta)
}

func resourceVolcengineEcsKeyPairRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsKeyPairService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineEcsKeyPair())
	if err != nil {
		return fmt.Errorf("error on reading ecs key pair %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEcsKeyPairUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsKeyPairService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineEcsKeyPair())
	if err != nil {
		return fmt.Errorf("error on updating ecs key pair  %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsKeyPairRead(d, meta)
}

func resourceVolcengineEcsKeyPairDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsKeyPairService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineEcsKeyPair())
	if err != nil {
		return fmt.Errorf("error on deleting ecs key pair %q, %s", d.Id(), err)
	}
	return err
}
