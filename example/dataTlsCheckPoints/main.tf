data "volcengine_tls_check_points" "default" {
  project_id          = "7a8ac13e-8e3e-4392-ae77-aea8efa49bbf"
  topic_id            = "33124cc3-15c4-4cdc-9a8a-cc64a9d593dd"
  shard_id            = "0"
  consumer_group_name = "tf-consumer-group"
}

