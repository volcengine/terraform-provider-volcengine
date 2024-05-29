resource "volcengine_transit_router" "foo" {
  transit_router_name = "test-tf-acc"
  description         = "test-tf-acc"
}

resource "volcengine_transit_router_route_table" "foo" {
  description = "tf-test-acc-description"
  transit_router_route_table_name = "tf-table-test-acc"
  transit_router_id = volcengine_transit_router.foo.id
  tags {
    key = "k1"
    value = "v1"
  }
}
