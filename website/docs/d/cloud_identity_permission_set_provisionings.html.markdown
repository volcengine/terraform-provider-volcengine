---
subcategory: "CLOUD_IDENTITY"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_identity_permission_set_provisionings"
sidebar_current: "docs-volcengine-datasource-cloud_identity_permission_set_provisionings"
description: |-
  Use this data source to query detailed information of cloud identity permission set provisionings
---
# volcengine_cloud_identity_permission_set_provisionings
Use this data source to query detailed information of cloud identity permission set provisionings
## Example Usage
```hcl
data "volcengine_cloud_identity_permission_set_provisionings" "foo" {
  target_id = "210026****"
}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `permission_set_id` - (Optional) The id of cloud identity permission set.
* `target_id` - (Optional) The target account id of cloud identity permission set.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `permission_provisionings` - The collection of query.
    * `create_time` - The create time of the cloud identity permission set provisioning.
    * `id` - The id of the cloud identity permission set.
    * `permission_set_id` - The id of the cloud identity permission set.
    * `permission_set_name` - The name of the cloud identity permission set.
    * `target_id` - The target account id of the cloud identity permission set provisioning.
    * `update_time` - The update time of the cloud identity permission set provisioning.
* `total_count` - The total count of query.


