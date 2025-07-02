data "volcengine_waf_cc_rules" "foo" {
  cc_type = [1]
  host = "www.tf-test.com"
  rule_name = "tf"
  path_order_by = "ASC"
}