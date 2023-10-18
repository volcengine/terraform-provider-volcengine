resource "volcengine_vpc_prefix_list" "foo" {
  prefix_list_name = "acc-test-prefix"
  max_entries = 7
  description = "acc test description"
  ip_version = "IPv4"
  prefix_list_entries {
    cidr = "192.168.4.0/28"
    description = "acc-test-1"
  }
  prefix_list_entries {
    cidr = "192.168.9.0/28"
    description = "acc-test-4"
  }
  prefix_list_entries {
    cidr = "192.168.8.0/28"
    description = "acc-test-5"
  }
  tags {
    key = "tf-key1"
    value = "tf-value1"
  }
}