resource "volcengine_tls_rule" "foo" {
  topic_id = "7bfa2cdc-4f8b-4cf9-b4c9-0ed05c33349f"
  rule_name = "test"
  //paths = ["/data/nginx/log/xx.log"]
  log_type = "minimalist_log"
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
          "json":{
            "field":"__content__",
            "trim_keys":{
              "mode":"all",
              "chars":"#"
            },
            "trim_values":{
              "mode":"all",
              "chars":"#t"
            },
            "allow_overwrite_keys":true,
            "allow_empty_values":true
          },
        },
      ),
        jsonencode(
          {
            "json":{
              "field":"__content__",
              "trim_keys":{
                "mode":"all",
                "chars":"#xx"
              },
              "trim_values":{
                "mode":"all",
                "chars":"#txxxt"
              },
              "allow_overwrite_keys":true,
              "allow_empty_values":true
            },
          },
        )
        ]
    }
    advanced {
      close_inactive = 10
      close_removed = false
      close_renamed = false
      close_eof = false
      close_timeout = 1
    }
  }
  container_rule {
    stream = "all"
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
      workload_type = "Deployment"
      workload_name_regex = ".*test.*"
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