resource "volcengine_privatelink_vpc_endpoint_service" "foo" {
  resources {
    resource_id   = "clb-2bzxccdjo9uyo2dx0eg0orzla"
    resource_type = "CLB"
  }
  description         = "tftest"
  auto_accept_enabled = true
}