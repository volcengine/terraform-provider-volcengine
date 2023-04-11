resource "volcengine_network_interface" "foo" {
  subnet_id = "subnet-im67x70vxla88gbssz1hy1z2"
  security_group_ids = ["sg-im67wp9lx3i88gbssz3d22b2"]
  primary_ip_address = "192.168.0.253"
  network_interface_name = "tf-test-up"
  description = "tf-test-up"
  port_security_enabled = false
  project_name = "default"
}