---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_oidc_provider"
sidebar_current: "docs-volcengine-resource-iam_oidc_provider"
description: |-
  Provides a resource to manage iam oidc provider
---
# volcengine_iam_oidc_provider
Provides a resource to manage iam oidc provider
## Example Usage
```hcl
resource "volcengine_iam_oidc_provider" "foo" {
  oidc_provider_name  = "oidc_provider"
  issuer_url          = "https://security-api.snssdk.com/qa/sso/oidc/6c505fb67d32417c8de287ee1fa89fc1"
  description         = "acc-test-oidc-modify"
  issuance_limit_time = 10
  client_ids          = ["6c505fb67d32417c8de287ee1fa89fd2"]
  thumbprints         = ["9b1afaa2dfca349fe38c5ef3e72ee03cb0696d65ea2e11f597ea9aa55fcff44d"]
}
```
## Argument Reference
The following arguments are supported:
* `client_ids` - (Required) The client IDs of the OIDC provider.
* `issuer_url` - (Required, ForceNew) The URL of the OIDC provider.
* `oidc_provider_name` - (Required, ForceNew) The name of the OIDC provider.
* `thumbprints` - (Required) The thumbprints of the OIDC provider.
* `description` - (Optional) The description of the OIDC provider.
* `issuance_limit_time` - (Optional) The issuance limit time of the OIDC provider.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_date` - The create date of the OIDC provider.
* `trn` - The trn of OIDC provider.
* `update_date` - The update date of the OIDC provider.


## Import
IamOidcProvider can be imported using the id, e.g.
```
$ terraform import volcengine_iam_oidc_provider.default resource_id
```

