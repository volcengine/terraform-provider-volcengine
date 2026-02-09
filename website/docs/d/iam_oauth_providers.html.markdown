---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_oauth_providers"
sidebar_current: "docs-volcengine-datasource-iam_oauth_providers"
description: |-
  Use this data source to query detailed information of iam oauth providers
---
# volcengine_iam_oauth_providers
Use this data source to query detailed information of iam oauth providers
## Example Usage
```hcl
data "volcengine_iam_oauth_providers" "default" {
  oauth_provider_name = "acc-test-oauth"
}
```
## Argument Reference
The following arguments are supported:
* `oauth_provider_name` - (Required) The name of the OAuth provider.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `providers` - The collection of OAuth providers.
    * `authorize_template` - The authorize template of the OAuth provider.
    * `authorize_url` - The authorize url of the OAuth provider.
    * `client_id` - The client id of the OAuth provider.
    * `client_secret` - The client secret of the OAuth provider.
    * `create_date` - The create date of the OAuth provider.
    * `description` - The description of the OAuth provider.
    * `identity_map_type` - The identity map type of the OAuth provider.
    * `idp_identity_key` - The idp identity key of the OAuth provider.
    * `oauth_provider_name` - The name of the OAuth provider.
    * `provider_id` - The id of the OAuth provider.
    * `scope` - The scope of the OAuth provider.
    * `sso_type` - The SSO type of the OAuth provider.
    * `status` - The status of the OAuth provider.
    * `token_url` - The token url of the OAuth provider.
    * `trn` - The trn of the OAuth provider.
    * `update_date` - The update date of the OAuth provider.
    * `user_info_url` - The user info url of the OAuth provider.
* `total_count` - The total count of query.


