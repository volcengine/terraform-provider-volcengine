resource "volcengine_transit_router_bandwidth_package" "foo" {
  transit_router_bandwidth_package_name = "acc-tf-test"
  description                           = "acc-test"
  local_geographic_region_set_id        = "China"
  peer_geographic_region_set_id         = "China"
  bandwidth                             = 2
  period                                = 1
  renew_type                            = "Manual"
  project_name                          = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}