resource "volcengine_rds_ip_list" "foo" {
  instance_id = "mysql-42b38c769c4b"
  group_name = "qwecsa"
  ip_list = ["1.1.1.1"]
}