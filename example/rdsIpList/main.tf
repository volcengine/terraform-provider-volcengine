resource "volcengine_rds_ip_list" "foo" {
  instance_id = "mysql-0fdd3bab2e7c"
  group_name = "foo"
  ip_list = ["1.1.1.1", "2.2.2.2"]
}