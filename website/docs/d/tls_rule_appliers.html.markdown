---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_rule_appliers"
sidebar_current: "docs-volcengine-datasource-tls_rule_appliers"
description: |-
  Use this data source to query detailed information of tls rule appliers
---
# volcengine_tls_rule_appliers
Use this data source to query detailed information of tls rule appliers
## Example Usage
```hcl
variable "name" {
  default = "tf-test-rule-applier"
}

resource "volcengine_tls_project" "foo" {
  project_name = var.name
  description  = "tf-test-project-desc"
}

resource "volcengine_tls_topic" "foo" {
  project_id      = volcengine_tls_project.foo.id
  topic_name      = var.name
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
}

resource "volcengine_tls_rule" "foo" {
  topic_id   = volcengine_tls_topic.foo.id
  rule_name  = "tf-test-rule"
  log_type   = "delimiter_log"
  log_sample = "2018-05-22 15:35:53.850,INFO,XXXX"
  input_type = 1

  #  exclude_paths {
  #    type = "File"
  #    value = "/data/nginx/log/*/*/exclude.log"
  #  }
  #  exclude_paths {
  #    type = "Path"
  #    value = "/data/nginx/log/*/exclude/"
  #  }

  extract_rule {
    delimiter   = ","
    keys        = ["time", "level", "msg"]
    time_key    = "time"
    time_format = "%Y-%m-%d %H:%M:%S.%f"
    quote       = "\""
    time_zone   = "GMT+08:00"
    #    begin_regex = ""
    #    log_regex = ""
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
  }

  user_define_rule {
    enable_raw_log = true
    tail_files     = true
    #    fields = {
    #      cluster_id = "dabaad5f-7a10-4771-b3ea-d821f73e****"
    #    }
    #    parse_path_rule {
    #      path_sample = "/data/nginx/log/dabaad5f-7a10/tls/app.log"
    #      regex = "\\/data\\/nginx\\/log\\/(\\w+)-(\\w+)\\/tls\\/app\\.log"
    #      keys = ["instance-id", "pod-name"]
    #    }
    shard_hash_key {
      hash_key = "3C"
    }
    #    plugin {
    #      processors = [
    #        jsonencode(
    #        {
    #          "json":{
    #            "field":"__content__",
    #            "trim_keys":{
    #              "mode":"all",
    #              "chars":"#"
    #            },
    #            "trim_values":{
    #              "mode":"all",
    #              "chars":"#t"
    #            },
    #            "allow_overwrite_keys":true,
    #            "allow_empty_values":true
    #          },
    #        },
    #      ),
    #      ]
    #    }
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

resource "volcengine_tls_host_group" "foo" {
  host_group_name = var.name
  host_group_type = "Label"
  host_identifier = "tf-controller"
  auto_update     = false
  service_logging = false
}

# resource "volcengine_tls_rule_applier" "foo" {
#   rule_id       = volcengine_tls_rule.foo.id
#   host_group_id = volcengine_tls_host_group.foo.id
# }

data "volcengine_tls_rule_appliers" "foo" {
  rule_id = volcengine_tls_rule.foo.id
  #depends_on = [volcengine_tls_rule_applier.foo]
}

# output "foo_host_group_infos" {
#   value = data.volcengine_tls_rule_appliers.foo.host_group_infos
# }
```
## Argument Reference
The following arguments are supported:
* `rule_id` - (Required) The rule id.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `host_group_infos` - The host group info list.
    * `auto_update` - Whether to auto update.
    * `create_time` - The create time.
    * `host_count` - The host count.
    * `host_group_id` - The host group id.
    * `host_group_name` - The host group name.
    * `host_group_type` - The host group type.
    * `host_identifier` - The host identifier.
    * `iam_project_name` - The iam project name.
    * `modify_time` - The modify time.
    * `rule_count` - The rule count.
    * `service_logging` - Whether to service logging.
    * `update_end_time` - The update end time.
    * `update_start_time` - The update start time.
* `total_count` - The total count of query.


