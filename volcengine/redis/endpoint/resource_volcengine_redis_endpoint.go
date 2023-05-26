package endpoint

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Redis Endpoint can be imported using the instanceId:eipId, e.g.
```
$ terraform import volcengine_redis_endpoint.default redis-asdljioeixxxx:eip-2fef2qcfbfw8w5oxruw3w****
```
*/

func ResourceVolcengineRedisEndpoint() *schema.Resource {
	resource := &schema.Resource{
		Read:   resourceVolcengineRedisEndpointRead,
		Create: resourceVolcengineRedisEndpointCreate,
		Delete: resourceVolcengineRedisEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: redisEndpointAssociateImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of instance.",
			},
			"eip_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of eip.",
			},
		},
	}

	return resource
}

func resourceVolcengineRedisEndpointRead(d *schema.ResourceData, meta interface{}) (err error) {
	redisEndpointService := NewRedisEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(redisEndpointService, d, ResourceVolcengineRedisEndpoint())
	if err != nil {
		return fmt.Errorf("error on reading redis endpoint %v, %v", d.Id(), err)
	}
	return nil
}

func resourceVolcengineRedisEndpointCreate(d *schema.ResourceData, meta interface{}) (err error) {
	redisEndpointService := NewRedisEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(redisEndpointService, d, ResourceVolcengineRedisEndpoint())
	if err != nil {
		return fmt.Errorf("error on creating redis endpoint %v, %v", d.Id(), err)
	}
	return resourceVolcengineRedisEndpointRead(d, meta)
}

func resourceVolcengineRedisEndpointDelete(d *schema.ResourceData, meta interface{}) (err error) {
	redisEndpointService := NewRedisEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(redisEndpointService, d, ResourceVolcengineRedisEndpoint())
	if err != nil {
		return fmt.Errorf("error on deleting redis endpoint %q, %s", d.Id(), err)
	}
	return err
}
