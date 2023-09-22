data "volcengine_nas_zones" "foo" {

}

resource "volcengine_nas_file_system" "foo" {
  file_system_name = "acc-test-fs"
  description = "acc-test"
  zone_id = data.volcengine_nas_zones.foo.zones[0].id
  capacity = 103
  project_name = "default"
  tags {
    key = "k1"
    value = "v1"
  }
}