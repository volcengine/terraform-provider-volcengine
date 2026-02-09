---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_shard"
sidebar_current: "docs-volcengine-resource-tls_shard"
description: |-
  Provides a resource to manage tls shard
---
# volcengine_tls_shard
Provides a resource to manage tls shard
## Example Usage
```hcl
resource "volcengine_tls_shard" "foo" {
  topic_id = "176b62c7-c482-4a6e-b983-4697fda9294a"
  shard_id = 1
  number   = 2
}
```
## Argument Reference
The following arguments are supported:
* `number` - (Required) The number of splits. Must be a non-zero even number, such as 2, 4, 8, or 16.
* `shard_id` - (Required, ForceNew) The ID of the shard to split.
* `topic_id` - (Required, ForceNew) The ID of the topic.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `shards` - The collection of shards after split.
    * `exclusive_end_key` - The exclusive end key of the shard.
    * `inclusive_begin_key` - The inclusive begin key of the shard.
    * `modify_time` - The modification time of the shard.
    * `shard_id` - The ID of the shard.
    * `status` - The status of the shard.
    * `stop_write_time` - The stop write time of the shard.
    * `topic_id` - The ID of the topic.


## Import
The TlsShard is not support import.

