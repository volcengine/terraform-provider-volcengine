resource "volcengine_rds_postgresql_endpoint_public_address" "example" {
  instance_id = "postgres-ac541555dd74"
  endpoint_id = "postgres-ac541555dd74-cluster"
  eip_id      = "eip-1c0x0ehrbhb7k5e8j71k84ryd"
}