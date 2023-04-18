resource "volcengine_rds_mysql_instance_readonly_node" "foo" {
  instance_id = "mysql-b3fca7f571d6"
  node_spec = "rds.mysql.1c2g"
  zone_id = "cn-guilin-b"
}