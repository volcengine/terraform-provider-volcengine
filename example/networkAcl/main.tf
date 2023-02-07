resource "volcengine_network_acl" "foo" {
  vpc_id = "vpc-ru0wv9alfoxsu3nuld85rpp"
  network_acl_name = "tf-test-acl"
  ingress_acl_entries {
    network_acl_entry_name = "ingress1"
    policy = "accept"
    protocol = "all"
    source_cidr_ip = "192.168.0.0/24"
  }
  egress_acl_entries {
    network_acl_entry_name = "egress2"
    policy = "accept"
    protocol = "all"
    destination_cidr_ip = "192.168.0.0/16"
  }
}
