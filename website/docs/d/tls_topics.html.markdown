---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_topics"
sidebar_current: "docs-volcengine-datasource-tls_topics"
description: |-
  Use this data source to query detailed information of tls topics
---
# volcengine_tls_topics
Use this data source to query detailed information of tls topics
## Example Usage
```hcl
data "volcengine_tls_topics" "default" {
  project_id = "bdb87e4d-7dad-4b96-ac43-e1b09e9dc8ac"
  topic_id   = "9b756385-1dfb-4306-a094-0c88e04b34a5"
}
```
## Argument Reference
The following arguments are supported:
* `project_id` - (Required) The project id of tls topic.
* `name_regex` - (Optional) A Name Regex of tls topic.
* `output_file` - (Optional) File name where to save data source results.
* `tags` - (Optional) Tags.
* `topic_id` - (Optional) The id of tls topic. This field supports fuzzy queries. It is not supported to specify both TopicName and TopicId at the same time.
* `topic_name` - (Optional) The name of tls topic. This field supports fuzzy queries. It is not supported to specify both TopicName and TopicId at the same time.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `tls_topics` - The collection of tls topic query.
    * `archive_ttl` - Archive storage duration, valid when enable_hot_ttl is true.
    * `auto_split` - Whether to enable automatic partition splitting function of the tls topic.
    * `cold_ttl` - Infrequent storage duration, valid when enable_hot_ttl is true.
    * `create_time` - The create time of the tls topic.
    * `description` - The description of the tls topic.
    * `enable_hot_ttl` - Whether to enable tiered storage.
    * `enable_tracking` - Whether to enable WebTracking function of the tls topic.
    * `encrypt_conf` - Data encryption configuration.
        * `enable` - Whether to enable data encryption.
        * `encrypt_type` - The encryption type.
        * `user_cmk_info` - The user custom key.
            * `region_id` - The key region.
            * `trn` - The key trn.
            * `user_cmk_id` - The key id.
    * `hot_ttl` - Standard storage duration, valid when enable_hot_ttl is true.
    * `id` - The ID of the tls topic.
    * `log_public_ip` - Whether to enable the function of recording public IP.
    * `max_split_shard` - The max count of shards in the tls topic.
    * `modify_time` - The modify time of the tls topic.
    * `project_id` - The project id of the tls topic.
    * `shard_count` - The count of shards in the tls topic.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `time_format` - The format of the time field.
    * `time_key` - The name of the time field.
    * `topic_id` - The ID of the tls topic.
    * `topic_name` - The name of the tls topic.
    * `ttl` - The data storage time of the tls topic. Unit: Day.
* `total_count` - The total count of tls topic query.


