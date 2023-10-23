resource "volcengine_transit_router_route_entry" "foo" {
  transit_router_route_table_id = "tr-rtb-12b7qd3fmzf2817q7y2jkbd55"
  destination_cidr_block = "192.168.0.0/24"
  transit_router_route_entry_next_hop_type = "BlackHole"
  //transit_router_route_entry_next_hop_id = ""
  description = "tf test 23"
  transit_router_route_entry_name = "tf-entry-23"
}