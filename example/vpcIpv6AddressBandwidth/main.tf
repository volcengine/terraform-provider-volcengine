data "volcengine_ecs_instances" "dataEcs" {
  ids = ["i-ycal1mtpucl8j0hjiihy"]
}

data "volcengine_vpc_ipv6_addresses" "dataIpv6" {
  associated_instance_id = data.volcengine_ecs_instances.dataEcs.instances.0.instance_id
}

resource "volcengine_vpc_ipv6_address_bandwidth" "foo" {
  ipv6_address = data.volcengine_vpc_ipv6_addresses.dataIpv6.ipv6_addresses.0.ipv6_address
  billing_type = "PostPaidByBandwidth"
  bandwidth = 5
}