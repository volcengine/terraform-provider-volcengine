resource "volcengine_vefaas_kafka_trigger" "foo" {
  function_id = "f0zvcxxx"
  name = "123"
  mq_instance_id = "kafka-cnnguc4426wysxxx"
  topic_name = "topic-1"
  kafka_credentials {
    password = "Wasdfgg123"
    username = "test-1"
    mechanism = "PLAIN"
  }
  lifecycle {
    ignore_changes = [kafka_credentials]
  }
}
