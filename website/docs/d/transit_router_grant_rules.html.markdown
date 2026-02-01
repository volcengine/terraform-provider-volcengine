---
subcategory: "TRANSIT_ROUTER"
layout: "volcengine"
page_title: "Volcengine: volcengine_transit_router_grant_rules"
sidebar_current: "docs-volcengine-datasource-transit_router_grant_rules"
description: |-
  Use this data source to query detailed information of transit router grant rules
---
# volcengine_transit_router_grant_rules
Use this data source to query detailed information of transit router grant rules
## Example Usage
```hcl
data "volcengine_transit_router_grant_rules" "foo" {
  transit_router_id = "tr-2bzy39uy6u3282dx0efxiqyq0"
}
```
## Argument Reference
The following arguments are supported:
* `transit_router_id` - (Required) The id of the transit router.
* `grant_account_id` - (Optional) The id of the grant account.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rules` - The collection of query.
    * `creation_time` - The creation time of the rule.
    * `description` - The description of the rule.
    * `grant_account_id` - The id of the grant account.
    * `status` - The status of the rule.
    * `transit_router_id` - The id of the transaction router.
    * `update_time` - The update time of the rule.
* `total_count` - The total count of query.


