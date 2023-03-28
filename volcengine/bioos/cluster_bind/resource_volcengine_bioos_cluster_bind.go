package cluster_bind

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Cluster binder can be imported using the workspace id and cluster id, e.g.
```
$ terraform import volcengine_bioos_cluster_bind.default wc*****:uc***
```

*/

func ResourceVolcengineBioosClusterBind() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineBioosClusterBindCreate,
		Delete: resourceVolcengineBioosClusterDelete,
		Read:   resourceVolcengineBioosClusterBindRead,
		Importer: &schema.ResourceImporter{
			State: importBioosClusterBind,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the workspace.",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the cluster.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the cluster bind.",
				ValidateFunc: validation.StringInSlice([]string{
					"workflow",
					"notebook",
				}, false),
			},
		},
	}
}

func resourceVolcengineBioosClusterBindCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineBioosClusterBindService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineBioosClusterBind())
	if err != nil {
		return fmt.Errorf("error on creating volcengine bioos cluster bind: %q, %w", d.Id(), err)
	}
	return resourceVolcengineBioosClusterBindRead(d, meta)
}

func resourceVolcengineBioosClusterBindRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineBioosClusterBindService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineBioosClusterBind())
	if err != nil {
		return fmt.Errorf("error on reading volcengine bioos cluster bind: %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineBioosClusterDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineBioosClusterBindService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineBioosClusterBind())
	if err != nil {
		return fmt.Errorf("error on deleting volcengine bioos cluster bind: %q, %w", d.Id(), err)
	}
	return nil
}

func importBioosClusterBind(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form ID:ClusterID")
	}
	err = data.Set("workspace_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("cluster_id", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
