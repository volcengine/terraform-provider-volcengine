resource "volcengine_rds_postgresql_allowlist" "foo" {
  allow_list_name = "acc-test-allowlist"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = ["10.0.0.0/24"]
  security_group_bind_infos {
      security_group_id = "sg-1jojfhw8rca9s1n7ampztrq6w"
      bind_mode         = "IngressDirectionIp"
  }
}
resource "volcengine_rds_postgresql_allowlist" "example" {
  instance_ids = ["postgres-72715e0d9f58","postgres-eb3a578a6d73"]
  allow_list_name = "unify_new"
}