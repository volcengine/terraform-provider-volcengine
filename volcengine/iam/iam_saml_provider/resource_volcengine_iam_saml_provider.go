package iam_saml_provider

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
IamSamlProvider can be imported using the id, e.g.
```
$ terraform import volcengine_iam_saml_provider.default SAMLProviderName
```

*/

func ResourceVolcengineIamSamlProvider() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineIamSamlProviderCreate,
		Read:   resourceVolcengineIamSamlProviderRead,
		Update: resourceVolcengineIamSamlProviderUpdate,
		Delete: resourceVolcengineIamSamlProviderDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"saml_provider_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the SAML provider.",
			},
			"encoded_saml_metadata_document": {
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.Replace(old, "\n", "", -1) == strings.Replace(new, "\n", "", -1)
				},
				Description: "Metadata document, encoded in Base64.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the SAML provider.",
			},
			"sso_type": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "SSO types, 1. Role-based SSO, 2. User-based SSO.",
			},
			"status": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "User SSO status, 1. Enabled, 2. Disable other console login methods after enabling, " +
					"3. Disabled, is a required field when creating user SSO.",
			},
			"trn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The format for the resource name of an identity provider is trn:iam::${accountID}:saml-provider/{$SAMLProviderName}.",
			},
			"create_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identity provider creation time, such as 20150123T123318Z.",
			},
			"update_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identity provider update time, such as: 20150123T123318Z.",
			},
		},
	}
	return resource
}

func resourceVolcengineIamSamlProviderCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamSamlProviderService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineIamSamlProvider())
	if err != nil {
		return fmt.Errorf("error on creating iam_saml_provider %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamSamlProviderRead(d, meta)
}

func resourceVolcengineIamSamlProviderRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamSamlProviderService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineIamSamlProvider())
	if err != nil {
		return fmt.Errorf("error on reading iam_saml_provider %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineIamSamlProviderUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamSamlProviderService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineIamSamlProvider())
	if err != nil {
		return fmt.Errorf("error on updating iam_saml_provider %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamSamlProviderRead(d, meta)
}

func resourceVolcengineIamSamlProviderDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamSamlProviderService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineIamSamlProvider())
	if err != nil {
		return fmt.Errorf("error on deleting iam_saml_provider %q, %s", d.Id(), err)
	}
	return err
}
