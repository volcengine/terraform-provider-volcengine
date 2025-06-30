---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_keyrings"
sidebar_current: "docs-volcengine-datasource-kms_keyrings"
description: |-
  Use this data source to query detailed information of kms keyrings
---
# volcengine_kms_keyrings
Use this data source to query detailed information of kms keyrings
## Example Usage
```hcl
data "volcengine_kms_keyrings" "default" {
  keyring_name = ["tf-test-1", "tf-test-2", "tf-test-3"]
  description  = ["tf-1", "tf-2"]
}
```
## Argument Reference
The following arguments are supported:
* `creation_date_range` - (Optional) The creation date of the keyring.
* `description` - (Optional) The description of the keyring.
* `keyring_name` - (Optional) The name of the keyring.
* `keyring_type` - (Optional) The type of the keyring.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The name of the project.
* `update_date_range` - (Optional) The update date of the keyring.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `keyrings` - The information about the keyring.
    * `creation_date` - The date when the keyring was created.
    * `description` - The description of the keyring.
    * `id` - The unique ID of the keyring. The value is in the UUID format.
    * `key_count` - Key ring key count.
    * `keyring_name` - The name of the keyring.
    * `keyring_type` - The type of the keyring.
    * `trn` - The information about the tenant resource name (TRN).
    * `uid` - The tenant ID of the keyring.
    * `update_date` - The date when the keyring was updated.
* `total_count` - The total count of query.


