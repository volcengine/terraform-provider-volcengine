package project

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Tls Project can be imported using the id, e.g.
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
			"iam_project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The IAM project name of the tls project.",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The region of the tls project.",
			},
			"tags": ve.TagsSchema(),
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
