resource "volcengine_waf_instance_ctl" "foo" {
  allow_enable = 0
  block_enable = 1
  project_name = "default"
}
