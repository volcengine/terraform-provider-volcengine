---
subcategory: "CLOUD_IDENTITY"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_identity_permission_sets"
sidebar_current: "docs-volcengine-datasource-cloud_identity_permission_sets"
description: |-
  Use this data source to query detailed information of cloud identity permission sets
---
# volcengine_cloud_identity_permission_sets
Use this data source to query detailed information of cloud identity permission sets
## Example Usage
```hcl
resource "volcengine_cloud_identity_permission_set" "foo" {
  name             = "acc-test-permission_set-${count.index}"
  description      = "tf"
  session_duration = 5000
  permission_policies {
    permission_policy_type = "System"
    permission_policy_name = "AdministratorAccess"
    inline_policy_document = ""
  }
  permission_policies {
    permission_policy_type = "System"
    permission_policy_name = "ReadOnlyAccess"
    inline_policy_document = ""
  }
  permission_policies {
    permission_policy_type = "Inline"
    inline_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
  }

  count = 2
}

data "volcengine_cloud_identity_permission_sets" "foo" {
  ids = volcengine_cloud_identity_permission_set.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of cloud identity permission set IDs.
* `name_regex` - (Optional) A Name Regex of cloud identity permission set.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `permission_sets` - The collection of query.
    * `created_time` - The create time of the cloud identity permission set.
    * `description` - The description of the cloud identity permission set.
    * `id` - The id of the cloud identity permission set.
    * `name` - The name of the cloud identity permission set.
    * `permission_policies` - The policies of the cloud identity permission set.
        * `create_time` - The create time of the cloud identity permission set policy.
        * `permission_policy_document` - The document of the cloud identity permission set policy.
        * `permission_policy_name` - The name of the cloud identity permission set policy.
        * `permission_policy_type` - The type of the cloud identity permission set policy.
    * `permission_set_id` - The id of the cloud identity permission set.
    * `relay_state` - The relay state of the cloud identity permission set.
    * `session_duration` - The session duration of the cloud identity permission set.
    * `updated_time` - The updated time of the cloud identity permission set.
* `total_count` - The total count of query.


