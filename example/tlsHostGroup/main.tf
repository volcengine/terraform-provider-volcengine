resource "volcengine_tls_host_group" "foo" {
  host_group_name = "tfgroup"
  host_group_type = "Label"
  host_identifier = "tf-controller"
  auto_update     = false
  service_logging = false
}