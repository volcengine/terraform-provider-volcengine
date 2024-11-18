provider "volcengine" {
  access_key = "access_key_1"
  secret_key = "secret_key_1"
  region     = "region_1"
}

provider "volcengine" {
  access_key = "access_key_2"
  secret_key = "secret_key_2"
  region     = "region_2"
  alias      = "second_account"
}

resource "volcengine_transit_router" "foo" {
  transit_router_name = "acc-test-tr"
  description         = "acc-test"
}

resource "volcengine_transit_router_grant_rule" "foo" {
  grant_account_id  = "2000xxxxx"
  description       = "acc-test-tf"
  transit_router_id = volcengine_transit_router.foo.id
}

resource "volcengine_transit_router_shared_transit_router_state" "foo" {
  transit_router_id = volcengine_transit_router.foo.id
  action            = "Accept"
  provider          = volcengine.second_account
}
