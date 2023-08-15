data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-security-group"
  vpc_id              = volcengine_vpc.foo.id
}

data "volcengine_images" "foo" {
  os_type          = "Linux"
  visibility       = "public"
  instance_type_id = "ecs.g1.large"
}

resource "volcengine_scaling_group" "foo" {
  scaling_group_name        = "acc-test-scaling-group"
  subnet_ids                = [volcengine_subnet.foo.id]
  multi_az_policy           = "BALANCE"
  desire_instance_number    = 0
  min_instance_number       = 0
  max_instance_number       = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown          = 10
}

resource "volcengine_scaling_configuration" "foo" {
  scaling_configuration_name    = "tf-test"
  scaling_group_id              = volcengine_scaling_group.foo.id
  image_id                      = data.volcengine_images.foo.images[0].image_id
  instance_types                = ["ecs.g2i.large"]
  instance_name                 = "tf-test"
  instance_description          = ""
  host_name                     = ""
  password                      = ""
  key_pair_name                 = "tf-keypair"
  security_enhancement_strategy = "InActive"
  volumes {
    volume_type          = "ESSD_PL0"
    size                 = 20
    delete_with_instance = false
  }
  volumes {
    volume_type          = "ESSD_PL0"
    size                 = 50
    delete_with_instance = true
  }
  security_group_ids = [volcengine_security_group.foo.id]
  eip_bandwidth      = 10
  eip_isp            = "ChinaMobile"
  eip_billing_type   = "PostPaidByBandwidth"
  user_data          = "IyEvYmluL2Jhc2gKZWNobyAidGVzdCI="
  tags {
    key   = "tf-key1"
    value = "tf-value1"
  }
  tags {
    key   = "tf-key2"
    value = "tf-value2"
  }
  project_name   = "default"
  hpc_cluster_id = ""
  spot_strategy  = "NoSpot"
}

