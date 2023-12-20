package cloud_monitor_contact_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudMonitorContactGroup can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_monitor_contact_group.default resource_id
```

*/

func ResourceVolcengineCloudMonitorContactGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudMonitorContactGroupCreate,
		Read:   resourceVolcengineCloudMonitorContactGroupRead,
		Update: resourceVolcengineCloudMonitorContactGroupUpdate,
		Delete: resourceVolcengineCloudMonitorContactGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the contact group.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the contact group.",
			},
			"contacts_id_list": {
				Type:     schema.TypeSet,
				Set:      schema.HashString,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "When creating a contact group, contacts should be added with their contact ID. " +
					"The maximum number of IDs allowed is 10, meaning that the maximum number of members in a single contact group is 10.",
			},
		},
	}
	return resource
}

func resourceVolcengineCloudMonitorContactGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudMonitorContactGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCloudMonitorContactGroup())
	if err != nil {
		return fmt.Errorf("error on creating cloud_monitor_contact_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudMonitorContactGroupRead(d, meta)
}

func resourceVolcengineCloudMonitorContactGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudMonitorContactGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCloudMonitorContactGroup())
	if err != nil {
		return fmt.Errorf("error on reading cloud_monitor_contact_group %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudMonitorContactGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudMonitorContactGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCloudMonitorContactGroup())
	if err != nil {
		return fmt.Errorf("error on updating cloud_monitor_contact_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudMonitorContactGroupRead(d, meta)
}

func resourceVolcengineCloudMonitorContactGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudMonitorContactGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCloudMonitorContactGroup())
	if err != nil {
		return fmt.Errorf("error on deleting cloud_monitor_contact_group %q, %s", d.Id(), err)
	}
	return err
}
