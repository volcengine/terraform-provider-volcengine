---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_rule"
sidebar_current: "docs-volcengine-resource-tls_rule"
description: |-
  Provides a resource to manage tls rule
---
# volcengine_tls_rule
Provides a resource to manage tls rule
## Example Usage
```hcl
resource "volcengine_tls_rule" "foo" {
  topic_id  = "7bfa2cdc-4f8b-4cf9-b4c9-0ed05c33349f"
  rule_name = "test"
  //paths = ["/data/nginx/log/xx.log"]
  log_type   = "minimalist_log"
  log_sample = "2018-05-22 15:35:53.850 INFO XXXX"
  input_type = 1
  #  exclude_paths {
  #    type = "File"
  #    value = "/data/nginx/log/*/*/exclude.log"
  #  }
  #  exclude_paths {
  #    type = "Path"
  #    value = "/data/nginx/log/*/exclude/"
  #  }
  # extract_rule {
  #    delimiter = ""
  #    begin_regex = ""
  #    log_regex = ""
  #    keys = [""]
  #    time_key = ""
  #    time_format = ""
  #    filter_key_regex {
  #      key = "__content__"
  #      regex = ".*ERROR.*"
  #    }
  #    un_match_up_load_switch = true
  #    un_match_log_key = "LogParseFailed"
  #    log_template {
  #      type = ""
  #      format = ""
  #    }
  # }
  user_define_rule {
    enable_raw_log = false
    #    fields = {
    #      cluster_id = "dabaad5f-7a10-4771-b3ea-d821f73e****"
    #    }
    tail_files = true
    #    parse_path_rule {
    #      path_sample = "/data/nginx/log/dabaad5f-7a10/tls/app.log"
    #      regex = "\\/data\\/nginx\\/log\\/(\\w+)-(\\w+)\\/tls\\/app\\.log"
    #      keys = ["instance-id", "pod-name"]
    #    }
    shard_hash_key {
      hash_key = "3C"
    }
    #    plugin {
    #      processors = <<PROCESSORS
    #    {
    #        "json":{
    #            "field":"__content__",
    #            "trim_keys":{
    #              "mode":"all",
    #              "chars":"#"
    #            },
    #            "trim_values":{
    #              "mode":"all",
    #              "chars":"#"
    #            },
    #            "allow_overwrite_keys":true,
    #            "allow_empty_values":true
    #        }
    #    }
    #  PROCESSORS
    #    }
    plugin {
      processors = [
        jsonencode(
          {
            "json" : {
              "field" : "__content__",
              "trim_keys" : {
                "mode" : "all",
                "chars" : "#"
              },
              "trim_values" : {
                "mode" : "all",
                "chars" : "#t"
              },
              "allow_overwrite_keys" : true,
              "allow_empty_values" : true
            },
          },
        ),
        jsonencode(
          {
            "json" : {
              "field" : "__content__",
              "trim_keys" : {
                "mode" : "all",
                "chars" : "#xx"
              },
              "trim_values" : {
                "mode" : "all",
                "chars" : "#txxxt"
              },
              "allow_overwrite_keys" : true,
              "allow_empty_values" : true
            },
          },
        )
      ]
    }
    advanced {
      close_inactive = 10
      close_removed  = false
      close_renamed  = false
      close_eof      = false
      close_timeout  = 1
    }
  }
  container_rule {
    stream               = "all"
    container_name_regex = ".*test.*"
    include_container_label_regex = {
      Key1 = "Value12",
      Key2 = "Value23"
    }
    exclude_container_label_regex = {
      Key1 = "Value12",
      Key2 = "Value22"
    }
    include_container_env_regex = {
      Key1 = "Value1",
      Key2 = "Value2"
    }
    exclude_container_env_regex = {
      Key1 = "Value1",
      Key2 = "Value2"
    }
    env_tag = {
      Key1 = "Value1",
      Key2 = "Value2"
    }
    kubernetes_rule {
      namespace_name_regex = ".*test.*"
      workload_type        = "Deployment"
      workload_name_regex  = ".*test.*"
      include_pod_label_regex = {
        Key1 = "Value1",
        Key2 = "Value2"
      }
      exclude_pod_label_regex = {
        Key1 = "Value1",
        Key2 = "Value2"
      }
      pod_name_regex = ".*test.*"
      label_tag = {
        Key1 = "Value1",
        Key2 = "Value2"
      }
      annotation_tag = {
        Key1 = "Value1",
        Key2 = "Value2"
      }
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `rule_name` - (Required) The name of the collection configuration.
* `topic_id` - (Required, ForceNew) The ID of the log topic to which the collection configuration belongs.
* `container_rule` - (Optional) Container collection rules.
* `exclude_paths` - (Optional) Collect the blacklist list.
* `extract_rule` - (Optional) The extract rule.
* `input_type` - (Optional) The type of the collection configuration. Validate value can be `0`(host log file), `1`(K8s container standard output) and `2`(Log files in the K8s container).
* `log_sample` - (Optional) The sample of the log.
* `log_type` - (Optional) The log type. The value can be one of the following: `minimalist_log`, `json_log`, `delimiter_log`, `multiline_log`, `fullregex_log`.
* `paths` - (Optional) Collection path list.
* `user_define_rule` - (Optional) User-defined collection rules.

The `advanced` object supports the following:

* `close_eof` - (Optional) Whether to release the log file handle after reading to the end of the log file. The default is false.
* `close_inactive` - (Optional) The wait time to release the log file handle. When the log file has not written a new log for more than the specified time, release the handle of the log file.
* `close_removed` - (Optional) After the log file is removed, whether to release the handle of the log file. The default is false.
* `close_renamed` - (Optional) After the log file is renamed, whether to release the handle of the log file. The default is false.
* `close_timeout` - (Optional) The maximum length of time that LogCollector monitors log files. The unit is seconds, and the default is 0 seconds, which means that there is no limit to the length of time LogCollector monitors log files.

The `container_rule` object supports the following:

* `stream` - (Required) The collection mode.
* `container_name_regex` - (Optional) The name of the container to be collected.
* `env_tag` - (Optional) Whether to add environment variables as log tags to raw log data.
* `exclude_container_env_regex` - (Optional) The container environment variable blacklist is used to specify the range of containers not to be collected.
* `exclude_container_label_regex` - (Optional) The container Label blacklist is used to specify the range of containers not to be collected.
* `include_container_env_regex` - (Optional) The container environment variable whitelist specifies the container to be collected through the container environment variable. If the whitelist is not enabled, it means that all containers are specified to be collected.
* `include_container_label_regex` - (Optional) The container label whitelist specifies the containers to be collected through the container label. If the whitelist is not enabled, all containers are specified to be collected.
* `kubernetes_rule` - (Optional) Collection rules for Kubernetes containers.

The `exclude_paths` object supports the following:

* `type` - (Required) Collection path type. The path type can be `File` or `Path`.
* `value` - (Required) Collection path.

The `extract_rule` object supports the following:

* `begin_regex` - (Optional) The first log line needs to match the regular expression.
* `delimiter` - (Optional) The delimiter of the log.
* `filter_key_regex` - (Optional) The filter key list.
* `keys` - (Optional) A list of log field names (Key).
* `log_regex` - (Optional) The entire log needs to match the regular expression.
* `log_template` - (Optional) Automatically extract log fields according to the specified log template.
* `time_format` - (Optional) Parsing format of the time field.
* `time_key` - (Optional) The field name of the log time field.
* `un_match_log_key` - (Optional) When uploading the failed log, the key name of the failed log.
* `un_match_up_load_switch` - (Optional) Whether to upload the log of parsing failure.

The `filter_key_regex` object supports the following:

* `key` - (Required) The name of the filter key.
* `regex` - (Required) The log content of the filter field needs to match the regular expression.

The `kubernetes_rule` object supports the following:

* `annotation_tag` - (Optional) Whether to add Kubernetes Annotation as a log tag to the raw log data.
* `exclude_pod_label_regex` - (Optional) Specify the containers not to be collected through the Pod Label blacklist, and not enable means to collect all containers.
* `include_pod_label_regex` - (Optional) The Pod Label whitelist is used to specify containers to be collected. When the Pod Label whitelist is not enabled, it means that all containers are collected.
* `label_tag` - (Optional) Whether to add Kubernetes Label as a log label to the original log data.
* `namespace_name_regex` - (Optional) The name of the Kubernetes Namespace to be collected. If no Namespace name is specified, all containers will be collected. Namespace names support regular matching.
* `pod_name_regex` - (Optional) The Pod name is used to specify the container to be collected. When no Pod name is specified, it means to collect all containers.
* `workload_name_regex` - (Optional) Specify the container to be collected by the name of the workload. When no workload name is specified, all containers are collected. The workload name supports regular matching.
* `workload_type` - (Optional) Specify the container to be collected by the type of workload. Only one type can be selected. When no type is specified, it means to collect all types of containers.

The `log_template` object supports the following:

* `format` - (Required) Log template content.
* `type` - (Required) The type of the log template.

The `parse_path_rule` object supports the following:

* `keys` - (Optional) A list of field names. Log Service will parse the path sample (PathSample) into multiple fields according to the regular expression (Regex), and Keys is used to specify the field name of each field.
* `path_sample` - (Optional) Sample capture path for a real scene.
* `regex` - (Optional) Regular expression for extracting path fields. It must match the collection path sample, otherwise it cannot be extracted successfully.

The `plugin` object supports the following:

* `processors` - (Required) LogCollector plugin.

The `shard_hash_key` object supports the following:

* `hash_key` - (Required) The HashKey of the log group is used to specify the partition (shard) to be written to by the current log group.

The `user_define_rule` object supports the following:

* `advanced` - (Optional) LogCollector extension configuration.
* `enable_raw_log` - (Optional) Whether to upload raw logs.
* `fields` - (Optional) Add constant fields to logs.
* `parse_path_rule` - (Optional) Rules for parsing collection paths. After the rules are set, the fields in the collection path will be extracted through the regular expressions specified in the rules, and added to the log data as metadata.
* `plugin` - (Optional) Plugin configuration. After the plugin configuration is enabled, one or more LogCollector processor plugins can be added to parse logs with complex or variable structures.
* `shard_hash_key` - (Optional) Rules for routing log partitions. Setting this parameter indicates that the HashKey routing shard mode is used when collecting logs, and Log Service will write the data to the shard containing the specified Key value.
* `tail_files` - (Optional) LogCollector collection strategy, which specifies whether LogCollector collects incremental logs or full logs. The default is false, which means to collect all logs.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `rule_id` - The id of the rule.


## Import
tls rule can be imported using the id, e.g.
```
$ terraform import volcengine_tls_rule.default fa************
```

