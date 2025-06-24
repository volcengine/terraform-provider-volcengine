resource "volcengine_tls_etl_task" "foo" {
  dsl_type = "NORMAL"
  description = "for-tf-test"
  enable = "true"
  from_time = 1750649545
  name = "tf-test-etl-task-1"
  script = ""
  source_topic_id = "8ba48bd7-2493-4300-b1d0-cb7xxxxxxx"
  to_time = 1750735958
  target_resources {
    alias = "tf-test-1"
    topic_id = "b966e41a-d6a6-4999-bd75-39962xxxxxx"
  }
  target_resources {
    alias = "tf-test-2"
    topic_id = "0ed72ac8-9531-4967-b216-ac3xxxxx"
  }
  task_type = "Resident"

}

