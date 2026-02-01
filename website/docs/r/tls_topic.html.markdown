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
  project_id      = "bdb87e4d-7dad-4b96-ac43-e1b09e9dc8ac"
  topic_name      = "tf-topic-5"
  description     = "test"
  ttl             = 60
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
  log_public_ip  = true
  enable_hot_ttl = true
  hot_ttl        = 30
  cold_ttl       = 30
  archive_ttl    = 0
  encrypt_conf {
    enable       = true
    encrypt_type = "default"
  }
}
```
## Argument Reference
The following arguments are supported:
* `project_id` - (Required, ForceNew) The project id of the tls topic.
* `shard_count` - (Required, ForceNew) The count of shards in the tls topic. Valid value range: 1-10. This field is only valid when creating tls topic.
* `topic_name` - (Required) The name of the tls topic.
* `ttl` - (Required) The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
* `archive_ttl` - (Optional) Archive storage duration, valid when enable_hot_ttl is true.
* `auto_split` - (Optional) Whether to enable automatic partition splitting function of the tls topic.
true: (default) When the amount of data written exceeds the capacity of existing partitions for 5 consecutive minutes, Log Service will automatically split partitions based on the data volume to meet business needs. However, the number of partitions after splitting cannot exceed the maximum number of partitions. Newly split partitions within the last 15 minutes will not be automatically split again.
false: Disables automatic partition splitting.
* `cold_ttl` - (Optional) Infrequent storage duration, valid when enable_hot_ttl is true.
* `description` - (Optional) The description of the tls project.
* `enable_hot_ttl` - (Optional) Whether to enable tiered storage.
* `enable_tracking` - (Optional) Whether to enable WebTracking function of the tls topic.
* `encrypt_conf` - (Optional) Data encryption configuration.
* `hot_ttl` - (Optional) Standard storage duration, valid when enable_hot_ttl is true.
* `log_public_ip` - (Optional) Whether to enable the function of recording public IP.
* `manual_split_shard_id` - (Optional) The id of shard to be manually split. This field is valid only when modifying the topic. 
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `manual_split_shard_number` - (Optional) The split number of shard. The valid number should be a non-zero even number, such as 2, 4, 8, or 16. The total number of read-write status shards after splitting cannot exceed 50. 
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `max_split_shard` - (Optional) The maximum number of partitions, which is the maximum number of partitions after partition splitting. The value range is 1 to 10, with a default of 10.
* `tags` - (Optional) Tags.
* `time_format` - (Optional) The format of the time field.
* `time_key` - (Optional) The name of the time field.

The `encrypt_conf` object supports the following:

* `enable` - (Optional) Whether to enable data encryption.
* `encrypt_type` - (Optional) The encryption type.
* `user_cmk_info` - (Optional) The user custom key.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

The `user_cmk_info` object supports the following:

* `region_id` - (Optional) The key region.
* `trn` - (Optional) The key trn.
* `user_cmk_id` - (Optional) The key id.

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

