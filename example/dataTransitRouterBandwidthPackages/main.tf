resource "volcengine_transit_router_bandwidth_package" "foo" {
  transit_router_bandwidth_package_name = "acc-tf-test"
  description = "acc-test"
  bandwidth = 2
  period = 1
  renew_type = "Manual"
}

data "volcengine_transit_router_bandwidth_packages" "foo" {
  ids = [volcengine_transit_router_bandwidth_package.foo.id]
}