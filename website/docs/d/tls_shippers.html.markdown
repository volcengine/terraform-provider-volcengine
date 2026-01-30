---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_shippers"
sidebar_current: "docs-volcengine-datasource-tls_shippers"
description: |-
  Use this data source to query detailed information of tls shippers
---
# volcengine_tls_shippers
Use this data source to query detailed information of tls shippers
## Example Usage
```hcl
data "volcengine_tls_shippers" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `iam_project_name` - (Optional) Specify the IAM project name for querying the data delivery configuration under the specified IAM project.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_id` - (Optional) Specify the log item ID for querying the data delivery configuration under the specified log item.
* `project_name` - (Optional) Specify the name of the log item for querying the data delivery configuration under the specified log item. Support fuzzy matching.
* `shipper_id` - (Optional) Delivery configuration ID.
* `shipper_name` - (Optional) Delivery configuration name.
* `shipper_type` - (Optional) Specify the delivery type for querying the delivery configuration related to that delivery type.
* `topic_id` - (Optional) Specify the log topic ID for querying the data delivery configuration related to this log topic.
* `topic_name` - (Optional) Specify the name of the log topic for querying the data delivery configuration related to this log topic. Support fuzzy matching.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `shippers` - Submit the relevant information of the configuration.
    * `content_info` - The content format configuration of the delivery log.
        * `csv_info` - CSV format log content configuration.
            * `delimiter` - Delimiters are supported, including commas, tabs, vertical bars, semicolons, and Spaces.
            * `escape_char` - When the field content contains a delimiter, use an escape character to wrap the field. Currently, only single quotes, double quotes, and null characters are supported.
            * `keys` - Configure the fields that need to be delivered.
            * `non_field_content` - Invalid field filling content, with a length ranging from 0 to 128.
            * `print_header` - Whether to print the Key on the first line.
        * `format` - Log content parsing format.
        * `json_info` - JSON format log content configuration.
            * `enable` - Enable the flag.
            * `escape` - Whether to escape or not. It must be configured as true.
            * `keys` - When delivering in JSON format, if this parameter is not configured, it indicates that all fields have been delivered. Including __content__ (choice), __source__, __path__, __time__, __image_name__, __container_name__, __pod_name__, __pod_uid__, namespace, __tag____client_ip__, __tag____receive_time__.
    * `create_time` - Processing task creation time.
    * `dashboard_id` - The default built-in dashboard ID for delivery.
    * `kafka_shipper_info` - JSON format log content configuration.
        * `compress` - Compression formats currently supported include snappy, gzip, lz4, and none.
        * `end_time` - Delivery end time, millisecond timestamp. If not configured, it will keep delivering.
        * `instance` - Kafka instance.
        * `kafka_topic` - The name of the Kafka Topic.
        * `start_time` - Delivery start time, millisecond timestamp. If not configured, the default is the current time.
    * `modify_time` - The most recent modification time of the processing task.
    * `project_id` - The log project ID where the log to be delivered is located.
    * `project_name` - The name of the log item where the log to be delivered is located.
    * `role_trn` - The role trn.
    * `shipper_end_time` - Delivery end time, millisecond timestamp. If not configured, it will keep delivering.
    * `shipper_id` - Deliver configuration ID.
    * `shipper_name` - Delivery configuration name.
    * `shipper_start_time` - Delivery start time, millisecond timestamp. If not configured, it defaults to the current time.
    * `shipper_type` - The type of delivery.
    * `status` - Whether to enable the delivery configuration.
    * `topic_id` - The log topic ID where the log to be delivered is located.
    * `topic_name` - The name of the log topic where the log to be delivered is located.
    * `tos_shipper_info` - Deliver the relevant configuration to the object storage (TOS).
        * `bucket` - When choosing a TOS bucket, it must be located in the same region as the source log topic.
        * `compress` - Compression formats currently supported include snappy, gzip, lz4, and none.
        * `interval` - The delivery time interval, measured in seconds, ranges from 300 to 900.
        * `max_size` - The maximum size of the original file that can be delivered to each partition (Shard), that is, the size of the uncompressed log file. The unit is MiB, and the value range is 5 to 256.
        * `partition_format` - Partition rules for delivering logs.
        * `prefix` - The top-level directory name of the storage bucket. All log data delivered through this delivery configuration will be delivered to this directory.
* `total_count` - The total count of query.


