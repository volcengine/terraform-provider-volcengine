---
subcategory: "RABBITMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rabbitmq_zones"
sidebar_current: "docs-volcengine-datasource-rabbitmq_zones"
description: |-
  Use this data source to query detailed information of rabbitmq zones
---
# volcengine_rabbitmq_zones
Use this data source to query detailed information of rabbitmq zones
## Example Usage
```hcl
data "volcengine_rabbitmq_zones" "foo" {

}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `zones` - The collection of query.
    * `description` - The description of zone.
    * `status` - The status of zone.
    * `zone_id` - The ID of zone.
    * `zone_name` - The name of zone.


