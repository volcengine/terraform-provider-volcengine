resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByTraffic"
}
data "volcengine_eip_addresses" "foo"{
  ids = ["${volcengine_eip_address.foo.id}"]
}