resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByBandwidth"
  bandwidth = 1
  isp = "ChinaUnicom"
  name = "acc-eip"
  description = "acc-test"
  project_name = "default"
}