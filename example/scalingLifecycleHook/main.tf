resource "vestack_scaling_lifecycle_hook" "foo" {
  scaling_group_id = "scg-ybru8pazhgl8j1di4tyd"
  lifecycle_hook_name = "tf-test"
  lifecycle_hook_timeout = 30
  lifecycle_hook_type = "SCALE_IN"
  lifecycle_hook_policy = "CONTINUE"
}