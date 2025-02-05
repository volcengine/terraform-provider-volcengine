resource "volcengine_mongodb_allow_list" "foo" {
  allow_list_name = "acc-test-allow-list"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = "10.1.1.3,10.2.3.0/24,10.1.1.1"
  project_name    = "default"
}