---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_grant_rule"
sidebar_current: "docs-volcengine-resource-transit_router_grant_rule"
description: |-
  Provides a resource to manage transit router grant rule
---
# volcengine_transit_router_grant_rule
Provides a resource to manage transit router grant rule
## Example Usage
```hcl
resource "volcengine_transit_router" "foo" {
  transit_router_name = "acc-test-tf"
  description         = "acc-test-tf"
}

resource "volcengine_transit_router_grant_rule" "foo" {
  grant_account_id  = "2000xxxxx"
  description       = "acc-test-tf"
  transit_router_id = volcengine_transit_router.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `grant_account_id` - (Required, ForceNew) Account ID awaiting authorization for intermediate router instance.
* `transit_router_id` - (Required, ForceNew) The id of the transit router.
* `description` - (Optional) The description of the rule.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TransitRouterGrantRule can be imported using the transit router id and accountId, e.g.
```
$ terraform import volcengine_transit_router_grant_rule.default trId:accountId
```

