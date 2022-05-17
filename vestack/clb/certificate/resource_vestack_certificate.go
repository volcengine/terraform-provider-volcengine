package certificate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Certificate can be imported using the id, e.g.
```
$ terraform import vestack_certificate.default cert-2fe5k****c16o5oxruvtk3qf5
```

*/

func ResourceVestackCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceVestackCertificateCreate,
		Read:   resourceVestackCertificateRead,
		Delete: resourceVestackCertificateDelete,
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
			},
			"private_key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The private key of the Certificate.",
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

func resourceVestackCertificateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	certificateService := NewCertificateService(meta.(*ve.SdkClient))
	err = certificateService.Dispatcher.Create(certificateService, d, ResourceVestackCertificate())
	if err != nil {
		return fmt.Errorf("error on creating certificate  %q, %w", d.Id(), err)
	}
	return resourceVestackCertificateRead(d, meta)
}

func resourceVestackCertificateRead(d *schema.ResourceData, meta interface{}) (err error) {
	certificateService := NewCertificateService(meta.(*ve.SdkClient))
	err = certificateService.Dispatcher.Read(certificateService, d, ResourceVestackCertificate())
	if err != nil {
		return fmt.Errorf("error on reading certificate %q, %w", d.Id(), err)
	}
	return err
}

func resourceVestackCertificateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	certificateService := NewCertificateService(meta.(*ve.SdkClient))
	err = certificateService.Dispatcher.Update(certificateService, d, ResourceVestackCertificate())
	if err != nil {
		return fmt.Errorf("error on updating certificate  %q, %w", d.Id(), err)
	}
	return resourceVestackCertificateRead(d, meta)
}

func resourceVestackCertificateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	certificateService := NewCertificateService(meta.(*ve.SdkClient))
	err = certificateService.Dispatcher.Delete(certificateService, d, ResourceVestackCertificate())
	if err != nil {
		return fmt.Errorf("error on deleting certificate %q, %w", d.Id(), err)
	}
	return err
}
