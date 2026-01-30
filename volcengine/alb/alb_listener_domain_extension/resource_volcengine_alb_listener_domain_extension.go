package alb_listener_domain_extension

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
AlbListenerDomainExtension can be imported using the listener id and domain extension id, e.g.
```
$ terraform import volcengine_alb_listener_domain_extension.default listenerId:extensionId
```

*/

func ResourceVolcengineAlbListenerDomainExtension() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineAlbListenerDomainExtensionCreate,
		Read:   resourceVolcengineAlbListenerDomainExtensionRead,
		Update: resourceVolcengineAlbListenerDomainExtensionUpdate,
		Delete: resourceVolcengineAlbListenerDomainExtensionDelete,
		Importer: &schema.ResourceImporter{
			State: importAlbListenerDomainExtension,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"listener_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The listener id. Only HTTPS listener is effective.",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The domain name. The maximum number of associated domain names for an HTTPS listener is 20, with a value range of 1 to 20.",
			},
			"certificate_source": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "alb",
				Description: "The source of the certificate. Valid values: `alb`, `cert_center`, `pca_leaf`. Default is `alb`.",
			},
			"certificate_id": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"certificate_id", "cert_center_certificate_id", "pca_leaf_certificate_id"},
				Description:  "Server certificate used for the domain name. Valid when the certificate_source is `alb`.",
			},
			"cert_center_certificate_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The server certificate ID used by the domain name. Valid when the certificate_source is `cert_center`.",
			},
			"pca_leaf_certificate_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The server certificate ID used by the domain name. Valid when the certificate source is `pca_leaf`.",
			},
			"domain_extension_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the domain extension.",
			},
		},
	}
	return resource
}

func resourceVolcengineAlbListenerDomainExtensionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbListenerDomainExtensionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineAlbListenerDomainExtension())
	if err != nil {
		return fmt.Errorf("error on creating alb_listener_domain_extension %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbListenerDomainExtensionRead(d, meta)
}

func resourceVolcengineAlbListenerDomainExtensionRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbListenerDomainExtensionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineAlbListenerDomainExtension())
	if err != nil {
		return fmt.Errorf("error on reading alb_listener_domain_extension %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineAlbListenerDomainExtensionUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbListenerDomainExtensionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineAlbListenerDomainExtension())
	if err != nil {
		return fmt.Errorf("error on updating alb_listener_domain_extension %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbListenerDomainExtensionRead(d, meta)
}

func resourceVolcengineAlbListenerDomainExtensionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbListenerDomainExtensionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineAlbListenerDomainExtension())
	if err != nil {
		return fmt.Errorf("error on deleting alb_listener_domain_extension %q, %s", d.Id(), err)
	}
	return err
}

func importAlbListenerDomainExtension(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form listenerId:extensionId")
	}
	err = data.Set("listener_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("domain_extension_id", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
