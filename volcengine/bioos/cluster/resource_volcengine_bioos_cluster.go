package cluster

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Cluster can be imported using the id, e.g.
```
$ terraform import volcengine_bioos_cluster.default *****
```

*/

func ResourceVolcengineBioosCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineBioosClusterCreate,
		Read:   resourceVolcengineBioosClusterRead,
		Delete: resourceVolcengineBioosClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the cluster.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The description of the cluster.",
			},
			"vke_config": {
				Type: schema.TypeList,
				Description: "The configuration of the vke cluster. This cluster type is not recommended. " +
					"It is recommended to use a shared cluster.",
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				MaxItems:      1,
				ConflictsWith: []string{"shared_config"},
				AtLeastOneOf:  []string{"shared_config"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The id of the vke cluster.",
						},
						"storage_class": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The name of the StorageClass that the vke cluster has installed.",
						},
					},
				},
			},
			"shared_config": {
				Type:          schema.TypeList,
				Description:   "The configuration of the shared cluster.",
				Optional:      true,
				ForceNew:      true,
				MaxItems:      1,
				Computed:      true,
				ConflictsWith: []string{"vke_config"},
				AtLeastOneOf:  []string{"vke_config"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: "Whether to enable a shared cluster. This value must be `true`.",
						},
					},
				},
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the bioos cluster.",
			},
		},
	}
}

func resourceVolcengineBioosClusterCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineBioosClusterService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineBioosCluster())
	if err != nil {
		return fmt.Errorf("error on creating volcengine bioos cluster: %q, %w", d.Id(), err)
	}
	return resourceVolcengineBioosClusterRead(d, meta)
}

func resourceVolcengineBioosClusterRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineBioosClusterService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineBioosCluster())
	if err != nil {
		return fmt.Errorf("error on reading volcengine bioos cluster: %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineBioosClusterDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineBioosClusterService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineBioosCluster())
	if err != nil {
		return fmt.Errorf("error on deleting volcengine bioos cluster: %q, %w", d.Id(), err)
	}
	return nil
}
