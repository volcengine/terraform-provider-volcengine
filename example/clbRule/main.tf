resource "volcengine_clb_rule" "foo" {
  listener_id = "lsn-273ywvnmiu70g7fap8u2xzg9d"
  server_group_id = "rsp-273yxuqfova4g7fap8tyemn6t"
  domain = "test-volc123.com"
  url = "/yyyy"
}