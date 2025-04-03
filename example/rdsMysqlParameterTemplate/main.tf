resource "volcengine_rds_mysql_parameter_template" "foo" {
    template_name = "test"
    template_type = "Mysql"
    template_type_version = "MySQL_8_0"
    template_params {
        name = "auto_increment_increment"
        running_value = "1"
    }
    template_desc = "test"
}