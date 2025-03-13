package veecp_kubeconfig

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VeecpKubeconfig can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_kubeconfig.default resource_id
```

*/

func ResourceVolcengineVeecpKubeconfig() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVeecpKubeconfigCreate,
		Read:   resourceVolcengineVeecpKubeconfigRead,
		Delete: resourceVolcengineVeecpKubeconfigDelete,
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
	return resource
}

func resourceVolcengineVeecpKubeconfigCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpKubeconfigService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVeecpKubeconfig())
	if err != nil {
		return fmt.Errorf("error on creating veecp_kubeconfig %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpKubeconfigRead(d, meta)
}

func resourceVolcengineVeecpKubeconfigRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpKubeconfigService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVeecpKubeconfig())
	if err != nil {
		return fmt.Errorf("error on reading veecp_kubeconfig %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVeecpKubeconfigDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpKubeconfigService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVeecpKubeconfig())
	if err != nil {
		return fmt.Errorf("error on deleting veecp_kubeconfig %q, %s", d.Id(), err)
	}
	return err
}
