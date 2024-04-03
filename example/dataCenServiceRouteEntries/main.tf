resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
  count      = 3
}

resource "volcengine_cen" "foo" {
  cen_name     = "acc-test-cen"
  description  = "acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_cen_attach_instance" "foo" {
  cen_id             = volcengine_cen.foo.id
  instance_id        = volcengine_vpc.foo[count.index].id
  instance_region_id = "cn-beijing"
  instance_type      = "VPC"
  count              = 3
}

resource "volcengine_cen_service_route_entry" "foo" {
  cen_id                 = volcengine_cen.foo.id
  destination_cidr_block = "100.64.0.0/11"
  service_region_id      = "cn-beijing"
  service_vpc_id         = volcengine_cen_attach_instance.foo[0].instance_id
  description            = "acc-test"
  publish_mode           = "Custom"
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type      = "VPC"
    instance_id        = volcengine_cen_attach_instance.foo[1].instance_id
  }
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type      = "VPC"
    instance_id        = volcengine_cen_attach_instance.foo[2].instance_id
  }
}

data "volcengine_cen_service_route_entries" "foo" {
  cen_id                 = volcengine_cen.foo.id
  destination_cidr_block = volcengine_cen_service_route_entry.foo.destination_cidr_block
}
