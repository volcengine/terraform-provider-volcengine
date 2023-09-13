data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet_new"
  description = "tfdesc"
  cidr_block = "172.16.0.0/24"
  zone_id = data.volcengine_zones.foo.zones[0].id
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_escloud_instance" "foo" {
  instance_configuration {
    version            = "V6_7"
    zone_number        = 1
    enable_https       = true
    admin_user_name    = "admin"
    admin_password     = "Password@@"
    charge_type        = "PostPaid"
    configuration_code = "es.standard"
    enable_pure_master = true
    instance_name      = "acc-test-0"
    node_specs_assigns {
      type               = "Master"
      number             = 3
      resource_spec_name = "es.x4.medium"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 100
    }
    node_specs_assigns {
      type               = "Hot"
      number             = 2
      resource_spec_name = "es.x4.large"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 100
    }
    node_specs_assigns {
      type               = "Kibana"
      number             = 1
      resource_spec_name = "kibana.x2.small"
    }
    subnet_id = volcengine_subnet.foo.id
    project_name = "default"
    force_restart_after_scale = false
  }
}