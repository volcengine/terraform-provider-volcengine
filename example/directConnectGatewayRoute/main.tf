resource "volcengine_direct_connect_gateway_route" "foo" {
  direct_connect_gateway_id = "dcg-172frxs5utjb44d1w33op****"
  destination_cidr_block    = "192.168.40.0/24"
  next_hop_id               = "dcv-1729lrbfx7fuo4d1w34pk****"
}