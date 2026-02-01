---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_oidc_providers"
sidebar_current: "docs-volcengine-datasource-iam_oidc_providers"
description: |-
  Use this data source to query detailed information of iam oidc providers
---
# volcengine_iam_oidc_providers
Use this data source to query detailed information of iam oidc providers
## Example Usage
```hcl
data "volcengine_iam_oidc_providers" "foo" {

}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `oidc_providers` - The collection of query.
    * `client_ids` - The client IDs of the OIDC provider.
    * `create_date` - The create date of the OIDC provider.
    * `description` - The description of the OIDC provider.
    * `issuance_limit_time` - The issuance limit time of the OIDC provider.
    * `issuer_url` - The URL of the OIDC provider.
    * `provider_name` - The name of the OIDC provider.
    * `thumbprints` - The thumbprints of the OIDC provider.
    * `trn` - The trn of OIDC provider.
    * `update_date` - The update date of the OIDC provider.
* `total_count` - The total count of query.


