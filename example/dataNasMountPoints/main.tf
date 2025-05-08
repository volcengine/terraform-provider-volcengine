data "volcengine_nas_zones" "foo" {

}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-project1"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-test-2"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_nas_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_nas_permission_group" "foo" {
  permission_group_name = "acc-test"
  description           = "acctest"
  permission_rules {
    cidr_ip  = "*"
    rw_mode  = "RW"
    use_mode = "All_squash"
  }
  permission_rules {
    cidr_ip  = "192.168.0.0"
    rw_mode  = "RO"
    use_mode = "All_squash"
  }
}

resource "volcengine_nas_file_system" "foo" {
  file_system_name = "acc-test-fs"
  description      = "acc-test"
  zone_id          = data.volcengine_nas_zones.foo.zones[0].id
  capacity         = 103
  project_name     = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_nas_mount_point" "foo" {
  file_system_id      = volcengine_nas_file_system.foo.id
  mount_point_name    = "acc-test"
  permission_group_id = volcengine_nas_permission_group.foo.id
  subnet_id           = volcengine_subnet.foo.id
}

data "volcengine_nas_mount_points" "foo" {
  file_system_id = volcengine_nas_file_system.foo.id
  mount_point_id = volcengine_nas_mount_point.foo.mount_point_id
}