resource "volcengine_cloud_identity_group" "foo" {
  group_name   = "acc-test-group"
  display_name = "tf-test-group"
  join_type    = "Manual"
  description  = "tf"
}
