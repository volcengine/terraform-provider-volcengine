data "volcengine_rds_postgresql_instance_price_differences" "example" {
  instance_id = "postgres-72715e0d9f58"

  modify_type = "Usually"
  
  node_info {
    node_id = "postgres-72715e0d9f58"
    zone_id   = "cn-beijing-a"
    node_type = "Primary"
    node_spec = "rds.postgres.2c4g"
  }
  node_info {
    node_id = "postgres-72715e0d9f58-iyys"
    zone_id   = "cn-beijing-a"
    node_type = "Secondary"
    node_spec = "rds.postgres.2c4g"
  }

  storage_type  = "LocalSSD"
  storage_space = 100

  charge_info {
    charge_type = "PostPaid"
    number      = 1
  }
}