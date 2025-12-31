package alb_replace_certificate

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The AlbReplaceCertificate is not support import.

*/

func ResourceVolcengineAlbReplaceCertificate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineAlbReplaceCertificateCreate,
		Read:   resourceVolcengineAlbReplaceCertificateRead,
		//Update: resourceVolcengineAlbReplaceCertificateUpdate,
		Delete: resourceVolcengineAlbReplaceCertificateDelete,
		// Importer: &schema.ResourceImporter{
		// 	State: schema.ImportStatePassthrough,
		// },
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"certificate_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the certificate. Valid values: 'server' for server certificates, 'ca' for CA certificates.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					value := val.(string)
					if value != "server" && value != "ca" {
						errs = append(errs, fmt.Errorf("%s must be 'server' or 'ca', got: %s", key, value))
					}
					return warns, errs
				},
			},
			"old_certificate_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the old certificate to be replaced.",
			},
			"update_mode": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The mode of certificate replacement. Valid values: 'new' for uploading new certificate, 'stock' for using existing certificate.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					value := val.(string)
					if value != "new" && value != "stock" {
						errs = append(errs, fmt.Errorf("%s must be 'new' or 'stock', got: %s", key, value))
					}
					return warns, errs
				},
			},
			"certificate_source": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"alb", "cert_center"}, false),
				Description:  "The source of the server certificate. Valid values: `alb`, `cert_center`. Required when update_mode is 'stock'.",
			},
			"certificate_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The ID of the new certificate or CA certificate. Required when certificate_source is 'alb' and update_mode is 'stock'.",
			},
			"cert_center_certificate_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The ID of the new certificate. Required when certificate_source is 'cert_center' and update_mode is 'stock'.",
			},
			"certificate_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The name of the certificate.",
			},
			"description": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "The description of the certificate.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The project name of the certificate.",
			},
			// Server certificate specific fields (when certificate_type is 'server')
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				//Sensitive:   true,
				Description: "The public key of the server certificate. Required when certificate_type is 'server' and update_mode is 'new'.",
			},
			"private_key": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "The private key of the server certificate. Required when certificate_type is 'server' and update_mode is 'new'.",
			},
			// CA certificate specific fields (when certificate_type is 'ca')
			"ca_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				//Sensitive:   true,
				Description: "The content of the CA certificate. Required when certificate_type is 'ca' and update_mode is 'new'.",
			},
		},
	}
	return resource
}

func resourceVolcengineAlbReplaceCertificateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbReplaceCertificateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineAlbReplaceCertificate())
	if err != nil {
		return fmt.Errorf("error on creating alb_replace_certificate %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbReplaceCertificateRead(d, meta)
}

func resourceVolcengineAlbReplaceCertificateRead(d *schema.ResourceData, meta interface{}) (err error) {
	return nil
}

func resourceVolcengineAlbReplaceCertificateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbReplaceCertificateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineAlbReplaceCertificate())
	if err != nil {
		return fmt.Errorf("error on updating alb_replace_certificate %q, %s", d.Id(), err)
	}
	return resourceVolcengineAlbReplaceCertificateRead(d, meta)
}

func resourceVolcengineAlbReplaceCertificateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewAlbReplaceCertificateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineAlbReplaceCertificate())
	if err != nil {
		return fmt.Errorf("error on deleting alb_replace_certificate %q, %s", d.Id(), err)
	}
	return err
}
