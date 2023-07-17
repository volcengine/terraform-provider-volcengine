resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByBandwidth"
  bandwidth = 1
  isp = "ChinaUnicom"
  name = "tf-eip"
  description = "tf-test"
  project_name = "default"
}