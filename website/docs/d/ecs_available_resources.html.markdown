---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_available_resources"
sidebar_current: "docs-volcengine-datasource-ecs_available_resources"
description: |-
  Use this data source to query detailed information of ecs available resources
---
# volcengine_ecs_available_resources
Use this data source to query detailed information of ecs available resources
## Example Usage
```hcl
data "volcengine_ecs_available_resources" "foo" {
  destination_resource = "InstanceType"
}
```
## Argument Reference
The following arguments are supported:
* `destination_resource` - (Required) The type of resource to query. Valid values: `InstanceType`, `DedicatedHost`.
* `instance_charge_type` - (Optional) The charge type of instance. Valid values: `PostPaid`, `PrePaid`, `ReservedInstance`. Default is `PostPaid`.
* `instance_type_id` - (Optional) The id of instance type.
* `output_file` - (Optional) File name where to save data source results.
* `spot_strategy` - (Optional) The spot strategy of PostPaid instance. Valid values: `NoSpot`, `SpotAsPriceGo`. Default is `NoSpot`.
* `zone_id` - (Optional) The id of available zone.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `available_zones` - The collection of query.
    * `available_resources` - The resource information of the available zone.
        * `supported_resources` - The supported resource information.
            * `status` - The status of the resource. Valid values: `Available`, `SoldOut`.
            * `value` - The value of the resource.
        * `type` - The type of resource. Valid values: `InstanceType`, `DedicatedHost`.
    * `region_id` - The id of the region.
    * `status` - The resource status of the available zone. Valid values: `Available`, `SoldOut`.
    * `zone_id` - The id of the available zone.
* `total_count` - The total count of query.


