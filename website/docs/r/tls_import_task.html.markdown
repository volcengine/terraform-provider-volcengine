---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_import_task"
sidebar_current: "docs-volcengine-resource-tls_import_task"
description: |-
  Provides a resource to manage tls import task
---
# volcengine_tls_import_task
Provides a resource to manage tls import task
## Example Usage
```hcl
resource "volcengine_tls_import_task" "foo" {
  description = "tf-test"
  import_source_info {
    kafka_source_info {
      encode              = "UTF-8"
      host                = "1.1.1.1"
      initial_offset      = 0
      time_source_default = 1
      topic               = "topic-1,topic-2,topic-3"
    }
  }
  source_type = "kafka"
  target_info {
    region   = "cn-beijing"
    log_type = "json_log"
    extract_rule {
      un_match_log_key        = "key-failed"
      un_match_up_load_switch = true
    }
  }
  task_name = "tf-test-task-name-kafka"
  topic_id  = "b966e41a-d6a6-4999-bd75-39xxxxxxx"
}
```
## Argument Reference
The following arguments are supported:
* `import_source_info` - (Required) The source information of the data import task.
* `target_info` - (Required) The output information of the data import task.
* `description` - (Optional) Data import task description.
* `project_id` - (Optional) The log project ID used for storing data.
* `source_type` - (Optional) Import the source type.
* `status` - (Optional) The status of the data import task.
* `task_name` - (Optional) Data import task name.
* `topic_id` - (Optional) The log topic ID used for storing data.

The `extract_rule` object supports the following:

* `begin_regex` - (Optional) The regular expression used to identify the first line in each log, and its matching part will serve as the beginning of the log.
* `delimiter` - (Optional) Log delimiter.
* `keys` - (Optional) List of log field names (Keys).
* `quote` - (Optional) Reference symbol. The content wrapped by the reference will not be separated but will be parsed into a complete field. It is valid if and only if the LogType is delimiter_log.
* `skip_line_count` - (Optional) The number of log lines skipped.
* `time_extract_regex` - (Optional) A regular expression for extracting time, used to extract the time value in the TimeKey field and parse it into the corresponding collection time.
* `time_format` - (Optional) The parsing format of the time field.
* `time_key` - (Optional) The field name of the log time field.
* `time_zone` - (Optional) Time zone, supporting both machine time zone (default) and custom time zone. Among them, the custom time zone supports GMT and UTC.
* `un_match_log_key` - (Optional) When uploading a log that failed to parse, the key name of the parse failed log.
* `un_match_up_load_switch` - (Optional) Whether to upload the logs of failed parsing.

The `import_source_info` object supports the following:

* `kafka_source_info` - (Optional) TOS imports source information.
* `tos_source_info` - (Optional) TOS imports source information.

The `kafka_source_info` object supports the following:

* `encode` - (Optional) The encoding format of the data.
* `group` - (Optional) Kafka consumer group.
* `host` - (Optional) The service addresses corresponding to different types of Kafka clusters are different.
* `initial_offset` - (Optional) The starting position of data import.
* `instance_id` - (Optional) When you are using the Volcano Engine Message Queue Kafka version, it should be set to the Kafka instance ID.
* `is_need_auth` - (Optional) Whether to enable authentication.
* `mechanism` - (Optional) Password authentication mechanism.
* `password` - (Optional) The Kafka SASL user password used for identity authentication.
* `protocol` - (Optional) Secure Transport protocol.
* `time_source_default` - (Optional) Specify the log time.
* `topic` - (Optional) Kafka Topic name.
* `username` - (Optional) The Kafka SASL username used for identity authentication.

The `target_info` object supports the following:

* `log_type` - (Required) Specify the log parsing type when importing.
* `region` - (Required) Regional ID.
* `extract_rule` - (Optional) Log extraction rules.
* `log_sample` - (Optional) Log sample.

The `tos_source_info` object supports the following:

* `bucket` - (Optional) The TOS bucket where the log file is located.
* `compress_type` - (Optional) The compression mode of data in the TOS bucket.
* `prefix` - (Optional) The path of the file to be imported in the TOS bucket.
* `region` - (Optional) The region where the TOS bucket is located. Support cross-regional data import.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ImportTask can be imported using the id, e.g.
```
$ terraform import volcengine_import_task.default resource_id
```

