package bandwidth_package_attachment

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
BandwidthPackageAttachment can be imported using the bandwidth package id and eip id, e.g.
```
$ terraform import volcengine_bandwidth_package_attachment.default BandwidthPackageId:EipId
```

*/

func ResourceVolcengineBandwidthPackageAttachment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineBandwidthPackageAttachmentCreate,
		Read:   resourceVolcengineBandwidthPackageAttachmentRead,
		Delete: resourceVolcengineBandwidthPackageAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				var err error
				items := strings.Split(d.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{d}, fmt.Errorf("import id must be of the form BandwidthPackageId:EipId")
				}
				err = d.Set("bandwidth_package_id", items[0])
				if err != nil {
					return []*schema.ResourceData{d}, err
				}
				err = d.Set("allocation_id", items[1])
				if err != nil {
					return []*schema.ResourceData{d}, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"bandwidth_package_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The bandwidth package id.",
			},
			"allocation_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the public IP or IPv6 public bandwidth to be added to the shared bandwidth package instance.",
			},
		},
	}
	return resource
}

func resourceVolcengineBandwidthPackageAttachmentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewBandwidthPackageAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineBandwidthPackageAttachment())
	if err != nil {
		return fmt.Errorf("error on creating bandwidth_package_attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineBandwidthPackageAttachmentRead(d, meta)
}

func resourceVolcengineBandwidthPackageAttachmentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewBandwidthPackageAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineBandwidthPackageAttachment())
	if err != nil {
		return fmt.Errorf("error on reading bandwidth_package_attachment %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineBandwidthPackageAttachmentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewBandwidthPackageAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineBandwidthPackageAttachment())
	if err != nil {
		return fmt.Errorf("error on deleting bandwidth_package_attachment %q, %s", d.Id(), err)
	}
	return err
}
