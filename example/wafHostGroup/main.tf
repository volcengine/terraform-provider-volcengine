resource "volcengine_waf_host_group" "foo" {
  description = "tf-test"
  host_list = ["www.tf-test.com"]
  name = "tf-test"
}