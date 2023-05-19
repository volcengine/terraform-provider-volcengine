resource "volcengine_rds_mysql_instance" "foo" {
  db_engine_version = "MySQL_5_7"
  node_spec = "rds.mysql.1c2g"
  primary_zone_id = "cn-guilin-a"
  secondary_zone_id = "cn-guilin-b"
  storage_space = 80
  subnet_id = "subnet-2d72yi377stts58ozfdrlk9f6"
  instance_name = "tf-test"
  lower_case_table_names = "1"

  charge_info {
    charge_type = "PostPaid"
  }

  allow_list_ids = ["acl-2dd8f8317e4d4159b21630d13ae2e6ec", "acl-2eaa2a053b2a4a58b988e38ae975e81c"]

  parameters {
    parameter_name = "auto_increment_increment"
    parameter_value = "2"
  }
  parameters {
    parameter_name = "auto_increment_offset"
    parameter_value = "4"
  }
}

resource "volcengine_rds_mysql_instance_readonly_node" "readonly" {
  instance_id = volcengine_rds_mysql_instance.foo.id
  node_spec = "rds.mysql.2c4g"
  zone_id = "cn-guilin-a"
}