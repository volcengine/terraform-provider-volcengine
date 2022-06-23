resource "vestack_cen_bandwidth_package" "foo" {
  local_geographic_region_set_id = "China"
  peer_geographic_region_set_id = "China"
  bandwidth = 32
  cen_bandwidth_package_name = "tf-test"
  description = "tf-test1"
  billing_type = "PrePaid"
  period_unit = "Year"
  period = 1
}