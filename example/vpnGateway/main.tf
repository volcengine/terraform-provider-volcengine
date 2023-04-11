resource "volcengine_vpn_gateway" "foo" {
  vpc_id = "vpc-2fe19q1dn2g3k59gp68n7w3rr"
  subnet_id = "subnet-2fe19qp20f3sw59gp67w8om25"
  bandwidth = 20
  vpn_gateway_name = "tf-test"
  description = "tf-test"
  period = 2
  project_name = "yuwenhao"
}