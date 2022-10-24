---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_registry"
sidebar_current: "docs-volcengine-resource-cr_registry"
description: |-
  Provides a resource to manage cr registry
---
# volcengine_cr_registry
Provides a resource to manage cr registry
## Example Usage
```hcl
resource "volcengine_cr_registry" "foo" {
  name               = "tf-1"
  delete_immediately = false
  password           = "1qaz!QAZ"
}
```
## Argument Reference
The following arguments are supported:
* `name` - (Required) The name of registry.
* `delete_immediately` - (Optional) Whether delete registry immediately.
* `password` - (Optional) The password of registry user.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `charge_type` - The charge type of registry.
* `create_time` - The creation time of registry.
* `domains` - The domain of registry.
    * `domain` - The domain of registry.
    * `type` - The domain type of registry.
* `status` - The status of registry.
    * `conditions` - The condition of registry.
    * `phase` - The phase status of registry.
* `type` - The type of registry.
* `user_status` - The status of user.
* `username` - The username of cr instance.


## Import
CR Instance can be imported using the name, e.g.
```
$ terraform import volcengine_cr_instance.default enterprise-x
```

