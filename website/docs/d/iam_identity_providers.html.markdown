---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_identity_providers"
sidebar_current: "docs-volcengine-datasource-iam_identity_providers"
description: |-
  Use this data source to query detailed information of iam identity providers
---
# volcengine_iam_identity_providers
Use this data source to query detailed information of iam identity providers
## Example Usage
```hcl
data "volcengine_iam_identity_providers" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `providers` - The collection of identity providers.
    * `create_date` - The create date of the identity provider.
    * `description` - The description of the identity provider.
    * `idp_type` - The type of the identity provider.
    * `provider_name` - The name of the identity provider.
    * `sso_type` - The SSO type of the identity provider.
    * `status` - The status of the identity provider.
    * `trn` - The TRN of the identity provider.
    * `update_date` - The update date of the identity provider.
* `total_count` - The total count of identity providers.


