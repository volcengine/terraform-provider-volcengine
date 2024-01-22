resource "volcengine_transit_router" "foo" {
  transit_router_name = "acc-test-tf"
  description         = "acc-test-tf"
}

resource "volcengine_transit_router_grant_rule" "foo" {
  grant_account_id = "2000xxxxx"
  description = "acc-test-tf"
  transit_router_id = volcengine_transit_router.foo.id
}
