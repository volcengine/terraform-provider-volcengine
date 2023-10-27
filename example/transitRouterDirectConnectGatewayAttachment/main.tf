resource "volcengine_transit_router_direct_connect_gateway_attachment" "foo" {
  transit_router_id = "tr-2bzy39x27qtxc2dx0eg5qaj05"
  direct_connect_gateway_id = "dcg-3reaq6ymdzegw5zsk2igxzusb"
  description = "tf-test-modify"
  transit_router_attachment_name = "tf-test-modify"
}