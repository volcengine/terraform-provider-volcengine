---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_roles"
sidebar_current: "docs-volcengine-datasource-iam_roles"
description: |-
  Use this data source to query detailed information of iam roles
---
# volcengine_iam_roles
Use this data source to query detailed information of iam roles
## Example Usage
```hcl
resource "volcengine_iam_role" "foo1" {
  role_name             = "acc-test-role1"
  display_name          = "acc-test1"
  description           = "acc-test1"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}"
  max_session_duration  = 3600
}

resource "volcengine_iam_role" "foo2" {
  role_name             = "acc-test-role2"
  display_name          = "acc-test2"
  description           = "acc-test2"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"ecs\"]}}]}"
  max_session_duration  = 3600
}

data "volcengine_iam_roles" "foo" {
  role_name = "${volcengine_iam_role.foo1.role_name},${volcengine_iam_role.foo2.role_name}"
}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of Role.
* `output_file` - (Optional) File name where to save data source results.
* `query` - (Optional) The query field of Role.
* `role_name` - (Optional) The name of the Role, comma separated.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `roles` - The collection of Role query.
    * `create_date` - The create time of the Role.
    * `description` - The description of the Role.
    * `id` - The ID of the Role.
    * `role_name` - The name of the Role.
    * `trn` - The resource name of the Role.
    * `trust_policy_document` - The trust policy document of the Role.
* `total_count` - The total count of Role query.


