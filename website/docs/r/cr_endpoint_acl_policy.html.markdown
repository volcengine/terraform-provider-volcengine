---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_endpoint_acl_policy"
sidebar_current: "docs-volcengine-resource-cr_endpoint_acl_policy"
description: |-
  Provides a resource to manage cr endpoint acl policy
---
# volcengine_cr_endpoint_acl_policy
Provides a resource to manage cr endpoint acl policy
## Example Usage
```hcl
resource "volcengine_cr_registry" "foo" {
  name    = "acc-test-cr-registry"
  project = "default"
}

resource "volcengine_cr_endpoint" "foo" {
  registry = volcengine_cr_registry.foo.id
  enabled  = true
}

resource "volcengine_cr_endpoint_acl_policy" "foo" {
  registry    = volcengine_cr_endpoint.foo.registry
  type        = "Public"
  entry       = "192.168.0.${count.index}"
  description = "test-${count.index}"
  count       = 3
}
```
## Argument Reference
The following arguments are supported:
* `description` - (Required, ForceNew) The description of the acl policy.
* `entry` - (Required, ForceNew) The ip list of the acl policy.
* `registry` - (Required, ForceNew) The registry name.
* `type` - (Required, ForceNew) The type of the acl policy. Valid values: `Public`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
CrEndpointAclPolicy can be imported using the registry:entry, e.g.
```
$ terraform import volcengine_cr_endpoint_acl_policy.default resource_id
```

