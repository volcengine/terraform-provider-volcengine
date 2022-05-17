resource "vestack_eip_address" "foo" {
  billing_type = "PostPaidByBandwidth"
  bandwidth = 1
  isp = "BGP"
  name = "tf-test"
  description = "tf-test"
}