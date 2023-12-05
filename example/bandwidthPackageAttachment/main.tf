resource "volcengine_eip_address" "foo" {
    billing_type = "PostPaidByBandwidth"
    bandwidth = 1
    isp = "BGP"
    name = "acc-eip"
    description = "acc-test"
    project_name = "default"
}

resource "volcengine_bandwidth_package" "foo" {
    bandwidth_package_name = "acc-test"
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

resource "volcengine_bandwidth_package_attachment" "foo" {
    allocation_id = volcengine_eip_address.foo.id
    bandwidth_package_id = volcengine_bandwidth_package.foo.id
}