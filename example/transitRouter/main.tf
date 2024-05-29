resource "volcengine_transit_router" "foo" {
  transit_router_name = "acc-test-tr"
  description         = "acc-test"
  project_name = "default"
  tags {
    key = "k1"
    value = "v1"
  }
}