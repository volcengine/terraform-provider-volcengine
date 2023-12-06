resource "volcengine_bandwidth_package" "foo" {
    bandwidth_package_name = "tf-test"
    billing_type = "PostPaidByBandwidth"
    //billing_type = "PrePaid"
    isp = "BGP"
    description = "tftest-description"
    bandwidth = 10
    protocol = "IPv4"
    //period = 1
    security_protection_types = ["AntiDDoS_Enhanced"]
    tags {
        key = "tftest"
        value = "tftest"
    }
}