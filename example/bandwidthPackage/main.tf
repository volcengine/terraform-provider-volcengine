resource "volcengine_bandwidth_package" "foo" {
    bandwidth_package_name = "tf-test"
    billing_type = "PostPaidByBandwidth"
    isp = "BGP"
    description = "tftest-description"
    bandwidth = 10
    protocol = "IPv4"
    tags {
        key = "tftest"
        value = "tftest"
    }
}