data "volcengine_rds_postgresql_parameter_template_apply_diffs" "diffs" {
  instance_id = "postgres-72715e0d9f58"
  template_id = "postgresql-ef66e3807988595a"
}