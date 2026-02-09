---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_role"
sidebar_current: "docs-volcengine-resource-iam_role"
description: |-
  Provides a resource to manage iam role
---
# volcengine_iam_role
Provides a resource to manage iam role
## Example Usage
```hcl
resource "volcengine_iam_role" "foo" {
  role_name             = "tf-test-role"
  display_name          = "tf-test-modify"
  description           = "tf-test-modify"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}"
  max_session_duration  = 3600
  tags {
    key   = "key-modify"
    value = "value-modify"
  }
}
```
## Argument Reference
The following arguments are supported:
* `role_name` - (Required, ForceNew) The name of the Role.
* `description` - (Optional) The description of the Role.
* `display_name` - (Optional) The display name of the Role.
* `max_session_duration` - (Optional) The max session duration of the Role.
* `tags` - (Optional) Tags.
* `trust_policy_document` - (Optional) The trust policy document of the Role.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `is_service_linked_role` - Whether the Role is a service linked role.
* `role_id` - The id of the Role.
* `trn` - The resource name of the Role.


## Import
Iam role can be imported using the id, e.g.
```
$ terraform import volcengine_iam_role.default TerraformTestRole
```

