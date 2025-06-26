---
subcategory: "RABBITMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rabbitmq_regions"
sidebar_current: "docs-volcengine-datasource-rabbitmq_regions"
description: |-
  Use this data source to query detailed information of rabbitmq regions
---
# volcengine_rabbitmq_regions
Use this data source to query detailed information of rabbitmq regions
## Example Usage
```hcl
data "volcengine_rabbitmq_regions" "foo" {

}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `regions` - The collection of query.
    * `description` - The description of region.
    * `region_id` - The ID of region.
    * `region_name` - The name of region.
    * `status` - The status of region.
* `total_count` - The total count of query.


