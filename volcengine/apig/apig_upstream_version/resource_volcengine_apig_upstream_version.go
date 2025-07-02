package apig_upstream_version

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ApigUpstreamVersion can be imported using the id, e.g.
```
$ terraform import volcengine_apig_upstream_version.default resource_id
```

*/

func ResourceVolcengineApigUpstreamVersion() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineApigUpstreamVersionCreate,
		Read:   resourceVolcengineApigUpstreamVersionRead,
		Update: resourceVolcengineApigUpstreamVersionUpdate,
		Delete: resourceVolcengineApigUpstreamVersionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"upstream_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the apig upstream.",
			},
			"upstream_version": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The version of the apig upstream.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The name of apig upstream version.",
						},
						"labels": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The labels of apig upstream version.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The key of apig upstream version label.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The value of apig upstream version label.",
									},
								},
							},
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of apig upstream version.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineApigUpstreamVersionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigUpstreamVersionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineApigUpstreamVersion())
	if err != nil {
		return fmt.Errorf("error on creating apig_upstream_version %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigUpstreamVersionRead(d, meta)
}

func resourceVolcengineApigUpstreamVersionRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigUpstreamVersionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineApigUpstreamVersion())
	if err != nil {
		return fmt.Errorf("error on reading apig_upstream_version %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineApigUpstreamVersionUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigUpstreamVersionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineApigUpstreamVersion())
	if err != nil {
		return fmt.Errorf("error on updating apig_upstream_version %q, %s", d.Id(), err)
	}
	return resourceVolcengineApigUpstreamVersionRead(d, meta)
}

func resourceVolcengineApigUpstreamVersionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewApigUpstreamVersionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineApigUpstreamVersion())
	if err != nil {
		return fmt.Errorf("error on deleting apig_upstream_version %q, %s", d.Id(), err)
	}
	return err
}

var upstreamVersionImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("upstream_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("upstream_version.0.name", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
