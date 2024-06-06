resource "volcengine_cloud_identity_group" "foo" {
  group_name   = "acc-test-group-${count.index}"
  display_name = "tf-test-group-${count.index}"
  join_type    = "Manual"
  description  = "tf"

  count = 2
}

data "volcengine_cloud_identity_groups" "foo" {
  group_name = "acc-test-group"
  join_type  = "Manual"
}
