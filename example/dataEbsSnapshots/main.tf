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

resource "volcengine_ebs_snapshot" "foo" {
  volume_id      = volcengine_volume.foo.id
  snapshot_name  = "acc-test-snapshot"
  description    = "acc-test"
  retention_days = 3
  project_name   = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  count = 2
}

data "volcengine_ebs_snapshots" "foo" {
  ids = volcengine_ebs_snapshot.foo[*].id
}