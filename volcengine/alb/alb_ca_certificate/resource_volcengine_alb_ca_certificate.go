package alb_ca_certificate

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
AlbCaCertificate can be imported using the id, e.g.
```
$ terraform import volcengine_alb_ca_certificate.default cert-*****
```

*/

func ResourceVolcengineAlbCaCertificate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineAlbCaCertificateCreate,
		Read:   resourceVolcengineAlbCaCertificateRead,
		Update: resourceVolcengineAlbCaCertificateUpdate,
		Delete: resourceVolcengineAlbCaCertificateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"ca_certificate_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the CA certificate.",
			},
			"ca_certificate": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The content of the CA certificate.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the CA certificate.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the CA certificate.",
			},
			// "tags": ve.TagsSchema(),
			// computed fields
			"san": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The san extension of the Certificate.",
			},
			"listeners": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "The ID list of the Listener.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the CA Certificate.",
			},
			"expired_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The expire time of the CA Certificate.",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The domain name of the CA Certificate.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the CA Certificate.",
			},
			// 文档与接口实际返回不同
			"certificate_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the CA Certificate.",
			},
		},
	}
	return resource
}

func resourceVolcengineAlbCaCertificateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbCaCertificateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineAlbCaCertificate())
	if err != nil {
		return fmt.Errorf("error on creating alb_ca_certificate %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbCaCertificateRead(d, meta)
}

func resourceVolcengineAlbCaCertificateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbCaCertificateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineAlbCaCertificate())
	if err != nil {
		return fmt.Errorf("error on reading alb_ca_certificate %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineAlbCaCertificateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbCaCertificateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineAlbCaCertificate())
	if err != nil {
		return fmt.Errorf("error on updating alb_ca_certificate %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbCaCertificateRead(d, meta)
}

func resourceVolcengineAlbCaCertificateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbCaCertificateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineAlbCaCertificate())
	if err != nil {
		return fmt.Errorf("error on deleting alb_ca_certificate %q, %s", d.Id(), err)
	}
	return err
}
