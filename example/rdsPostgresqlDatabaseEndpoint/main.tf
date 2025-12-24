resource "volcengine_rds_postgresql_database_endpoint" "cluster" {
  endpoint_id                      = "postgres-72715e0d9f58-cluster"
  endpoint_name                    = "默认终端"
  endpoint_type                    = "Cluster"
  instance_id                      = "postgres-72715e0d9f58"
  read_only_node_distribution_type = "Custom"
  read_only_node_max_delay_time    = 40
  read_write_mode                  = "ReadWrite"
  read_write_proxy_connection      = 20
  write_node_halt_writing          = false
  read_write_splitting             = true
  read_only_node_weight {
      node_id   = null
      node_type = "Primary"
      weight    = 200
  }
  dns_visibility                   = true
  port                             = 5432
}

resource "volcengine_rds_postgresql_database_endpoint" "example"{
  instance_id                      = "postgres-72715e0d9f58"
  endpoint_name                    = "tf-test"
  endpoint_type                    = "Custom"
  nodes                            = "Primary"
  read_write_mode                  = "ReadWrite"
}