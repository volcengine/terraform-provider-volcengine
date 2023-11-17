---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_endpoint"
sidebar_current: "docs-volcengine-resource-cr_endpoint"
description: |-
  Provides a resource to manage cr endpoint
---
# volcengine_cr_endpoint
Provides a resource to manage cr endpoint
## Example Usage
```hcl
resource "volcengine_cr_endpoint" "default" {
  registry = "acc-test-cr"
  enabled  = true
}
```
## Argument Reference
The following arguments are supported:
* `registry` - (Required, ForceNew) The CrRegistry name.
* `enabled` - (Optional) Whether enable public endpoint.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `status` - The status of public endpoint.


## Import
CR endpoints can be imported using the endpoint:registryName, e.g.
```
$ terraform import volcengine_cr_endpoint.default endpoint:cr-basic
```

