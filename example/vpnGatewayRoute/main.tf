resource "vestack_vpn_gateway_route" "foo" {
  vpn_gateway_id = "vgw-2d689v3lxs0zk58ozfebct3fc"
  destination_cidr_block = "192.168.0.0/20"
  next_hop_id = "vgc-2bytcbwy2txj42dx0efb93tag"
}