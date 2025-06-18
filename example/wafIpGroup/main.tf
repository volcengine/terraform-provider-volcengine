resource "volcengine_waf_ip_group" "foo" {
  add_type = "List"
  ip_list = ["1.1.1.1", "1.1.1.2", "1.1.1.3"]
  name = "tf-test"
}