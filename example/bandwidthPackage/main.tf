resource "volcengine_bandwidth_package" "foo" {
  bandwidth_package_name    = "acc-test-bp"
  billing_type              = "PostPaidByBandwidth"
  isp                       = "BGP"
  description               = "acc-test"
  bandwidth                 = 10
  protocol                  = "IPv4"
  security_protection_types = ["AntiDDoS_Enhanced"]
  tags {
    key   = "k1"
    value = "v1"
  }
}
