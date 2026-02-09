---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_oauth_provider"
sidebar_current: "docs-volcengine-resource-iam_oauth_provider"
description: |-
  Provides a resource to manage iam oauth provider
---
# volcengine_iam_oauth_provider
Provides a resource to manage iam oauth provider
## Example Usage
```hcl
resource "volcengine_iam_oauth_provider" "foo" {
  oauth_provider_name = "acc-test-oauth"
  sso_type            = 2
  status              = 1
  description         = "acc-test-modify"
  client_id           = "test_client_id_modify"
  client_secret       = ""
  user_info_url       = "https://example.com/user_info_modify"
  token_url           = "https://example.com/access_token_modify"
  authorize_url       = "https://example.com/authorize_modify"
  authorize_template  = "$${authEndpoint}?client_id=$${clientId}&scope=$${scope}&response_type=code&state=12345"
  scope               = "openid"
  identity_map_type   = 1
  idp_identity_key    = "username_modify"
}
```
## Argument Reference
The following arguments are supported:
* `authorize_template` - (Required) The authorize template of the OAuth provider.
* `authorize_url` - (Required) The authorize url of the OAuth provider.
* `client_id` - (Required) The client id of the OAuth provider.
* `client_secret` - (Required) The client secret of the OAuth provider.
* `identity_map_type` - (Required) The identity map type of the OAuth provider.
* `idp_identity_key` - (Required) The idp identity key of the OAuth provider.
* `oauth_provider_name` - (Required, ForceNew) The name of the OAuth provider.
* `sso_type` - (Required, ForceNew) The SSO type of the OAuth provider.
* `token_url` - (Required) The token url of the OAuth provider.
* `user_info_url` - (Required) The user info url of the OAuth provider.
* `description` - (Optional) The description of the OAuth provider.
* `scope` - (Optional) The scope of the OAuth provider.
* `status` - (Optional) The status of the OAuth provider.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_date` - The create date of the OAuth provider.
* `provider_id` - The id of the OAuth provider.
* `trn` - The trn of the OAuth provider.
* `update_date` - The update date of the OAuth provider.


## Import
IamOAuthProvider can be imported using the id, e.g.
```
$ terraform import volcengine_iam_oauth_provider.default oidc_provider_name
```

