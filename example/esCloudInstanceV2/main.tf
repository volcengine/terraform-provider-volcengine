data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  description = "tfdesc"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_escloud_instance_v2" "foo" {
  instance_name       = "acc-test-escloud-instance"
  version             = "V7_10"
  zone_ids            = [data.volcengine_zones.foo.zones[0].id, data.volcengine_zones.foo.zones[1].id, data.volcengine_zones.foo.zones[2].id]
  subnet_id           = volcengine_subnet.foo.id
  enable_https        = false
  admin_password      = "Password@@123"
  charge_type         = "PostPaid"
  auto_renew          = false
  period              = 1
  configuration_code  = "es.standard"
  enable_pure_master  = true
  deletion_protection = false
  project_name        = "default"

  node_specs_assigns {
    type               = "Master"
    number             = 3
    resource_spec_name = "es.x2.medium"
    storage_spec_name  = "es.volume.essd.pl0"
    storage_size       = 20
  }
  node_specs_assigns {
    type               = "Hot"
    number             = 6
    resource_spec_name = "es.x2.medium"
    storage_spec_name  = "es.volume.essd.flexpl-standard"
    storage_size       = 500
    extra_performance {
      throughput = 65
    }
  }
  node_specs_assigns {
    type               = "Kibana"
    number             = 1
    resource_spec_name = "kibana.x2.small"
    storage_spec_name  = ""
    storage_size       = 0
  }

  network_specs {
    type      = "Elasticsearch"
    bandwidth = 1
    is_open   = true
    spec_name = "es.eip.bgp_fixed_bandwidth"
  }

  network_specs {
    type      = "Kibana"
    bandwidth = 1
    is_open   = true
    spec_name = "es.eip.bgp_fixed_bandwidth"
  }

  tags {
    key   = "k1"
    value = "v1"
  }

  #  maintenance_time = "02:00-08:00"
  #  maintenance_day = ["FRIDAY", "MONDAY"]
}
