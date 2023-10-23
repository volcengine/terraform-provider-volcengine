resource "volcengine_transit_router_vpn_attachment" "foo" {
  transit_router_id = "tr-2d6frp10q687458ozfep4****"
  vpn_connection_id = "vgc-3reidwjf1t1c05zsk2hik****"
  zone_id = "cn-beijing-a"
  transit_router_attachment_name = "tf-test"
  description = "desc"
}