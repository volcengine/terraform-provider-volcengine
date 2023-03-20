resource "volcengine_network_acl" "foo" {
  vpc_id = "vpc-12bk4qjc69reo17q7y36shv6z"
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
  project_name = "default"
}
