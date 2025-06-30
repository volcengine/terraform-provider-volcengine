---
subcategory: "EBS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ebs_max_extra_performances"
sidebar_current: "docs-volcengine-datasource-ebs_max_extra_performances"
description: |-
  Use this data source to query detailed information of ebs max extra performances
---
# volcengine_ebs_max_extra_performances
Use this data source to query detailed information of ebs max extra performances
## Example Usage
```hcl
data "volcengine_ebs_max_extra_performances" "foo" {
  volume_type = "TSSD_TL0"
  size        = 60
}
```
## Argument Reference
The following arguments are supported:
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `size` - (Optional) The size of the volume. Unit: GiB.
* `volume_id` - (Optional) The id of the volume.
* `volume_type` - (Optional) The type of the volume. Valid values: `TSSD_TL0`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `performances` - The collection of query.
    * `baseline` - The baseline of the performance.
        * `iops` - The baseline of the iops.
        * `throughput` - The baseline of the throughput.
    * `limit` - The limit of the performance.
        * `iops` - The limit of the iops.
        * `throughput` - The limit of the throughput.
    * `max_extra_performance_can_purchase` - The max extra performance can purchase.
        * `extra_performance_type_id` - The type of the extra performance.
        * `limit` - The limit of the extra performance.
* `total_count` - The total count of query.


