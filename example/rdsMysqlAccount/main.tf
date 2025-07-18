data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
     vpc_name   = "acc-test-vpc"
     cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
     subnet_name = "acc-test-subnet"
     cidr_block = "172.16.0.0/24"
     zone_id = data.volcengine_zones.foo.zones[0].id
     vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_rds_mysql_instance" "foo" {
     instance_name = "acc-test-rds-mysql"
     db_engine_version = "MySQL_5_7"
     node_spec = "rds.mysql.1c2g"
     primary_zone_id = data.volcengine_zones.foo.zones[0].id
     secondary_zone_id = data.volcengine_zones.foo.zones[0].id
     storage_space = 80
     subnet_id = volcengine_subnet.foo.id
     lower_case_table_names = "1"
     charge_info {
          charge_type = "PostPaid"
     }
     parameters {
          parameter_name = "auto_increment_increment"
          parameter_value = "2"
     }
     parameters {
          parameter_name = "auto_increment_offset"
          parameter_value = "4"
     }
}

resource "volcengine_rds_mysql_database" "foo1" {
     db_name = "acc-test-db1"
     instance_id = volcengine_rds_mysql_instance.foo.id
     #instance_id = "mysql-b51d37110dd1"
}

resource "volcengine_rds_mysql_database" "foo" {
     db_name = "acc-test-db"
     instance_id = volcengine_rds_mysql_instance.foo.id
}

resource "volcengine_rds_mysql_account" "foo" {
     account_name = "acc-test-account"
     account_password = "93f0cb0614Aab12"
     account_type = "Normal"
     instance_id = volcengine_rds_mysql_instance.foo.id
     account_privileges {
          db_name = volcengine_rds_mysql_database.foo.db_name
          account_privilege = "Custom"
          account_privilege_detail = "SELECT,INSERT,UPDATE"
     }
     account_privileges {
          db_name = volcengine_rds_mysql_database.foo1.db_name
          account_privilege = "DDLOnly"
     }
     host = "192.10.10.%"
#     table_column_privileges {
#          db_name = volcengine_rds_mysql_database.foo.db_name
#          table_privileges {
#               table_name = "test"
#               account_privilege_detail = "SELECT,INSERT,UPDATE"
#          }
#          column_privileges {
#               table_name = "test"
#               column_name = "test"
#               account_privilege_detail = "SELECT,INSERT,UPDATE"
#          }
#     }
}