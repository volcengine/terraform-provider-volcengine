resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = "cn-beijing-a"
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_vepfs_file_system" "foo" {
  file_system_name = "acc-test-file-system"
  subnet_id        = volcengine_subnet.foo.id
  store_type       = "Advance_100"
  description      = "tf-test"
  capacity         = 12
  project          = "default"
  enable_restripe  = false
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_vepfs_mount_service" "foo" {
  mount_service_name = "acc-test-mount-service"
  subnet_id          = volcengine_subnet.foo.id
  node_type          = "ecs.g1ie.large"
  project            = "default"
}

resource "volcengine_vepfs_mount_service_attachment" "foo" {
  mount_service_id = volcengine_vepfs_mount_service.foo.id
  file_system_id   = volcengine_vepfs_file_system.foo.id
}
