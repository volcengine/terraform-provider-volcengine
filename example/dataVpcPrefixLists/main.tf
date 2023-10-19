resource "volcengine_vpc_prefix_list" "foo" {
  prefix_list_name = "acc-test-prefix"
  max_entries = 3
  description = "acc test description"
  ip_version = "IPv4"
  prefix_list_entries {
    cidr = "192.168.4.0/28"
    description = "acc-test-1"
  }
  prefix_list_entries {
    cidr = "192.168.5.0/28"
    description = "acc-test-2"
  }
  tags {
    key = "tf-key1"
    value = "tf-value1"
  }
}

data "volcengine_vpc_prefix_lists" "foo" {
  ids = [volcengine_vpc_prefix_list.foo.id]
}