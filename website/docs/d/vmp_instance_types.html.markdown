---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_instance_types"
sidebar_current: "docs-volcengine-datasource-vmp_instance_types"
description: |-
  Use this data source to query detailed information of vmp instance types
---
# volcengine_vmp_instance_types
Use this data source to query detailed information of vmp instance types
## Example Usage
```hcl
data "volcengine_vmp_instance_types" "default" {
  ids = ["vmp.standard.15d"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Instance Type IDs.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instance_types` - The collection of query.
    * `active_series` - Maximum number of active sequences.
    * `availability_zone_replicas` - Number of zone.
    * `dedicated` - Whether the workspace is exclusive.
    * `id` - The ID of instance type.
    * `ingest_samples_per_second` - Maximum write samples per second.
    * `query_concurrency` - Maximum number of concurrent queries.
    * `query_per_second` - Maximum Query QPS.
    * `replicas_per_zone` - Data replicas per az.
    * `retention_period` - Maximum data retention time.
    * `scan_samples_per_second` - Maximum scan samples per second.
    * `scan_series_per_second` - Maximum number of scan sequences per second.
* `total_count` - The total count of query.


