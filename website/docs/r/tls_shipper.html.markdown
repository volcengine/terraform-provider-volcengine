---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_shipper"
sidebar_current: "docs-volcengine-resource-tls_shipper"
description: |-
  Provides a resource to manage tls shipper
---
# volcengine_tls_shipper
Provides a resource to manage tls shipper
## Example Usage
```hcl
resource "volcengine_tls_shipper" "foo" {
  content_info {
    format = "json"
    json_info {
      enable = true
      keys   = ["__content", "__pod_name__"]
    }
  }
  shipper_end_time   = 1751255700021
  shipper_name       = "tf-test-modify"
  shipper_start_time = 1750737324521
  shipper_type       = "tos"
  topic_id           = "8ba48bd7-2493-4300-b1d0-cb760b89e51b"
  role_trn           = ""
  tos_shipper_info {
    bucket           = "tf-test"
    prefix           = "terraform_1.9.4_linux_amd64.zip"
    max_size         = 50
    interval         = 200
    compress         = "snappy"
    partition_format = "%Y/%m/%d/%H/%M"
  }
}
```
## Argument Reference
The following arguments are supported:
* `content_info` - (Required) Configuration of the delivery format for log content.
* `shipper_name` - (Required) Delivery configuration name.
* `topic_id` - (Required, ForceNew) The log topic ID where the log to be delivered is located.
* `kafka_shipper_info` - (Optional) JSON format log content configuration.
* `role_trn` - (Optional) The role trn.
* `shipper_end_time` - (Optional) Delivery end time, millisecond timestamp. If not configured, it will keep delivering. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `shipper_start_time` - (Optional) Delivery start time, millisecond timestamp. If not configured, it defaults to the current time. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `shipper_type` - (Optional) The type of delivery.
* `status` - (Optional) Whether to enable the delivery configuration. The default value is true.
* `tos_shipper_info` - (Optional) Deliver the relevant configuration to the object storage (TOS).

The `content_info` object supports the following:

* `csv_info` - (Optional) CSV format log content configuration.
* `format` - (Optional) Log content parsing format.
* `json_info` - (Optional) JSON format log content configuration.

The `csv_info` object supports the following:

* `delimiter` - (Required) Delimiters are supported, including commas, tabs, vertical bars, semicolons, and Spaces.
* `escape_char` - (Required) When the field content contains a delimiter, use an escape character to wrap the field. Currently, only single quotes, double quotes, and null characters are supported.
* `keys` - (Required) Configure the fields that need to be delivered.
* `non_field_content` - (Required) Invalid field filling content, with a length ranging from 0 to 128.
* `print_header` - (Required) Whether to print the Key on the first line.

The `json_info` object supports the following:

* `enable` - (Required) Enable the flag.
* `escape` - (Optional) Whether to escape or not. It must be configured as true.
* `keys` - (Optional) When delivering in JSON format, if this parameter is not configured, it indicates that all fields have been delivered. Including __content__ (choice), __source__, __path__, __time__, __image_name__, __container_name__, __pod_name__, __pod_uid__, namespace, __tag____client_ip__, __tag____receive_time__.

The `kafka_shipper_info` object supports the following:

* `compress` - (Required) Compression formats currently supported include snappy, gzip, lz4, and none.
* `instance` - (Required) Kafka instance.
* `kafka_topic` - (Required) The name of the Kafka Topic.
* `end_time` - (Optional, ForceNew) Delivery end time, millisecond timestamp. If not configured, it will keep delivering.
* `start_time` - (Optional, ForceNew) Delivery start time, millisecond timestamp. If not configured, the default is the current time.

The `tos_shipper_info` object supports the following:

* `bucket` - (Required, ForceNew) When choosing a TOS bucket, it must be located in the same region as the source log topic.
* `compress` - (Optional) Compression formats currently supported include snappy, gzip, lz4, and none.
* `interval` - (Optional) The delivery time interval, measured in seconds, ranges from 300 to 900.
* `max_size` - (Optional) The maximum size of the original file that can be delivered to each partition (Shard), that is, the size of the uncompressed log file. The unit is MiB, and the value range is 5 to 256.
* `partition_format` - (Optional) Partition rules for delivering logs.
* `prefix` - (Optional) The top-level directory name of the storage bucket. All log data delivered through this delivery configuration will be delivered to this directory.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Shipper can be imported using the id, e.g.
```
$ terraform import volcengine_shipper.default resource_id
```

