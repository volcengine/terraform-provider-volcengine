data "volcengine_nas_zones" "foo" {

}

resource "volcengine_nas_file_system" "foo" {
  file_system_name = "acc-test-fs-${count.index}"
  description = "acc-test"
  zone_id = data.volcengine_nas_zones.foo.zones[0].id
  capacity = 103
  project_name = "default"
  tags {
    key = "k1"
    value = "v1"
  }
  count = 3
}

data "volcengine_nas_file_systems" "foo"{
  ids = volcengine_nas_file_system.foo[*].id
}