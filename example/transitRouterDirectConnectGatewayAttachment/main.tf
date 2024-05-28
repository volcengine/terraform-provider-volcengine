resource "volcengine_transit_router" "foo" {
  transit_router_name = "acc-test-tf-acc"
  description         = "acc-test-tf-acc"
}

resource "volcengine_direct_connect_gateway" "foo"{
  direct_connect_gateway_name="acc-test-gateway-acc"
  description="acc-test-acc"
  tags{
    key="k1"
    value="v1"
  }
}

resource "volcengine_transit_router_direct_connect_gateway_attachment" "foo" {
  description = "acc-test-tf"
  transit_router_attachment_name = "acc-test-tf"
  transit_router_id = volcengine_transit_router.foo.id
  direct_connect_gateway_id = volcengine_direct_connect_gateway.foo.id
  tags {
    key = "k1"
    value = "v1"
  }
}
