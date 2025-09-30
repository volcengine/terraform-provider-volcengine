package iam_oidc_provider

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
IamOidcProvider can be imported using the id, e.g.
```
$ terraform import volcengine_iam_oidc_provider.default resource_id
```

*/

func ResourceVolcengineIamOidcProvider() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineIamOidcProviderCreate,
		Read:   resourceVolcengineIamOidcProviderRead,
		Update: resourceVolcengineIamOidcProviderUpdate,
		Delete: resourceVolcengineIamOidcProviderDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"oidc_provider_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the OIDC provider.",
			},
			"issuer_url": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The URL of the OIDC provider.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the OIDC provider.",
			},
			"issuance_limit_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The issuance limit time of the OIDC provider.",
			},
			"client_ids": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The client IDs of the OIDC provider.",
			},
			"thumbprints": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The thumbprints of the OIDC provider.",
			},

			// computed fields
			"trn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The trn of OIDC provider.",
			},
			"create_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create date of the OIDC provider.",
			},
			"update_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update date of the OIDC provider.",
			},
		},
	}
	return resource
}

func resourceVolcengineIamOidcProviderCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamOidcProviderService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineIamOidcProvider())
	if err != nil {
		return fmt.Errorf("error on creating iam_oidc_provider %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamOidcProviderRead(d, meta)
}

func resourceVolcengineIamOidcProviderRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamOidcProviderService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineIamOidcProvider())
	if err != nil {
		return fmt.Errorf("error on reading iam_oidc_provider %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineIamOidcProviderUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamOidcProviderService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineIamOidcProvider())
	if err != nil {
		return fmt.Errorf("error on updating iam_oidc_provider %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamOidcProviderRead(d, meta)
}

func resourceVolcengineIamOidcProviderDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamOidcProviderService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineIamOidcProvider())
	if err != nil {
		return fmt.Errorf("error on deleting iam_oidc_provider %q, %s", d.Id(), err)
	}
	return err
}
