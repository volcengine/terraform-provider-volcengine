package certificate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Certificate can be imported using the id, e.g.
```
$ terraform import volcengine_certificate.default cert-2fe5k****c16o5oxruvtk3qf5
```

*/

func ResourceVolcengineCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineCertificateCreate,
		Read:   resourceVolcengineCertificateRead,
		Delete: resourceVolcengineCertificateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"certificate_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The name of the Certificate.",
			},
			"public_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The public key of the Certificate.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时支持，修改不支持
					return d.Id() != ""
				},
			},
			"private_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The private key of the Certificate.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时支持，修改不支持
					return d.Id() != ""
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The description of the Certificate.",
			},
		},
	}
}

func resourceVolcengineCertificateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	certificateService := NewCertificateService(meta.(*ve.SdkClient))
	err = certificateService.Dispatcher.Create(certificateService, d, ResourceVolcengineCertificate())
	if err != nil {
		return fmt.Errorf("error on creating certificate  %q, %w", d.Id(), err)
	}
	return resourceVolcengineCertificateRead(d, meta)
}

func resourceVolcengineCertificateRead(d *schema.ResourceData, meta interface{}) (err error) {
	certificateService := NewCertificateService(meta.(*ve.SdkClient))
	err = certificateService.Dispatcher.Read(certificateService, d, ResourceVolcengineCertificate())
	if err != nil {
		return fmt.Errorf("error on reading certificate %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineCertificateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	certificateService := NewCertificateService(meta.(*ve.SdkClient))
	err = certificateService.Dispatcher.Update(certificateService, d, ResourceVolcengineCertificate())
	if err != nil {
		return fmt.Errorf("error on updating certificate  %q, %w", d.Id(), err)
	}
	return resourceVolcengineCertificateRead(d, meta)
}

func resourceVolcengineCertificateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	certificateService := NewCertificateService(meta.(*ve.SdkClient))
	err = certificateService.Dispatcher.Delete(certificateService, d, ResourceVolcengineCertificate())
	if err != nil {
		return fmt.Errorf("error on deleting certificate %q, %w", d.Id(), err)
	}
	return err
}
