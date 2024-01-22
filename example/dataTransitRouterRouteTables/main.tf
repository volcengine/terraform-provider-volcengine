resource "volcengine_transit_router" "foo" {
  transit_router_name = "test-tf-acc"
  description         = "test-tf-acc"
}

resource "volcengine_transit_router_route_table" "foo" {
  description = "tf-test-acc-description"
  transit_router_route_table_name = "tf-table-test-acc"
  transit_router_id = volcengine_transit_router.foo.id
}


data "volcengine_transit_router_route_tables" "default" {
  transit_router_id = volcengine_transit_router.foo.id
  ids = [volcengine_transit_router_route_table.foo.transit_router_route_table_id]
  //transit_router_route_table_type = ""
}
