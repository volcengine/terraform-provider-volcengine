resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByBandwidth"
  bandwidth = 1
  isp = "ChinaUnicom"
  name = "tf-project-1"
  description = "tf-test"
  project_name = "yuwenhao"
}