resource "volcengine_nlb_listener" "foo" {
  load_balancer_id   = "nlb-2bznov724ct8g2dx0eg5dofhl"
  protocol           = "TCP"
  port               = 80
  server_group_id    = "rsp-3rezrz8q5s64g5zsk2ijjvhk4"
  listener_name      = "nlb-lsn-test-tf-1"
  description        = "nlb lsn test by tf"
  connection_timeout = 900
  enabled            = true
  tags {
    key   = "k3"
    value = "v4"
  }
}
