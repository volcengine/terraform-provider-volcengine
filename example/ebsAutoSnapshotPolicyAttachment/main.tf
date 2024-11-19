data "volcengine_zones" "foo" {
}

resource "volcengine_volume" "foo" {
  volume_name        = "acc-test-volume"
  volume_type        = "ESSD_PL0"
  description        = "acc-test"
  kind               = "data"
  size               = 500
  zone_id            = data.volcengine_zones.foo.zones[0].id
  volume_charge_type = "PostPaid"
  project_name       = "default"
}

resource "volcengine_ebs_auto_snapshot_policy" "foo" {
  auto_snapshot_policy_name = "acc-test-auto-snapshot-policy"
  time_points               = [1, 5, 9]
  retention_days            = -1
  repeat_weekdays           = [2, 6]
  project_name              = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_ebs_auto_snapshot_policy_attachment" "foo" {
  auto_snapshot_policy_id = volcengine_ebs_auto_snapshot_policy.foo.id
  volume_id               = volcengine_volume.foo.id
}
