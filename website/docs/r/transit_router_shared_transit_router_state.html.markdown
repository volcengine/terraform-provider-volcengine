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
provider "volcengine" {
  access_key = "access_key_1"
  secret_key = "secret_key_1"
  region     = "region_1"
}

provider "volcengine" {
  access_key = "access_key_2"
  secret_key = "secret_key_2"
  region     = "region_2"
  alias      = "second_account"
}

resource "volcengine_transit_router" "foo" {
  transit_router_name = "acc-test-tr"
  description         = "acc-test"
}

resource "volcengine_transit_router_grant_rule" "foo" {
  grant_account_id  = "2000xxxxx"
  description       = "acc-test-tf"
  transit_router_id = volcengine_transit_router.foo.id
}

resource "volcengine_transit_router_shared_transit_router_state" "foo" {
  transit_router_id = volcengine_transit_router.foo.id
  action            = "Accept"
  provider          = volcengine.second_account
}
```
## Argument Reference
The following arguments are supported:
* `action` - (Required) `Accept` or `Reject` the shared transit router. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `transit_router_id` - (Required, ForceNew) The id of the transit router.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
SharedTransitRouterState can be imported using the id, e.g.
```
$ terraform import volcengine_transit_router_shared_transit_router_state.default state:transitRouterId
```

