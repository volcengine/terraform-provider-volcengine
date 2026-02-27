---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_secret_versions"
sidebar_current: "docs-volcengine-datasource-kms_secret_versions"
description: |-
  Use this data source to query detailed information of kms secret versions
---
# volcengine_kms_secret_versions
Use this data source to query detailed information of kms secret versions
## Example Usage
```hcl
data "volcengine_kms_secret_versions" "default" {
  secret_name = "secret-9527"
}
```
## Argument Reference
The following arguments are supported:
* `secret_name` - (Required) The name of the secret.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `secret_versions` - The version info of secret.
    * `creation_date` - The creation time of secret version.
    * `version_id` - The version ID of secret value.
    * `version_stage` - The version stage of secret value.
* `total_count` - The total count of query.


