resource "volcengine_tls_etl_task" "foo" {
  dsl_type = "NORMAL"
  description = "for-tf-test"
  enable = "false"
  from_time = 1750649545
  name = "tf-test-etl-task"
  script = ""
  source_topic_id = "9b756385-1dfb-4306-a094-0c88e04b34a5"
  to_time = 1750735958
  target_resources {
    alias = "tf-test-1"
    topic_id = "a690a9b8-72c1-40a3-b8c6-f89a81d3748e"
  }
  target_resources {
    alias = "tf-test-2-modify"
    topic_id = "bdf4f23b-a889-456c-ac5f-09d727427557"
  }
  task_type = "Resident"
}

