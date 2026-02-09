---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_oidc_provider_thumbprint"
sidebar_current: "docs-volcengine-resource-iam_oidc_provider_thumbprint"
description: |-
  Provides a resource to manage iam oidc provider thumbprint
---
# volcengine_iam_oidc_provider_thumbprint
Provides a resource to manage iam oidc provider thumbprint
## Example Usage
```hcl
resource "volcengine_iam_oidc_provider_thumbprint" "foo" {
  oidc_provider_name = "oidc_provider"
  thumbprint         = "9b1afaa2dfca349fe38c5ef3e72ee03cb0696d65ea2e11f597ea9aa55fcgg33a"
}
```
## Argument Reference
The following arguments are supported:
* `oidc_provider_name` - (Required, ForceNew) The name of the OIDC provider.
* `thumbprint` - (Required, ForceNew) The thumbprint of the OIDC provider.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Iam OidcProviderThumbprint key don't support import

