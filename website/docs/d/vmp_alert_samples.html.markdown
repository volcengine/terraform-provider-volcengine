---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_alert_samples"
sidebar_current: "docs-volcengine-datasource-vmp_alert_samples"
description: |-
  Use this data source to query detailed information of vmp alert samples
---
# volcengine_vmp_alert_samples
Use this data source to query detailed information of vmp alert samples
## Example Usage
```hcl
data "volcengine_vmp_alert_samples" "example" {
  alert_id     = "695257b0d00908b4e7511fe4"
  sample_since = 1766851200
  sample_until = 1767006860
  limit        = 100
}
```
## Argument Reference
The following arguments are supported:
* `alert_id` - (Required) Alert ID to filter samples.
* `limit` - (Optional) Limit of samples, default 100, max 500.
* `sample_since` - (Optional) Filter start timestamp (unix).
* `sample_until` - (Optional) Filter end timestamp (unix).

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `alert_samples` - Alert samples collection.
    * `alert_id` - Alert ID.
    * `level` - Alert sample level.
    * `phase` - Alert sample phase.
    * `timestamp` - Alert sample timestamp(unix).
    * `value` - Alert sample value.
* `total_count` - The total count of query.


