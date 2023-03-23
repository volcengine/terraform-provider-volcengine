package kubeconfig

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"time"
)

/*

Import
VkeKubeconfig can be imported using the id, e.g.
```
$ terraform import volcengine_vke_kubeconfig.default kce8simvqtofl0l6u4qd0
```

*/

func ResourceVolcengineVkeKubeconfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineVkeKubeconfigCreate,
		Read:   resourceVolcengineVkeKubeconfigRead,
		Delete: resourceVolcengineVkeKubeconfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The cluster id of the Kubeconfig.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the Kubeconfig, the value of type should be Public or Private.",
			},
			"valid_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Default:     26280,
				Description: "The ValidDuration of the Kubeconfig, the range of the ValidDuration is 1 hour to 43800 hour.",
			},
		},
	}
}

func resourceVolcengineVkeKubeconfigCreate(d *schema.ResourceData, meta interface{}) (err error) {
	kubeconfigService := NewVkeKubeconfigService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(kubeconfigService, d, ResourceVolcengineVkeKubeconfig())
	if err != nil {
		return fmt.Errorf("error on creating cluster  %q, %w", d.Id(), err)
	}
	return resourceVolcengineVkeKubeconfigRead(d, meta)
}

func resourceVolcengineVkeKubeconfigRead(d *schema.ResourceData, meta interface{}) (err error) {
	kubeconfigService := NewVkeKubeconfigService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(kubeconfigService, d, ResourceVolcengineVkeKubeconfig())
	if err != nil {
		return fmt.Errorf("error on reading cluster %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineVkeKubeconfigDelete(d *schema.ResourceData, meta interface{}) (err error) {
	kubeconfigService := NewVkeKubeconfigService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(kubeconfigService, d, ResourceVolcengineVkeKubeconfig())
	if err != nil {
		return fmt.Errorf("error on deleting cluster %q, %w", d.Id(), err)
	}
	return err
}