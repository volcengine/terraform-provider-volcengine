data "volcengine_rds_postgresql_parameter_templates" "templates" {
  template_category     = "DBEngine"
  template_type         = "PostgreSQL"
  template_type_version = "PostgreSQL_12"
}
