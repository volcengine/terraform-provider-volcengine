---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_tag"
sidebar_current: "docs-volcengine-resource-iam_tag"
description: |-
  Provides a resource to manage iam tag
---
# volcengine_iam_tag
Provides a resource to manage iam tag
## Example Usage
```hcl
resource "volcengine_iam_tag" "foo" {
  resource_type  = "User"
  resource_names = ["jonny"]
  tags {
    key   = "key4"
    value = "value4"
  }
  tags {
    key   = "key3"
    value = "value3"
  }
}
```
## Argument Reference
The following arguments are supported:
* `resource_names` - (Required, ForceNew) The names of the resource.
* `resource_type` - (Required, ForceNew) The type of the resource. Valid values: User, Role.
* `tags` - (Optional, ForceNew) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Iam tag can be imported using the ResourceType, ResourceName and TagKey, e.g.
```
$ terraform import volcengine_iam_tag.default User:jonny:key1
```

