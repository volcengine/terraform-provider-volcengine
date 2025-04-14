package escloud_ip_white_list

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
EscloudIpWhiteList can be imported using the instance_id:type:component, e.g.
```
$ terraform import volcengine_escloud_ip_white_list.default resource_id
```

*/

func ResourceVolcengineEscloudIpWhiteList() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEscloudIpWhiteListCreate,
		Read:   resourceVolcengineEscloudIpWhiteListRead,
		Update: resourceVolcengineEscloudIpWhiteListUpdate,
		Delete: resourceVolcengineEscloudIpWhiteListDelete,
		Importer: &schema.ResourceImporter{
			State: esCloudIpWhiteListImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the EsCloud instance.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the ip white list. Valid values: `private`, `public`.",
			},
			"component": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The component of the ip white list. Valid values: `es`, `kibana`.",
			},
			"ip_list": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The ip list of the ip white list.",
			},
		},
	}
	return resource
}

func resourceVolcengineEscloudIpWhiteListCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEscloudIpWhiteListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineEscloudIpWhiteList())
	if err != nil {
		return fmt.Errorf("error on creating escloud_ip_white_list %q, %s", d.Id(), err)
	}
	return resourceVolcengineEscloudIpWhiteListRead(d, meta)
}

func resourceVolcengineEscloudIpWhiteListRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEscloudIpWhiteListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineEscloudIpWhiteList())
	if err != nil {
		return fmt.Errorf("error on reading escloud_ip_white_list %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEscloudIpWhiteListUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEscloudIpWhiteListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineEscloudIpWhiteList())
	if err != nil {
		return fmt.Errorf("error on updating escloud_ip_white_list %q, %s", d.Id(), err)
	}
	return resourceVolcengineEscloudIpWhiteListRead(d, meta)
}

func resourceVolcengineEscloudIpWhiteListDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEscloudIpWhiteListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineEscloudIpWhiteList())
	if err != nil {
		return fmt.Errorf("error on deleting escloud_ip_white_list %q, %s", d.Id(), err)
	}
	return err
}

var esCloudIpWhiteListImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 3 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("instance_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("type", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("component", items[2]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
