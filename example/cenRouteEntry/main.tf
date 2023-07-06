resource "volcengine_cen_route_entry" "foo" {
  cen_id = "cen-12ar8uclj68sg17q7y20v9gil"
  destination_cidr_block = "192.168.0.0/24"
  instance_type = "VPC"
  instance_region_id = "cn-beijing"
  instance_id = "vpc-im67wjcikxkw8gbssx8ufpj8"
}

resource "volcengine_cen_route_entry" "foo1" {
  cen_id = "cen-12ar8uclj68sg17q7y20v9gil"
  destination_cidr_block = "192.168.17.0/24"
  instance_type = "VPC"
  instance_region_id = "cn-beijing"
  instance_id = "vpc-im67wjcikxkw8gbssx8ufpj8"
}