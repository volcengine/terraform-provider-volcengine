resource "volcengine_alb_rule" "foo" {
  listener_id = "lsn-1iidd19u4oni874adhezjkyj3"
  domain = "www.test.com"
  url = "/test"
  rule_action = "Redirect"
  server_group_id = "rsp-1g72w74y4umf42zbhq4k4hnln"
  description = "test"
  traffic_limit_enabled = "off"
  traffic_limit_qps = 100
  rewrite_enabled = "off"
  redirect_config {
    redirect_domain = "www.testtest.com"
    redirect_uri = "/testtest"
    redirect_port = "555"
    redirect_http_code = "302"
    //redirect_http_protocol = ""
  }
  rewrite_config {
    rewrite_path = "/test"
  }
}