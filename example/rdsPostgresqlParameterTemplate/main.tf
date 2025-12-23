resource "volcengine_rds_postgresql_parameter_template" "tpl_base" {
  template_name         = "tf-pg-pt-base"
  template_type         = "PostgreSQL"
  template_type_version = "PostgreSQL_12"
  template_desc         = "base template for clone"

  template_params {
    name  = "auto_explain.log_analyze"
    value = "off"
  }
  template_params {
    name  = "auto_explain.log_buffers"
    value = "on"
  }
}

resource "volcengine_rds_postgresql_parameter_template" "tpl_clone" {
  template_name = "tf-pg-pt-clone"
  src_template_id = "postgresql-b62f5687df914b1c"
  template_desc = "cloned by terraform"
  template_type = "PostgreSQL"
  template_type_version = "PostgreSQL_12"
}

resource "volcengine_rds_postgresql_parameter_template" "tpl_export" {
  template_name = "tf-pg-pt-export"
  instance_id   = "postgres-72715e0d9f58"
  template_desc = "exported from instance"
  template_type = "PostgreSQL"
  template_type_version = "PostgreSQL_12"
}