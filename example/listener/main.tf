resource "volcengine_listener" "foo" {
  load_balancer_id = "clb-274xltt3rfmyo7fap8sv1jq39"
  listener_name = "Demo-HTTP-90"
  protocol = "HTTP"
  port = 90
  server_group_id = "rsp-274xltv2sjoxs7fap8tlv3q3s"
  health_check {
    enabled = "on"
    interval = 10
    timeout = 3
    healthy_threshold = 5
    un_healthy_threshold = 2
    domain = "volcengine.com"
    http_code = "http_2xx"
    method = "GET"
    uri = "/"
  }
  enabled = "on"
}

resource "volcengine_listener" "bar" {
  load_balancer_id = "clb-274xltt3rfmyo7fap8sv1jq39"
  listener_name = "Demo-HTTP-91"
  protocol = "HTTP"
  port = 91
  server_group_id = "rsp-274xltv2sjoxs7fap8tlv3q3s"
  health_check {
    enabled = "on"
    interval = 10
    timeout = 3
    healthy_threshold = 5
    un_healthy_threshold = 2
    domain = "volcengine.com"
    http_code = "http_2xx"
    method = "GET"
    uri = "/"
  }
  enabled = "on"
}

resource "volcengine_listener" "demo" {
  load_balancer_id = "clb-274xltt3rfmyo7fap8sv1jq39"
  protocol = "TCP"
  port = 92
  server_group_id = "rsp-274xltv2sjoxs7fap8tlv3q3s"
  health_check {
    enabled = "on"
    interval = 10
    timeout = 3
    healthy_threshold = 5
    un_healthy_threshold = 2
  }
  enabled = "on"
  established_timeout = 10
}