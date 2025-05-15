resource "volcengine_vefaas_kafka_trigger" "foo" {
  function_id = "35ybaxxx"
  name = "tf-123"
  mq_instance_id = "kafka-cnngmbeq10mcxxxx"
  topic_name = "topic"
  kafka_credentials {
    password = "Waxxxxxx"
    username = "test-1"
    mechanism = "PLAIN"
  }
  batch_size = 100
  description = "modify"
  lifecycle {
    ignore_changes = [kafka_credentials]
  }
}
