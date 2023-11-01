resource "volcengine_direct_connect_bgp_peer" "foo" {
  virtual_interface_id = "dcv-62vi13v131tsn3gd6il****"
  remote_asn           = 2000
  description          = "tf-test"
}