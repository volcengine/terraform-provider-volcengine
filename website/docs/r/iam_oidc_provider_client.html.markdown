---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_oidc_provider_client"
sidebar_current: "docs-volcengine-resource-iam_oidc_provider_client"
description: |-
  Provides a resource to manage iam oidc provider client
---
# volcengine_iam_oidc_provider_client
Provides a resource to manage iam oidc provider client
## Example Usage
```hcl
resource "volcengine_iam_oidc_provider_client" "foo" {
  oidc_provider_name = "oidc_provider"
  client_id          = "test_client_id_2"
}
```
## Argument Reference
The following arguments are supported:
* `client_id` - (Required, ForceNew) The client id of the OIDC provider.
* `oidc_provider_name` - (Required, ForceNew) The name of the OIDC provider.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Iam OidcProvider key don't support import

