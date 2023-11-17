resource "volcengine_alb_customized_cfg" "foo" {
  customized_cfg_name = "acc-test-cfg1"
  description = "This is a test modify"
  customized_cfg_content = "proxy_connect_timeout 4s;proxy_request_buffering on;"
  project_name = "default"
}