---
subcategory: "CLOUD_IDENTITY"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_identity_user_provisionings"
sidebar_current: "docs-volcengine-datasource-cloud_identity_user_provisionings"
description: |-
  Use this data source to query detailed information of cloud identity user provisionings
---
# volcengine_cloud_identity_user_provisionings
Use this data source to query detailed information of cloud identity user provisionings
## Example Usage
```hcl
data "volcengine_cloud_identity_user_provisionings" "foo" {
  account_id = "210026****"
}
```
## Argument Reference
The following arguments are supported:
* `account_id` - (Optional) The account id.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `user_provisionings` - The collection of query.
    * `created_time` - The created time of the cloud identity user provisioning.
    * `deletion_strategy` - The deletion strategy of the cloud identity user provisioning.
    * `department_names` - The department names of the cloud identity user provisioning.
    * `description` - The description of the cloud identity user provisioning.
    * `duplication_strategy` - The duplication strategy of the cloud identity user provisioning.
    * `duplication_suffix` - The duplication suffix of the cloud identity user provisioning.
    * `id` - The id of the cloud identity user provisioning.
    * `identity_source_strategy` - The identity source strategy of the cloud identity user provisioning.
    * `principal_id` - The principal id of the cloud identity user provisioning.
    * `principal_name` - The principal name of the cloud identity user provisioning.
    * `principal_type` - The principal type of the cloud identity user provisioning.
    * `provision_status` - The status of the cloud identity user provisioning.
    * `target_id` - The target account id of the cloud identity user provisioning.
    * `updated_time` - The updated time of the cloud identity user provisioning.
    * `user_provisioning_id` - The id of the cloud identity user provisioning.


