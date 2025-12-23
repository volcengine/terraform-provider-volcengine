data "volcengine_rds_postgresql_instance_price_details" "example" {
  node_info {
    zone_id           = "cn-beijing-a"
    node_type         = "Primary"
    node_spec         = "rds.postgres.1c2g"
    node_operate_type = "Create"
  }
  node_info {
    zone_id           = "cn-beijing-a"
    node_type         = "Secondary"
    node_spec         = "rds.postgres.1c2g"
    node_operate_type = "Create"
  }
  node_info {
    zone_id           = "cn-beijing-a"
    node_type         = "ReadOnly"
    node_spec         = "rds.postgres.2c8g"
    node_operate_type = "Create"
  }
  storage_type  = "LocalSSD"
  storage_space = 100

  charge_info {
    charge_type = "PrePaid"
    period_unit = "Month"
    period      = 2
    number      = 4
  }
}