resource "volcengine_transit_router_bandwidth_package" "foo" {
  transit_router_bandwidth_package_name = "acc-tf-test"
  description = "acc-test"
  bandwidth = 2
  period = 1
  renew_type = "Manual"
  renew_period = 1
  remain_renew_times = -1
}

resource "volcengine_transit_router" "foo" {
  transit_router_name = "acc-test-tf"
  description         = "acc-test-tf"
}

resource "volcengine_transit_router_peer_attachment" "foo" {
  transit_router_id = volcengine_transit_router.foo.id
  transit_router_attachment_name = "acc-test-tf"
  description = "tf-test"
  peer_transit_router_id = "tr-xxx"
  peer_transit_router_region_id = "cn-xx"
  transit_router_bandwidth_package_id = volcengine_transit_router_bandwidth_package.foo.id
  bandwidth = 2
  tags {
    key = "k1"
    value = "v1"
  }
}
