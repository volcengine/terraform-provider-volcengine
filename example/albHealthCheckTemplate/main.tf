resource "volcengine_alb_health_check_template" "foo" {
  health_check_template_name="acc-test-template-1"
  description="acc-test3"
  health_check_interval=8
  health_check_timeout=11
  healthy_threshold=2
  unhealthy_threshold=3
  health_check_method="HEAD"
  health_check_domain="test.com"
  health_check_uri="/"
  health_check_http_code="http_2xx"
  health_check_protocol="HTTP"
  health_check_http_version = "HTTP1.1"
  tags {
      key   = "key1"
      value = "value2"
  }
}