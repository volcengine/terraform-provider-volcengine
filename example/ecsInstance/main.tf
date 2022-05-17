resource "vestack_vpc" "foo" {
  vpc_name = "tf-test-2"
  cidr_block = "172.16.0.0/16"
}

resource "vestack_subnet" "foo1" {
  subnet_name = "subnet-test-1"
  cidr_block = "172.16.1.0/24"
  zone_id = "cn-nantong-a"
  vpc_id = vestack_vpc.foo.id
}

resource "vestack_security_group" "foo1" {
  depends_on = [vestack_subnet.foo1]
  vpc_id = vestack_vpc.foo.id
}

resource "vestack_ecs_instance" "default" {
  zone_id = "cn-nantong-a"
  image_id = "image-cj79g0oghxjpvhifi3yu"
  instance_type = "ecs.g1.large"
  instance_name = "xym-tf-test-2"
  description = "xym-tf-test-desc-1"
  password = "93f0cb0614Aab12"
  instance_charge_type = "PostPaid"
  system_volume_type = "ESSD_PL0"
  system_volume_size = 60
  subnet_id = vestack_subnet.foo1.id
  security_group_ids = [vestack_security_group.foo1.id]
  data_volumes {
    volume_type = "ESSD_PL0"
    size = 100
    delete_with_instance = true
  }
#  secondary_network_interfaces {
#    subnet_id = vestack_subnet.foo1.id
#    security_group_ids = [vestack_security_group.foo1.id]
#  }
}
