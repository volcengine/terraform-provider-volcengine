package alb_customized_cfg

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
AlbCustomizedCfg can be imported using the id, e.g.
```
$ terraform import volcengine_alb_customized_cfg.default ccfg-3cj44nv0jhhxc6c6rrtet****
```

*/

func ResourceVolcengineAlbCustomizedCfg() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineAlbCustomizedCfgCreate,
		Read:   resourceVolcengineAlbCustomizedCfgRead,
		Update: resourceVolcengineAlbCustomizedCfgUpdate,
		Delete: resourceVolcengineAlbCustomizedCfgDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"customized_cfg_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of CustomizedCfg.",
			},
			"customized_cfg_content": {
				Type:     schema.TypeString,
				Required: true,
				Description: "The content of CustomizedCfg. The length cannot exceed 4096 characters. " +
					"Spaces and semicolons need to be escaped. " +
					"Currently supported configuration items are `ssl_protocols`, `ssl_ciphers`, `client_max_body_size`, `keepalive_timeout`, `proxy_request_buffering` and `proxy_connect_timeout`.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of CustomizedCfg.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the CustomizedCfg.",
			},
			"tags": ve.TagsSchema(),
		},
	}
	return resource
}

func resourceVolcengineAlbCustomizedCfgCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbCustomizedCfgService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineAlbCustomizedCfg())
	if err != nil {
		return fmt.Errorf("error on creating alb_customized_cfg %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbCustomizedCfgRead(d, meta)
}

func resourceVolcengineAlbCustomizedCfgRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbCustomizedCfgService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineAlbCustomizedCfg())
	if err != nil {
		return fmt.Errorf("error on reading alb_customized_cfg %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineAlbCustomizedCfgUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbCustomizedCfgService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineAlbCustomizedCfg())
	if err != nil {
		return fmt.Errorf("error on updating alb_customized_cfg %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbCustomizedCfgRead(d, meta)
}

func resourceVolcengineAlbCustomizedCfgDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbCustomizedCfgService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineAlbCustomizedCfg())
	if err != nil {
		return fmt.Errorf("error on deleting alb_customized_cfg %q, %s", d.Id(), err)
	}
	return err
}
