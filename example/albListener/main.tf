data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_alb" "foo" {
  address_ip_version = "IPv4"
  type               = "private"
  load_balancer_name = "acc-test-alb-private"
  description        = "acc-test"
  subnet_ids         = [volcengine_subnet.foo.id]
  project_name       = "default"
  delete_protection  = "off"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_alb_server_group" "foo" {
  vpc_id            = volcengine_vpc.foo.id
  server_group_name = "acc-test-server-group"
  description       = "acc-test"
  server_group_type = "instance"
  scheduler         = "wlc"
  project_name      = "default"
  health_check {
    enabled  = "on"
    interval = 3
    timeout  = 3
    method   = "GET"
  }
  sticky_session_config {
    sticky_session_enabled = "on"
    sticky_session_type    = "insert"
    cookie_timeout         = "1100"
  }
}

resource "volcengine_alb_certificate" "foo" {
  description = "tf-test"
  public_key  = "public key"
  private_key = "private key"
}

resource "volcengine_alb_listener" "foo" {
  load_balancer_id   = volcengine_alb.foo.id
  listener_name      = "acc-test-listener"
  protocol           = "HTTPS"
  port               = 6666
  enabled            = "off"
  certificate_source = "alb"
  #  cert_center_certificate_id = "cert-***"
  certificate_id  = volcengine_alb_certificate.foo.id
  server_group_id = volcengine_alb_server_group.foo.id
  description     = "acc test listener"
}
