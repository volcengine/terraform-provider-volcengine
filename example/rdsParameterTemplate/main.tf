resource "volcengine_rds_parameter_template" "foo" {
  template_desc = "created by terraform"
  template_name = "tf-template"
  template_type = "MySQL"
  template_type_version = "MySQL_Community_5_7"
  template_params {
    name = "auto_increment_increment"
    running_value = "2"
  }
  template_params {
    name = "slow_query_log"
    running_value = "ON"
  }
  template_params {
    name = "net_retry_count"
    running_value = "33"
  }
}