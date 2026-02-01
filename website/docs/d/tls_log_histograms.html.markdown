---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_log_histograms"
sidebar_current: "docs-volcengine-datasource-tls_log_histograms"
description: |-
  Use this data source to query detailed information of tls log histograms
---
# volcengine_tls_log_histograms
Use this data source to query detailed information of tls log histograms
## Example Usage
```hcl

```
## Argument Reference
The following arguments are supported:
* `end_time` - (Required) The end time.
* `query` - (Required) The query statement.
* `start_time` - (Required) The start time.
* `topic_id` - (Required) The topic id.
* `interval` - (Optional) The interval.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `histogram_infos` - The histogram info.
    * `count` - The count.
    * `end_time` - The end time.
    * `result_status` - The result status.
    * `start_time` - The start time.
* `result_status` - The result status.
* `total_count` - The total count.


