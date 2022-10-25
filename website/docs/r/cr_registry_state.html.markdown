---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_registry_state"
sidebar_current: "docs-volcengine-resource-cr_registry_state"
description: |-
  Provides a resource to manage cr registry state
---
# volcengine_cr_registry_state
Provides a resource to manage cr registry state
## Example Usage
```hcl
resource "volcengine_cr_registry_state" "foo" {
  name   = "tf-2"
  action = "Start"
}
```
## Argument Reference
The following arguments are supported:
* `action` - (Required) Start cr instance action.
* `name` - (Required) The cr instance id.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `status` - The status of cr instance.
    * `conditions` - The condition of instance.
    * `phase` - The phase status of instance.


## Import
CR registry state can be imported using the state:registry_name, e.g.
```
$ terraform import volcengine_cr_registry.default state:cr-basic
```

