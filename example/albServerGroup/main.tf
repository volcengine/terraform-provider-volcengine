resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_alb_server_group" "foo" {
  vpc_id = volcengine_vpc.foo.id
  server_group_name = "acc-test-server-group"
  description = "acc-test"
  server_group_type = "instance"
  scheduler = "wlc"
  protocol = "HTTP"
  ip_address_type = "IPv4"
  project_name = "default"
  health_check {
    enabled = "on"
    interval = 3
    timeout = 3
    method = "GET"
    domain = "www.test.com"
    uri = "/health"
    http_code = "http_2xx,http_3xx"
    protocol = "HTTP"
    port = 80
    http_version = "HTTP1.1"
  }
  sticky_session_config {
    sticky_session_enabled = "on"
    sticky_session_type = "insert"
    cookie_timeout = 1100
  }
  tags {
    key = "key1"
    value = "value2"
  }
}
