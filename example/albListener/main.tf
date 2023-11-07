resource "volcengine_alb_customized_cfg" "foo" {
  customized_cfg_name = "acc-test-cfg1"
  description = "This is a test modify"
  customized_cfg_content = "proxy_connect_timeout 4s;proxy_request_buffering on;"
  project_name = "default"
}

resource "volcengine_alb_listener" "foo" {
  load_balancer_id="alb-1iidd17v3klj474adhfrunyz9"
  listener_name="acc-test-listener-1"
  protocol="HTTP"
  port=6666
  enabled="on"
  # certificate_id = ""
  # ca_certificate_id = ""
  server_group_id="rsp-1g72w74y4umf42zbhq4k4hnln"
  enable_http2="off"
  enable_quic="off"
  acl_status="on"
  acl_type="white"
  acl_ids=["acl-1g72w6z11ighs2zbhq4v3rvh4"]
  description="acc test listener"
  customized_cfg_id = volcengine_alb_customized_cfg.foo.id
}