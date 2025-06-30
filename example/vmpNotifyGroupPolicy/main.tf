resource "volcengine_vmp_notify_group_policy" "foo" {
  name = "acc-test-1"
  description = "acc-test-1"
  levels {
    level = "P2"
    group_by = ["__rule__"]
    group_wait = "35"
    group_interval = "30"
    repeat_interval = "30"
  }
  levels {
    level = "P0"
    group_by = ["__rule__"]
    group_wait = "30"
    group_interval = "30"
    repeat_interval = "30"
  }
  levels {
    level = "P1"
    group_by = ["__rule__"]
    group_wait = "40"
    group_interval = "45"
    repeat_interval = "30"
  }
}
