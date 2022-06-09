resource "vestack_scaling_group" "foo" {
  scaling_group_name = "tf-test"
  subnet_ids = ["subnet-12bhi1j0k4buo17q7y2aet40a","subnet-12azazcacdjpc17q7y2588d57"]
  multi_az_policy = "PRIORITY"
  desire_instance_number = 0
  min_instance_number = 0
  max_instance_number = 10
  instance_terminate_policy = "OldestInstance"
  default_cooldown = 20
}