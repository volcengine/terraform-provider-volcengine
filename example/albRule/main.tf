# Basic edition
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

# Standard edition
resource "volcengine_alb_rule" "example" {
  listener_id = "lsn-bddjp5fcof0g8dv40naga1yd"
  rule_action = ""
  description = "standard edition alb rule"
  url = ""
  priority    = 1
  # Matching condition: Host + Path
  rule_conditions {
    type = "Host"
    host_config {
      values = ["www.example.com"]
    }
  }
  rule_conditions {
    type = "Path"
    path_config {
      values = ["/app/*"]
    }
  }
  rule_actions {
    type = "ForwardGroup"
    forward_group_config {
        server_group_tuples {
            server_group_id = "rsp-bdd1lpcbvv288dv40ov1sye0"
            weight          = 50
        }
        server_group_sticky_session {
            enabled = "off"
        }
    }
 }
}