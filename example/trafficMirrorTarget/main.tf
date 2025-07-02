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
  security_group_name = "acc-test-sg"
  vpc_id              = volcengine_vpc.foo.id
}

data "volcengine_images" "foo" {
  os_type          = "Linux"
  visibility       = "public"
  instance_type_id = "ecs.g3il.large"
}

resource "volcengine_ecs_instance" "foo" {
  instance_name        = "acc-test-ecs"
  description          = "acc-test"
  host_name            = "tf-acc-test"
  image_id             = data.volcengine_images.foo.images[0].image_id
  instance_type        = "ecs.g3il.large"
  password             = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type   = "ESSD_PL0"
  system_volume_size   = 40
  subnet_id            = volcengine_subnet.foo.id
  security_group_ids   = [volcengine_security_group.foo.id]
  project_name         = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_network_interface" "foo" {
  network_interface_name = "acc-test-eni"
  description            = "acc-test"
  subnet_id              = volcengine_subnet.foo.id
  security_group_ids     = [volcengine_security_group.foo.id]
  primary_ip_address     = "172.16.0.253"
  port_security_enabled  = false
  private_ip_address     = ["172.16.0.2"]
  project_name           = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_network_interface_attach" "foo" {
  network_interface_id = volcengine_network_interface.foo.id
  instance_id          = volcengine_ecs_instance.foo.id
}

resource "volcengine_traffic_mirror_target" "foo" {
  instance_type              = "NetworkInterface"
  instance_id                = volcengine_network_interface.foo.id
  traffic_mirror_target_name = "acc-test-traffic-mirror-target"
  description                = "acc-test"
  project_name               = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  depends_on = [volcengine_network_interface_attach.foo]
}
