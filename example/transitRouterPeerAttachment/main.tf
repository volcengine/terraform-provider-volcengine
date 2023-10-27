resource "volcengine_transit_router_peer_attachment" "foo" {
  transit_router_id = "tr-12bbdsa6ode6817q7y1f5****"
  transit_router_attachment_name = "tf-test-tra"
  description = "tf-test"
  peer_transit_router_id = "tr-3jgsfiktn0feo3pncmfb5****"
  peer_transit_router_region_id = "cn-beijing"
  transit_router_bandwidth_package_id = "tbp-cd-2felfww0i6pkw59gp68bq****"
  bandwidth = 2
}