---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_topic"
sidebar_current: "docs-volcengine-resource-tls_topic"
description: |-
  Provides a resource to manage tls topic
---
# volcengine_tls_topic
Provides a resource to manage tls topic
## Example Usage
```hcl
resource "volcengine_tls_topic" "foo" {
  project_id      = "e020c978-4f05-40e1-9167-0113d3ef****"
  topic_name      = "tf-test-topic"
  description     = "test"
  ttl             = 10
  shard_count     = 2
  auto_split      = true
  max_split_shard = 10
  enable_tracking = true
  time_key        = "request_time"
  time_format     = "%Y-%m-%dT%H:%M:%S,%f"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `project_id` - (Required, ForceNew) The project id of the tls topic.
* `shard_count` - (Required, ForceNew) The count of shards in the tls topic. Valid value range: 1-10.
* `topic_name` - (Required) The name of the tls topic.
* `ttl` - (Required) The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
* `auto_split` - (Optional) Whether to enable automatic partition splitting function of the tls topic.
true: (default) When the amount of data written exceeds the capacity of existing partitions for 5 consecutive minutes, Log Service will automatically split partitions based on the data volume to meet business needs. However, the number of partitions after splitting cannot exceed the maximum number of partitions. Newly split partitions within the last 15 minutes will not be automatically split again.
false: Disables automatic partition splitting.
* `description` - (Optional) The description of the tls project.
* `enable_tracking` - (Optional) Whether to enable WebTracking function of the tls topic.
* `max_split_shard` - (Optional) The maximum number of partitions, which is the maximum number of partitions after partition splitting. The value range is 1 to 10, with a default of 10.
* `tags` - (Optional) Tags.
* `time_format` - (Optional) The format of the time field.
* `time_key` - (Optional) The name of the time field.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of the tls topic.
* `modify_time` - The modify time of the tls topic.


## Import
Tls Topic can be imported using the id, e.g.
```
$ terraform import volcengine_tls_topic.default edf051ed-3c46-49ba-9339-bea628fe****
```

