resource "vestack_vpn_gateway_route" "foo" {
  vpn_gateway_id = "vgw-2c012ea9fm5mo2dx0efxg46qi"
  destination_cidr_block = "192.168.0.0/20"
  next_hop_id = "vgc-2d5ww3ww2lwcg58ozfe61ppc3"
}