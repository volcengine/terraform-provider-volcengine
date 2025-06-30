resource "volcengine_tls_import_task" "foo" {
  description = "tf-test"
  import_source_info {
    kafka_source_info {
      encode = "UTF-8"
      host = "1.1.1.1"
      initial_offset = 0
      time_source_default = 1
      topic = "topic-1,topic-2,topic-3"
    }
  }
  source_type = "kafka"
  target_info {
    region = "cn-beijing"
    log_type = "json_log"
    extract_rule {
      un_match_log_key = "key-failed"
      un_match_up_load_switch = true
    }
  }
  task_name = "tf-test-task-name-kafka"
  topic_id = "b966e41a-d6a6-4999-bd75-39xxxxxxx"
}