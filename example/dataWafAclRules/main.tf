data "volcengine_waf_acl_rules" "foo" {
  acl_type = "Block"
  action = ["observe"]
  defence_host = ["www.tf-test.com"]
  enable = [1]
  rule_name = "tf-test"
  time_order_by = "ASC"
  project_name = "default"
}