resource "volcengine_tls_consumer_group" "foo" {
  project_id = "7a8ac13e-8e3e-4392-ae77-aea8efa49bbf"
  topic_id_list = ["33124cc3-15c4-4cdc-9a8a-cc64a9d593dd","9c5c57de-c39f-4777-a4e2-d9b1e69688db"]
  consumer_group_name = "tf-consumer-group-hhh"
  heartbeat_ttl = 120
  ordered_consume = true
}