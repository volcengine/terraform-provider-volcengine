resource "vestack_scaling_lifecycle_hook" "foo" {
  scaling_group_id = "scg-ybqhkrmdekgh9zk8sdi7"
  lifecycle_hook_name = "tf-test1"
  lifecycle_hook_timeout = 200
  lifecycle_hook_type = "SCALE_OUT"
  lifecycle_hook_policy = "REJECT"
}