package iam_oidc_provider_client

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Iam OidcProvider key don't support import

*/

func ResourceVolcengineIamOidcProviderClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineIamOidcProviderClientCreate,
		Read:   resourceVolcengineIamOidcProviderClientRead,
		Delete: resourceVolcengineIamOidcProviderClientDelete,
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
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The client id of the OIDC provider.",
			},
		},
	}
}

func resourceVolcengineIamOidcProviderClientCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewIamOidcProviderClientService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Create(service, d, ResourceVolcengineIamOidcProviderClient())
}

func resourceVolcengineIamOidcProviderClientRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamOidcProviderClientService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Read(service, d, ResourceVolcengineIamOidcProviderClient())
}

func resourceVolcengineIamOidcProviderClientDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewIamOidcProviderClientService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineIamOidcProviderClient())
}
