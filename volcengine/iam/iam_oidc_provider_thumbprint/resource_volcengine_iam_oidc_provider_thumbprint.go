package iam_oidc_provider_thumbprint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Iam OidcProviderThumbprint key don't support import

*/

func ResourceVolcengineIamOidcProviderThumbprint() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineIamOidcProviderThumbprintCreate,
		Read:   resourceVolcengineIamOidcProviderThumbprintRead,
		Delete: resourceVolcengineIamOidcProviderThumbprintDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"oidc_provider_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the OIDC provider.",
			},
			"thumbprint": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The thumbprint of the OIDC provider.",
			},
		},
	}
}

func resourceVolcengineIamOidcProviderThumbprintCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewIamOidcProviderThumbprintService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Create(service, d, ResourceVolcengineIamOidcProviderThumbprint())
}

func resourceVolcengineIamOidcProviderThumbprintRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamOidcProviderThumbprintService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Read(service, d, ResourceVolcengineIamOidcProviderThumbprint())
}

func resourceVolcengineIamOidcProviderThumbprintDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewIamOidcProviderThumbprintService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineIamOidcProviderThumbprint())
}
