package project

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TlsProject can be imported using the id, e.g.
```
$ terraform import volcengine_tls_project.default e020c978-4f05-40e1-9167-0113d3ef****
```

*/

func ResourceVolcengineTlsProject() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTlsProjectCreate,
		Read:   resourceVolcengineTlsProjectRead,
		Update: resourceVolcengineTlsProjectUpdate,
		Delete: resourceVolcengineTlsProjectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"project_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the tls project.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the tls project.",
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the tls project.",
			},
			"inner_net_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The inner net domain of the tls project.",
			},
			"topic_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The count of topics in the tls project.",
			},
		},
	}
	return resource
}

func resourceVolcengineTlsProjectCreate(d *schema.ResourceData, meta interface{}) (err error) {
	tlsProjectService := NewTlsProjectService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(tlsProjectService, d, ResourceVolcengineTlsProject())
	if err != nil {
		return fmt.Errorf("error on creating TlsProject %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsProjectRead(d, meta)
}

func resourceVolcengineTlsProjectRead(d *schema.ResourceData, meta interface{}) (err error) {
	tlsProjectService := NewTlsProjectService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(tlsProjectService, d, ResourceVolcengineTlsProject())
	if err != nil {
		return fmt.Errorf("error on reading TlsProject %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineTlsProjectUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	tlsProjectService := NewTlsProjectService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(tlsProjectService, d, ResourceVolcengineTlsProject())
	if err != nil {
		return fmt.Errorf("error on updating TlsProject %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsProjectRead(d, meta)
}

func resourceVolcengineTlsProjectDelete(d *schema.ResourceData, meta interface{}) (err error) {
	tlsProjectService := NewTlsProjectService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(tlsProjectService, d, ResourceVolcengineTlsProject())
	if err != nil {
		return fmt.Errorf("error on deleting TlsProject %q, %w", d.Id(), err)
	}
	return err
}
