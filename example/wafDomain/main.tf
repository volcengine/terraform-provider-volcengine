resource "volcengine_waf_domain" "foo" {
  domain = "www.tf-test.com"
  access_mode = 10
  protocols = ["HTTP"]
  protocol_ports {
    http = [80]
  }
  enable_ipv6 = 0
  proxy_config = 1
  keep_alive_time_out = 100
  keep_alive_request = 200
  client_max_body_size = 1024
  lb_algorithm = "wlc"
  public_real_server = 0
  vpc_id = "vpc-2d6485y7p95og58ozfcvxxxxx"
  backend_groups {
    access_port = [80]
    backends {
      protocol = "HTTP"
      ip = "192.168.0.0"
      port = 80
      weight = 40
    }
    backends {
      protocol = "HTTP"
      ip = "192.168.1.0"
      port = 80
      weight = 60
    }
    name = "default"
  }
  client_ip_location = 0
  custom_header = ["x-top-1", "x-top-2"]
  proxy_connect_time_out = 10
  proxy_write_time_out = 120
  proxy_read_time_out = 200
  proxy_keep_alive = 101
  proxy_retry = 10
  proxy_keep_alive_time_out = 20
}