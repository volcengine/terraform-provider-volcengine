resource "volcengine_network_acl" "foo" {
  vpc_id           = "vpc-2d6jskar243k058ozfdae13ne"
  network_acl_name = "tf-test-acl"

  ingress_acl_entries {
    network_acl_entry_name = "ingress1"
    policy                 = "accept"
    protocol               = "all"
    source_cidr_ip         = "192.168.0.0/24"
  }

  egress_acl_entries {
    network_acl_entry_name = "egress2"
    policy                 = "accept"
    protocol               = "all"
    destination_cidr_ip    = "192.168.0.0/16"
  }

  ingress_acl_entries {
    network_acl_entry_name = "ingress3"
    policy                 = "accept"
    protocol               = "tcp"
    port                   = "80/80"
    source_cidr_ip         = "192.168.0.0/24"
  }

  project_name = "default"
}
