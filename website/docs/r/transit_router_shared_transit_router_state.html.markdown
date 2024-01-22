---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_shared_transit_router_state"
sidebar_current: "docs-volcengine-resource-transit_router_shared_transit_router_state"
description: |-
  Provides a resource to manage transit router shared transit router state
---
# volcengine_transit_router_shared_transit_router_state
Provides a resource to manage transit router shared transit router state
## Example Usage
```hcl
resource "volcengine_transit_router" "foo" {
  transit_router_name = "acc-test-tr"
  description         = "acc-test"
}

resource "volcengine_transit_router_shared_transit_router_state" "foo" {
  transit_router_id = volcengine_transit_router.foo.id
  action            = "Reject"
}
```
## Argument Reference
The following arguments are supported:
* `action` - (Required) `Accept` or `Reject` the shared transit router.
* `transit_router_id` - (Required, ForceNew) The id of the transit router.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
SharedTransitRouterState can be imported using the id, e.g.
```
$ terraform import volcengine_transit_router_shared_transit_router_state.default state:transitRouterId
```

