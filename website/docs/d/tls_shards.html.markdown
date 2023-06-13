---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_shards"
sidebar_current: "docs-volcengine-datasource-tls_shards"
description: |-
  Use this data source to query detailed information of tls shards
---
# volcengine_tls_shards
Use this data source to query detailed information of tls shards
## Example Usage
```hcl
data "volcengine_tls_shards" "default" {
  topic_id = "edf051ed-3c46-49ba-9339-bea628fedc15"
}
```
## Argument Reference
The following arguments are supported:
* `topic_id` - (Required) The id of topic.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `shards` - The collection of query.
    * `exclusive_end_key` - The end key info.
    * `inclusive_begin_key` - The begin key info.
    * `modify_time` - The modify time.
    * `shard_id` - The id of shard.
    * `status` - The status of shard.
    * `stop_write_time` - The stop write time.
    * `topic_id` - The ID of topic.
* `total_count` - The total count of query.


