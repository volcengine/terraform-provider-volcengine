resource "volcengine_transit_router" "foo" {
  transit_router_name = "test-tf-acc"
  description         = "test-tf-acc"
}

data "volcengine_transit_routers" "default" {
  ids                 = [volcengine_transit_router.foo.id]
  transit_router_name = "test"
}