package cloud_identity_user_attachment

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudIdentityUserAttachment can be imported using the group_id:user_id, e.g.
```
$ terraform import volcengine_cloud_identity_user_attachment.default resource_id
```

*/

func ResourceVolcengineCloudIdentityUserAttachment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudIdentityUserAttachmentCreate,
		Read:   resourceVolcengineCloudIdentityUserAttachmentRead,
		Delete: resourceVolcengineCloudIdentityUserAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: userAttachmentImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the cloud identity user.",
			},
			"group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the cloud identity group.",
			},
		},
	}
	return resource
}

func resourceVolcengineCloudIdentityUserAttachmentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityUserAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCloudIdentityUserAttachment())
	if err != nil {
		return fmt.Errorf("error on creating cloud_identity_user_attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudIdentityUserAttachmentRead(d, meta)
}

func resourceVolcengineCloudIdentityUserAttachmentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityUserAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCloudIdentityUserAttachment())
	if err != nil {
		return fmt.Errorf("error on reading cloud_identity_user_attachment %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudIdentityUserAttachmentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityUserAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCloudIdentityUserAttachment())
	if err != nil {
		return fmt.Errorf("error on deleting cloud_identity_user_attachment %q, %s", d.Id(), err)
	}
	return err
}

var userAttachmentImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("group_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("user_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
