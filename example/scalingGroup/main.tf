resource "vestack_scaling_group" "foo" {
  scaling_group_name = "tf-test"
  subnet_ids = ["subnet-3relrtnt9piww5zsk2hpje6c5", "subnet-2d6g96vyl3r4058ozfdjdoprz"]
  multi_az_policy = "BALANCE"
  desire_instance_number = 0
  min_instance_number = 0
  max_instance_number = 10
  instance_terminate_policy = "OldestInstance"
  default_cooldown = 10
  db_instance_ids = [ "mysql-acce5580ae97"]
}