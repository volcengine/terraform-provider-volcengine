resource "volcengine_nlb_server_group" "foo" {
  server_group_name             = "nlb-sg"
  project_name                  = "default"
  vpc_id                        = "vpc-2d64s88ovqb5s58ozfe3uj5mx"
  type                          = "instance"
  protocol                      = "TCP"
  scheduler                     = "wrr"
  ip_address_version            = "ipv4"
  description                   = "nlb sg test by tf"
  any_port_enabled              = false
  connection_drain_enabled      = true
  connection_drain_timeout      = 60
  preserve_client_ip_enabled    = true
  session_persistence_enabled   = true
  session_persistence_timeout   = 900
  proxy_protocol_type           = "off"
  bypass_security_group_enabled = false
  timestamp_remove_enabled      = true
  health_check {
    enabled             = true
    type                = "TCP"
    port                = 1
    method              = "GET"
    uri                 = "/"
    domain              = "volcengine.com"
    http_code           = "http_3xx"
    interval            = 13
    timeout             = 13
    healthy_threshold   = 4
    unhealthy_threshold = 4
    udp_request         = ""
    udp_expect          = ""
  }
  servers {
    instance_id = "i-yehgzql1c0r9cxybukct"
    type        = "ecs"
    ip          = "your-ip"
    port        = 60
    weight      = 100
    description = "ecs server"
    zone_id     = "cn-guilin-a"
  }

  tags {
    key   = "km"
    value = "vm"
  }
  tags {
      key   = "kmt"
      value = "vm"
    }
}
