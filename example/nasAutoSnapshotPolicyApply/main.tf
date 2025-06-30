data "volcengine_nas_zones" "foo" {

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

resource "volcengine_nas_auto_snapshot_policy" "foo" {
  auto_snapshot_policy_name = "acc-test-auto_snapshot_policy"
  repeat_weekdays           = "1,3,5,7"
  time_points               = "0,7,17"
  retention_days            = 20
}

resource "volcengine_nas_auto_snapshot_policy_apply" "foo" {
  file_system_id          = volcengine_nas_file_system.foo.id
  auto_snapshot_policy_id = volcengine_nas_auto_snapshot_policy.foo.id
}