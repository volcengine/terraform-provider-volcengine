resource "vestack_customer_gateway" "foo" {
  ip_address = "192.0.1.2"
  customer_gateway_name = "tf-test"
  description = "tf-test"
}