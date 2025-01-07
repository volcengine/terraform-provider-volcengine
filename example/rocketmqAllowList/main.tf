resource "volcengine_rocketmq_allow_list" "foo" {
  allow_list_name = "acc-test-allow-list"
  allow_list_desc = "acc-test"
  allow_list      = ["192.168.0.0/24", "192.168.2.0/24"]
}
