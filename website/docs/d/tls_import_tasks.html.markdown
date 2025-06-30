---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_import_tasks"
sidebar_current: "docs-volcengine-datasource-tls_import_tasks"
description: |-
  Use this data source to query detailed information of tls import tasks
---
# volcengine_tls_import_tasks
Use this data source to query detailed information of tls import tasks
## Example Usage
```hcl
data "volcengine_tls_import_tasks" "foo" {

}
```
## Argument Reference
The following arguments are supported:
* `iam_project_name` - (Optional) Specify the IAM project name to query the data import tasks under the specified IAM project.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_id` - (Optional) Specify the log item ID for querying the data import tasks under the specified log item.
* `project_name` - (Optional) Specify the name of the log item for querying the data import tasks under the specified log item. Support fuzzy query..
* `source_type` - (Optional) Specify the import type for querying the data import tasks related to this import type.
* `status` - (Optional) Specify the status of the import task.
* `task_id` - (Optional) Import the task ID of the data to be queried.
* `task_name` - (Optional) Import the task name of the data to be queried.
* `topic_id` - (Optional) Specify the log topic ID for querying the data import tasks related to this log topic.
* `topic_name` - (Optional) Specify the name of the log topic for querying the data import tasks related to this log topic. Support fuzzy query.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `task_info` - Data import task list.
    * `create_time` - The creation time of the data import task.
    * `description` - Data import task description.
    * `import_source_info` - The source information of the data import task.
        * `kafka_source_info` - TOS imports source information.
            * `encode` - The encoding format of the data.
            * `group` - Kafka consumer group.
            * `host` - The service addresses corresponding to different types of Kafka clusters are different.
            * `initial_offset` - The starting position of data import.
            * `instance_id` - When you are using the Volcano Engine Message Queue Kafka version, it should be set to the Kafka instance ID.
            * `is_need_auth` - Whether to enable authentication.
            * `mechanism` - Password authentication mechanism.
            * `password` - The Kafka SASL user password used for identity authentication.
            * `protocol` - Secure Transport protocol.
            * `time_source_default` - Specify the log time.
            * `topic` - Kafka Topic name.
            * `username` - The Kafka SASL username used for identity authentication.
        * `tos_source_info` - TOS imports source information.
            * `bucket` - The TOS bucket where the log file is located.
            * `compress_type` - The compression mode of data in the TOS bucket.
            * `prefix` - The path of the file to be imported in the TOS bucket.
            * `region` - The region where the TOS bucket is located. Support cross-regional data import.
    * `project_id` - Specify the log item ID for querying the data import tasks under the specified log item.
    * `project_name` - Specify the name of the log item for querying the data import tasks under the specified log item. Support fuzzy query..
    * `source_type` - Specify the import type for querying the data import tasks related to this import type.
    * `target_info` - The output information of the data import task.
        * `extract_rule` - Log extraction rules.
            * `begin_regex` - The regular expression used to identify the first line in each log, and its matching part will serve as the beginning of the log.
            * `delimiter` - Log delimiter.
            * `keys` - List of log field names (Keys).
            * `quote` - Reference symbol. The content wrapped by the reference will not be separated but will be parsed into a complete field. It is valid if and only if the LogType is delimiter_log.
            * `skip_line_count` - The number of log lines skipped.
            * `time_extract_regex` - A regular expression for extracting time, used to extract the time value in the TimeKey field and parse it into the corresponding collection time.
            * `time_format` - The parsing format of the time field.
            * `time_key` - The field name of the log time field.
            * `time_zone` - Time zone, supporting both machine time zone (default) and custom time zone. Among them, the custom time zone supports GMT and UTC.
            * `un_match_log_key` - When uploading a log that failed to parse, the key name of the parse failed log.
            * `un_match_up_load_switch` - Whether to upload the logs of failed parsing.
        * `log_sample` - Log sample.
        * `log_type` - Specify the log parsing type when importing.
        * `region` - Regional ID.
    * `task_id` - Import the task ID of the data to be queried.
    * `task_name` - Import the task name of the data to be queried.
    * `task_statistics` - The progress of the data import task.
        * `bytes_total` - The total number of resource bytes that have been listed.
        * `bytes_transferred` - The number of imported bytes.
        * `failed` - The number of resources that failed to import.
        * `not_exist` - The number of non-existent resources.
        * `skipped` - Skip the number of imported resources.
        * `task_status` - Import the status of the task.
        * `total` - The total number of resources that have been listed.
        * `transferred` - The number of imported resources.
    * `topic_id` - Specify the log topic ID for querying the data import tasks related to this log topic.
    * `topic_name` - Specify the name of the log topic for querying the data import tasks related to this log topic. Support fuzzy query.
* `total_count` - The total count of query.


