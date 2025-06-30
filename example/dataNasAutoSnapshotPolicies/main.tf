resource "volcengine_nas_auto_snapshot_policy" "foo" {
  auto_snapshot_policy_name = "acc-test-auto_snapshot_policy"
  repeat_weekdays           = "1,3,5,7"
  time_points               = "0,7,17"
  retention_days            = 20
}

data "volcengine_nas_auto_snapshot_policies" "foo" {
  auto_snapshot_policy_id = volcengine_nas_auto_snapshot_policy.foo.id
}