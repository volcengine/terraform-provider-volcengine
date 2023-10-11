resource "volcengine_network_interface" "foo" {
  subnet_id = "subnet-2fe79j7c8o5c059gp68ksxr93"
  security_group_ids = ["sg-2fepz3c793g1s59gp67y21r34"]
  primary_ip_address = "192.168.5.253"
  network_interface_name = "tf-test-up"
  description = "tf-test-up"
  port_security_enabled = false
  project_name = "default"
  private_ip_address = ["192.168.5.2"]
  #  secondary_private_ip_address_count = 0
  #  ipv6_address_count = 2
  #  ipv6_addresses = ["2000:1000:12ff:ff01:2f1a:ca69:8110:34f5", "2000:1000:12ff:ff01:df81:a2d2:e568:1715"]
}