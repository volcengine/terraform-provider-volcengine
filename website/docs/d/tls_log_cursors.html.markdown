---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_log_cursors"
sidebar_current: "docs-volcengine-datasource-tls_log_cursors"
description: |-
  Use this data source to query detailed information of tls log cursors
---
# volcengine_tls_log_cursors
Use this data source to query detailed information of tls log cursors
## Example Usage
```hcl
data "volcengine_tls_log_cursors" "default" {
  topic_id = "e101b8c8-77e7-4ae3-91c1-2532ee480e7d"
  shard_id = 0
  from     = "begin"
}
```
## Argument Reference
The following arguments are supported:
* `from` - (Required) The time point of the cursor. The value is a Unix timestamp in seconds, or "begin" or "end".
* `shard_id` - (Required) The ID of the shard.
* `topic_id` - (Required) The ID of the topic.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `log_cursors` - The list of log cursors.
    * `cursor` - The cursor value.
    * `from` - The time point of the cursor.
    * `shard_id` - The ID of the shard.
    * `topic_id` - The ID of the topic.


