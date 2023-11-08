resource "volcengine_alb_listener" "foo" {
  load_balancer_id="alb-1iidd17v3klj474adhfrunyz9"
  listener_name="acc-test-listener-1"
  protocol="HTTPS"
  port=6666
  enabled="on"
  certificate_id = "cert-1iidd2pahdyio74adhfr9ajwg"
  ca_certificate_id = "cert-1iidd2r9ii0hs74adhfeodxo1"
  server_group_id="rsp-1g72w74y4umf42zbhq4k4hnln"
  enable_http2="on"
  enable_quic="off"
  acl_status="on"
  acl_type="white"
  acl_ids=["acl-1g72w6z11ighs2zbhq4v3rvh4"]
  description="acc test listener"
}

resource "volcengine_alb_listener_domain_extension" "foo" {
  listener_id = volcengine_alb_listener.foo.id
  domain = "test-modify.com"
  certificate_id = "cert-1iidd2pahdyio74adhfr9ajwg"
}