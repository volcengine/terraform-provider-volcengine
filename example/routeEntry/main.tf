resource "volcengine_route_entry" "foo" {
  route_table_id = "vtb-2744hslq5b7r47fap8tjomgnj"
  destination_cidr_block = "0.0.0.0/2"
  next_hop_type = "NatGW"
  next_hop_id = "ngw-274gwbqe340zk7fap8spkzo7x"
  route_entry_name = "tf-test-up"
  description = "tf-test-up"
}