data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
     vpc_name = "acc-test-project1"
     cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
     subnet_name = "acc-subnet-test-2"
     cidr_block = "172.16.0.0/24"
     zone_id = data.volcengine_zones.foo.zones[0].id
     vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_rds_mysql_instance" "foo" {
     db_engine_version = "MySQL_5_7"
     node_spec = "rds.mysql.1c2g"
     primary_zone_id = data.volcengine_zones.foo.zones[0].id
     secondary_zone_id = data.volcengine_zones.foo.zones[0].id
     storage_space = 80
     subnet_id = volcengine_subnet.foo.id
     instance_name = "acc-test"
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

resource "volcengine_rds_mysql_database" "foo" {
     db_name = "acc-test"
     instance_id = volcengine_rds_mysql_instance.foo.id
     db_desc = "test-update"
     database_privileges {
          account_name = "acc-test"
          host = "%"
          account_privilege = "Custom"
          account_privilege_detail = "CREATE,DROP,ALTER,INDEX"
     }
}