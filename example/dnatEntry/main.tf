resource "volcengine_dnat_entry" "foo" {
  nat_gateway_id = "ngw-imw3aej7e96o8gbssxkfbybv"
  external_ip = "10.249.186.68"
  external_port = "23"
  internal_ip = "193.168.1.1"
  internal_port = "24"
  protocol = "tcp"
  dnat_entry_name = "terraform-test2"
}
