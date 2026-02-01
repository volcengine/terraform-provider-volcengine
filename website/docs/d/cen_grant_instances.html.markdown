---
subcategory: "CEN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_grant_instances"
sidebar_current: "docs-volcengine-datasource-cen_grant_instances"
description: |-
  Use this data source to query detailed information of cen grant instances
---
# volcengine_cen_grant_instances
Use this data source to query detailed information of cen grant instances
## Example Usage
```hcl
data "volcengine_cen_grant_instances" "foo" {
  instance_type      = "VPC"
  instance_id        = "vpc-2bysvq1xx543k2dx0eeul****"
  instance_region_id = "cn-beijing"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Optional) The ID of the instance.
* `instance_region_id` - (Optional) The region ID of the instance.
* `instance_type` - (Optional) The type of the instance. Valid values: `VPC`, `DCGW`.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `grant_rules` - The collection of query.
    * `cen_id` - The ID of the cen.
    * `cen_owner_id` - The owner ID of the cen.
    * `creation_time` - The creation time of the grant rule.
    * `instance_id` - The ID of the instance.
    * `instance_region_id` - The region ID of the instance.
    * `instance_type` - The type of the instance.
* `total_count` - The total count of query.


