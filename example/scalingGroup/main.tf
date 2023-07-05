resource "volcengine_scaling_group" "foo" {
  scaling_group_name = "test-tf"
  subnet_ids = ["subnet-2fe79j7c8o5c059gp68ksxr93"]
  multi_az_policy = "BALANCE"
  desire_instance_number = 0
  min_instance_number = 0
  max_instance_number = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown = 10
  project_name = "default"
  tags {
    key = "xx"
    value = "xxaaa"
  }
}