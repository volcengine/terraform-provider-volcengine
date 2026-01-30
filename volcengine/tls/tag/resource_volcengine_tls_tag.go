package tag

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
tls tag can be imported using the resource_id:resource_type, e.g.
```
$ terraform import volcengine_tls_tag.default resource-123456:project
```

*/

func ResourceVolcengineTlsTag() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineTlsTagCreate,
		Read:   resourceVolcengineTlsTagRead,
		Delete: resourceVolcengineTlsTagDelete,
		Importer: &schema.ResourceImporter{
			State: resourceVolcengineTlsTagImport,
		},
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the resource.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the resource. Valid values: project, topic, shipper, host_group, host, consumer_group, rule, alarm, alarm_notify_group, etl_task, import_task, schedule_sql_task, download_task, trace_instance.",
			},
			"tags": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Tags. The tag key must be unique within a resource, and the same tag key is not allowed to be repeated. The tag key must be 1 to 128 characters long, and can contain letters, digits, spaces, and the following special characters: _.:/=+-@. The tag value can be empty and must be 0 to 256 characters long, and can contain letters, digits, spaces, and the following special characters: _.:/=+-@.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Key of Tags.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Value of Tags.",
						},
					},
				},
			},
		},
	}
}

func resourceVolcengineTlsTagCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsTagService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Create(service, d, ResourceVolcengineTlsTag())
}

func resourceVolcengineTlsTagRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsTagService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Read(service, d, ResourceVolcengineTlsTag())
}

func resourceVolcengineTlsTagDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsTagService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineTlsTag())
}

func resourceVolcengineTlsTagImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), ":")
	if len(idParts) != 2 {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected resource_id:resource_type", d.Id())
	}

	resourceId := idParts[0]
	resourceType := idParts[1]

	d.Set("resource_id", resourceId)
	d.Set("resource_type", resourceType)
	d.SetId(resourceId)

	return []*schema.ResourceData{d}, nil
}
