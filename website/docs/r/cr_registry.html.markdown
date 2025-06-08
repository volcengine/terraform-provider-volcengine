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
# create cr registry
resource "volcengine_cr_registry" "foo" {
  name               = "acc-test-cr"
  delete_immediately = false
  password           = "1qaz!QAZ"
  project            = "default"
}

# create cr namespace
resource "volcengine_cr_namespace" "foo" {
  registry = volcengine_cr_registry.foo.id
  name     = "acc-test-namespace"
  project  = "default"
}

# create cr repository
resource "volcengine_cr_repository" "foo" {
  registry     = volcengine_cr_registry.foo.id
  namespace    = volcengine_cr_namespace.foo.name
  name         = "acc-test-repository"
  description  = "A test repository created by terraform."
  access_level = "Public"
}
```
## Argument Reference
The following arguments are supported:
* `name` - (Required, ForceNew) The name of registry.
* `delete_immediately` - (Optional) Whether delete registry immediately. Only effected in delete action.
* `password` - (Optional) The password of registry user.
* `project` - (Optional) The ProjectName of the cr registry.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `charge_type` - The charge type of registry.
* `create_time` - The creation time of registry.
* `domains` - The domain of registry.
    * `domain` - The domain of registry.
    * `type` - The domain type of registry.
* `resource_tags` - Tags.
    * `key` - The Key of Tags.
    * `value` - The Value of Tags.
* `status` - The status of registry.
    * `conditions` - The condition of registry.
    * `phase` - The phase status of registry.
* `type` - The type of registry.
* `user_status` - The status of user.
* `username` - The username of cr instance.


## Import
CR Registry can be imported using the name, e.g.
```
$ terraform import volcengine_cr_registry.default enterprise-x
```

