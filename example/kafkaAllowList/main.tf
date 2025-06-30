resource "volcengine_kafka_allow_list" "foo" {
    allow_list = ["192.168.0.1", "10.32.55.66", "10.22.55.66"]
    allow_list_name = "tf-test"
}