resource "volcengine_network_interface" "foo" {
  subnet_id = "subnet-2744ht7fhjthc7fap8tm10eqg"
  security_group_ids = ["sg-2744hspo7jbpc7fap8t7lef1p"]
  primary_ip_address = "192.168.0.253"
  network_interface_name = "tf-test-up"
  description = "tf-test-up"
  port_security_enabled = false
}