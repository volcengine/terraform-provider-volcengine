package cdn_certificate

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CdnCertificate can be imported using the id, e.g.
```
$ terraform import volcengine_cdn_certificate.default resource_id
```
You can delete the certificate hosted on the content delivery network.
You can configure the HTTPS module to associate the certificate and domain name through the domain_config field of volcengine_cdn_domain.
If the certificate to be deleted is already associated with a domain name, the deletion will fail.
To remove the association between the domain name and the certificate, you can disable the HTTPS function for the domain name in the Content Delivery Network console.
*/

func ResourceVolcengineCdnCertificate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCdnCertificateCreate,
		Read:   resourceVolcengineCdnCertificateRead,
		Delete: resourceVolcengineCdnCertificateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"certificate": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
				Description: "Represents a certificate object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: "Content of the specified certificate public key file. " +
								"Line breaks in the content should be replaced with `\\r\\n`. " +
								"The file extension for the certificate public key is `.crt` or `.pem`. " +
								"The public key must include the complete certificate chain. " +
								"When importing resources, this attribute will not be imported. " +
								"If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
						},
						"private_key": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: "The content of the specified certificate private key file. " +
								"Replace line breaks in the content with `\\r\\n`. " +
								"The file extension for the certificate private key is `.key` or `.pem`. " +
								"The private key must be unencrypted. " +
								"When importing resources, this attribute will not be imported. " +
								"If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
						},
					},
				},
			},
			"source": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "Specify the location for storing the certificate. " +
					"The parameter can take the following values: " +
					"`volc_cert_center`: indicates that the certificate will be stored in the certificate center." +
					"`cdn_cert_hosting`: indicates that the certificate will be hosted on the content delivery network.",
			},
			"cert_info": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Object representing a certificate remark.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"desc": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Note on the certificate.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineCdnCertificateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnCertificateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCdnCertificate())
	if err != nil {
		return fmt.Errorf("error on creating cdn_certificate %q, %s", d.Id(), err)
	}
	return resourceVolcengineCdnCertificateRead(d, meta)
}

func resourceVolcengineCdnCertificateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnCertificateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCdnCertificate())
	if err != nil {
		return fmt.Errorf("error on reading cdn_certificate %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCdnCertificateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnCertificateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCdnCertificate())
	if err != nil {
		return fmt.Errorf("error on deleting cdn_certificate %q, %s", d.Id(), err)
	}
	return err
}
