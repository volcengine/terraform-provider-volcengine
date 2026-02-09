---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_rules"
sidebar_current: "docs-volcengine-datasource-tls_rules"
description: |-
  Use this data source to query detailed information of tls rules
---
# volcengine_tls_rules
Use this data source to query detailed information of tls rules
## Example Usage
```hcl
data "volcengine_tls_rules" "default" {
  project_id = "47788404-8f1e-49fd-9472-aced5f4bf73f"
  topic_id   = "0a610439-d73f-4680-b365-24eefe98b4fc"
  rule_id    = "33b2607f-e213-42fb-a965-33a0f567ae23"
  log_type   = "delimiter_log"
  pause      = 0
}
```
## Argument Reference
The following arguments are supported:
* `iam_project_name` - (Optional) The iam project name.
* `log_type` - (Optional) The log type.
* `output_file` - (Optional) File name where to save data source results.
* `pause` - (Optional) Whether to pause collection configuration.
* `project_id` - (Optional) The project id.
* `project_name` - (Optional) The project name.
* `rule_id` - (Optional) The rule id.
* `rule_name` - (Optional) The rule name.
* `topic_id` - (Optional) The topic id.
* `topic_name` - (Optional) The topic name.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rules` - The rules list.
    * `container_rule` - Container collection rules.
        * `container_name_regex` - The name of the container to be collected.
        * `env_tag` - Whether to add environment variables as log tags to raw log data.
        * `exclude_container_env_regex` - The container environment variable blacklist is used to specify the range of containers not to be collected.
        * `exclude_container_label_regex` - The container Label blacklist is used to specify the range of containers not to be collected.
        * `include_container_env_regex` - The container environment variable whitelist specifies the container to be collected through the container environment variable. If the whitelist is not enabled, it means that all containers are specified to be collected.
        * `include_container_label_regex` - The container label whitelist specifies the containers to be collected through the container label. If the whitelist is not enabled, all containers are specified to be collected.
        * `kubernetes_rule` - Collection rules for Kubernetes containers.
            * `annotation_tag` - Whether to add Kubernetes Annotation as a log tag to the raw log data.
            * `exclude_pod_label_regex` - Specify the containers not to be collected through the Pod Label blacklist, and not enable means to collect all containers.
            * `include_pod_label_regex` - The Pod Label whitelist is used to specify containers to be collected. When the Pod Label whitelist is not enabled, it means that all containers are collected.
            * `label_tag` - Whether to add Kubernetes Label as a log label to the original log data.
            * `namespace_name_regex` - The name of the Kubernetes Namespace to be collected. If no Namespace name is specified, all containers will be collected. Namespace names support regular matching.
            * `pod_name_regex` - The Pod name is used to specify the container to be collected. When no Pod name is specified, it means to collect all containers.
            * `workload_name_regex` - Specify the container to be collected by the name of the workload. When no workload name is specified, all containers are collected. The workload name supports regular matching.
            * `workload_type` - Specify the container to be collected by the type of workload. Only one type can be selected. When no type is specified, it means to collect all types of containers.
        * `stream` - The collection mode.
    * `create_time` - The creation time.
    * `exclude_paths` - Collect the blacklist list.
        * `type` - Collection path type. The path type can be `File` or `Path`.
        * `value` - Collection path.
    * `extract_rule` - The extract rule.
        * `begin_regex` - The first log line needs to match the regular expression.
        * `delimiter` - The delimiter of the log.
        * `filter_key_regex` - The filter key list.
            * `key` - The name of the filter key.
            * `regex` - The log content of the filter field needs to match the regular expression.
        * `keys` - A list of log field names (Key).
        * `log_regex` - The entire log needs to match the regular expression.
        * `log_template` - Automatically extract log fields according to the specified log template.
            * `format` - Log template content.
            * `type` - The type of the log template.
        * `time_format` - Parsing format of the time field.
        * `time_key` - The field name of the log time field.
        * `un_match_log_key` - When uploading the failed log, the key name of the failed log.
        * `un_match_up_load_switch` - Whether to upload the log of parsing failure.
    * `input_type` - The collection type.
    * `log_sample` - Log sample.
    * `log_type` - The log type.
    * `modify_time` - The modification time.
    * `paths` - Collection path list.
    * `rule_id` - The rule id.
    * `rule_name` - The rule name.
    * `topic_id` - The topic id.
    * `topic_name` - The topic name.
    * `user_define_rule` - User-defined collection rules.
        * `advanced` - LogCollector extension configuration.
            * `close_eof` - Whether to release the log file handle after reading to the end of the log file. The default is false.
            * `close_inactive` - The wait time to release the log file handle. When the log file has not written a new log for more than the specified time, release the handle of the log file.
            * `close_removed` - After the log file is removed, whether to release the handle of the log file. The default is false.
            * `close_renamed` - After the log file is renamed, whether to release the handle of the log file. The default is false.
            * `close_timeout` - The maximum length of time that LogCollector monitors log files. The unit is seconds, and the default is 0 seconds, which means that there is no limit to the length of time LogCollector monitors log files.
        * `enable_raw_log` - Whether to upload raw logs.
        * `fields` - Add constant fields to logs.
        * `parse_path_rule` - Rules for parsing collection paths. After the rules are set, the fields in the collection path will be extracted through the regular expressions specified in the rules, and added to the log data as metadata.
            * `keys` - A list of field names. Log Service will parse the path sample (PathSample) into multiple fields according to the regular expression (Regex), and Keys is used to specify the field name of each field.
            * `path_sample` - Sample capture path for a real scene.
            * `regex` - Regular expression for extracting path fields. It must match the collection path sample, otherwise it cannot be extracted successfully.
        * `plugin` - Plugin configuration. After the plugin configuration is enabled, one or more LogCollector processor plugins can be added to parse logs with complex or variable structures.
            * `processors` - LogCollector plugin.
        * `shard_hash_key` - Rules for routing log partitions. Setting this parameter indicates that the HashKey routing shard mode is used when collecting logs, and Log Service will write the data to the shard containing the specified Key value.
            * `hash_key` - The HashKey of the log group is used to specify the partition (shard) to be written to by the current log group.
        * `tail_files` - LogCollector collection strategy, which specifies whether LogCollector collects incremental logs or full logs. The default is false, which means to collect all logs.
* `total_count` - The total count of query.


