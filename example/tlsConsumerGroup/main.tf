resource "volcengine_tls_consumer_group" "foo" {
  project_id = "17ba378d-de43-495e-8906-03aexxxxxx"
  topic_id_list = ["0ed72ac8-9531-4967-b216-ac30xxxxxx"]
  consumer_group_name = "tf-test-consumer-group"
  heartbeat_ttl = 120
  ordered_consume = false
}