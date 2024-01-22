resource "volcengine_transit_router" "foo" {
  transit_router_name = "acc-test-tr"
  description         = "acc-test"
}

resource "volcengine_transit_router_shared_transit_router_state" "foo" {
  transit_router_id = volcengine_transit_router.foo.id
  action = "Reject"
}