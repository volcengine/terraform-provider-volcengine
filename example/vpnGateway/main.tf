resource "vestack_vpn_gateway" "foo" {
  vpc_id = "vpc-2bysvq1xx543k2dx0eeulpeiv"
  subnet_id = "subnet-2d68bh74345q858ozfekrm8fj"
  bandwidth = 10
  vpn_gateway_name = "tf-test"
  description = "tf-test"
  period_unit = "Month"
  period = 1
  renew_type = "NoneRenew"
  renew_period = 1
  remain_renew_times = 3
}