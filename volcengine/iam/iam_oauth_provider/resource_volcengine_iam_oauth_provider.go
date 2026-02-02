package iam_oauth_provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
IamOAuthProvider can be imported using the id, e.g.
```
$ terraform import volcengine_iam_oauth_provider.default oidc_provider_name
```

*/

func ResourceVolcengineIamOAuthProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineIamOAuthProviderCreate,
		Read:   resourceVolcengineIamOAuthProviderRead,
		Update: resourceVolcengineIamOAuthProviderUpdate,
		Delete: resourceVolcengineIamOAuthProviderDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"oauth_provider_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the OAuth provider.",
			},
			"provider_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the OAuth provider.",
			},
			"sso_type": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The SSO type of the OAuth provider.",
			},
			"status": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The status of the OAuth provider.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the OAuth provider.",
			},
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The client id of the OAuth provider.",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The client secret of the OAuth provider.",
			},
			"user_info_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user info url of the OAuth provider.",
			},
			"token_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The token url of the OAuth provider.",
			},
			"authorize_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The authorize url of the OAuth provider.",
			},
			"authorize_template": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The authorize template of the OAuth provider.",
			},
			"scope": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The scope of the OAuth provider.",
			},
			"identity_map_type": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The identity map type of the OAuth provider.",
			},
			"idp_identity_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The idp identity key of the OAuth provider.",
			},
			"trn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The trn of the OAuth provider.",
			},
			"create_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create date of the OAuth provider.",
			},
			"update_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update date of the OAuth provider.",
			},
		},
	}
}

func resourceVolcengineIamOAuthProviderCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewIamOAuthProviderService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Create(service, d, ResourceVolcengineIamOAuthProvider())
}

func resourceVolcengineIamOAuthProviderRead(d *schema.ResourceData, meta interface{}) error {
	service := NewIamOAuthProviderService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Read(service, d, ResourceVolcengineIamOAuthProvider())
}

func resourceVolcengineIamOAuthProviderUpdate(d *schema.ResourceData, meta interface{}) error {
	service := NewIamOAuthProviderService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Update(service, d, ResourceVolcengineIamOAuthProvider())
}

func resourceVolcengineIamOAuthProviderDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewIamOAuthProviderService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineIamOAuthProvider())
}
