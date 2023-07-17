resource "volcengine_cen_service_route_entry" "foo" {
  cen_id = "cen-12ar8uclj68sg17q7y20v9gil"
  destination_cidr_block = "100.64.0.0/11"
  service_region_id = "cn-beijing"
  service_vpc_id = "vpc-im67wjcikxkw8gbssx8ufpj8"
  description = "test-tf"
  publish_mode = "Custom"
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type = "VPC"
    instance_id = "vpc-2fepz36a5ra4g59gp67w197xo"
  }
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type = "VPC"
    instance_id = "vpc-im67wjcikxkw8gbssx8ufpj8"
  }
}

resource "volcengine_cen_service_route_entry" "foo1" {
  cen_id = "cen-12ar8uclj68sg17q7y20v9gil"
  destination_cidr_block = "100.64.0.0/10"
  service_region_id = "cn-beijing"
  service_vpc_id = "vpc-im67wjcikxkw8gbssx8ufpj8"
  description = "test-tf"
  publish_mode = "Custom"
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type = "VPC"
    instance_id = "vpc-2fepz36a5ra4g59gp67w197xo"
  }
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type = "VPC"
    instance_id = "vpc-im67wjcikxkw8gbssx8ufpj8"
  }
}

resource "volcengine_cen_service_route_entry" "foo2" {
  cen_id = "cen-12ar8uclj68sg17q7y20v9gil"
  destination_cidr_block = "100.64.0.0/12"
  service_region_id = "cn-beijing"
  service_vpc_id = "vpc-im67wjcikxkw8gbssx8ufpj8"
  description = "test-tf"
  publish_mode = "Custom"
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type = "VPC"
    instance_id = "vpc-2fepz36a5ra4g59gp67w197xo"
  }
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type = "VPC"
    instance_id = "vpc-im67wjcikxkw8gbssx8ufpj8"
  }
}