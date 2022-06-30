resource "volcengine_listener" "foo" {
  load_balancer_id = "clb-273ylkl0a3i807fap8t4unbsq"
  listener_name = "Demo-HTTP-90"
  protocol = "HTTP"
  port = 90
  server_group_id = "rsp-273yv0kir1vk07fap8tt9jtwg"
  health_check {
    enabled = "on"
    interval = 10
    timeout = 3
    healthy_threshold = 5
    un_healthy_threshold = 2
    domain = "github.com"
    http_code = "http_2xx"
    method = "GET"
    uri = "/"
  }
  enabled = "on"
}