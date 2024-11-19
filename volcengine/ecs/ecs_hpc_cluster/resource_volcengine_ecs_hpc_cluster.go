package ecs_hpc_cluster

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
EcsHpcCluster can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_hpc_cluster.default resource_id
```

*/

func ResourceVolcengineEcsHpcCluster() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEcsHpcClusterCreate,
		Read:   resourceVolcengineEcsHpcClusterRead,
		Update: resourceVolcengineEcsHpcClusterUpdate,
		Delete: resourceVolcengineEcsHpcClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The zone id of the hpc cluster.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the hpc cluster.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the hpc cluster.",
			},
		},
	}
	return resource
}

func resourceVolcengineEcsHpcClusterCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsHpcClusterService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineEcsHpcCluster())
	if err != nil {
		return fmt.Errorf("error on creating ecs_hpc_cluster %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsHpcClusterRead(d, meta)
}

func resourceVolcengineEcsHpcClusterRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsHpcClusterService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineEcsHpcCluster())
	if err != nil {
		return fmt.Errorf("error on reading ecs_hpc_cluster %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEcsHpcClusterUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsHpcClusterService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineEcsHpcCluster())
	if err != nil {
		return fmt.Errorf("error on updating ecs_hpc_cluster %q, %s", d.Id(), err)
	}
	return resourceVolcengineEcsHpcClusterRead(d, meta)
}

func resourceVolcengineEcsHpcClusterDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEcsHpcClusterService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineEcsHpcCluster())
	if err != nil {
		return fmt.Errorf("error on deleting ecs_hpc_cluster %q, %s", d.Id(), err)
	}
	return err
}
