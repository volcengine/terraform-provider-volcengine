resource "volcengine_waf_acl_rule" "foo" {
  action = "block"
  description = "tf-test"
  name = "tf-test-1"
  url = "/"
  ip_add_type = 3
  host_add_type = 3
  enable = 1
  acl_type = "Allow"
  host_list = ["www.tf-test.com"]
  ip_list = ["1.2.2.2", "1.2.3.30"]
  accurate_group {
    accurate_rules {
      http_obj = "request.uri"
      obj_type = 1
      opretar = 2
      property = 0
      value_string = "GET"
    }
    logic = 1
  }
  advanced = 1
  project_name = "default"
}