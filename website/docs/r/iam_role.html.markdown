---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_role"
sidebar_current: "docs-volcengine-resource-iam_role"
description: |-
  Provides a resource to manage iam role
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_iam_role
Provides a resource to manage iam role
## Example Usage
```hcl
resource "volcengine_iam_role" "foo" {
  role_name             = "tf-test-lh"
  display_name          = "tf-test-lh"
  description           = "tf-test-lh"
  trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"IAM\":[\"trn:iam::2000000001:root\"]}}]}"
  max_session_duration  = 4800
  tags {
    key   = "key-1"
    value = "value-1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `role_name` - (Required) The name of the Role.
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

