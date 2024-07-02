resource "volcengine_transit_router" "foo" {
  transit_router_name = "acc-test-tr"
  description         = "acc-test"
  asn                 = 4294967294
  project_name        = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}