resource "volcengine_scaling_group" "foo" {
  scaling_group_name = "tf-test"
  subnet_ids = ["subnet-2ff1n75eyf08w59gp67qhnhqm"]
  multi_az_policy = "BALANCE"
  desire_instance_number = 0
  min_instance_number = 0
  max_instance_number = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown = 10
}