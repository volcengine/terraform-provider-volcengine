resource "volcengine_ebs_auto_snapshot_policy" "foo" {
  auto_snapshot_policy_name = "acc-test-auto-snapshot-policy"
  time_points               = [1, 5, 9]
  retention_days            = -1
  repeat_weekdays           = [2, 6]
#  repeat_days               = 5
  project_name              = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
