resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_alb_server_group" "foo" {
  vpc_id = volcengine_vpc.foo.id
  server_group_name = "acc-test-server-group-${count.index}"
  description = "acc-test"
  server_group_type = "instance"
  scheduler = "sh"
  project_name = "default"
  health_check {
    enabled = "on"
    interval = 3
    timeout = 3
    method = "GET"
  }
  sticky_session_config {
    sticky_session_enabled = "on"
    sticky_session_type = "insert"
    cookie_timeout = "1100"
  }
  count = 3
}

data "volcengine_alb_server_groups" "foo"{
  ids = volcengine_alb_server_group.foo[*].id
}
