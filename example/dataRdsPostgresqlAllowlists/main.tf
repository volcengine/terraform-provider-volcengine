data "volcengine_rds_postgresql_allowlists" "default" {
  name_regex = ".*allowlist.*"
  allow_list_id = "acl-e7846436e1e741edbd385868fa657436"
  allow_list_category = "Ordinary"
  allow_list_desc = "test allow list"
  allow_list_name = "test"
  ip_address = "100.64.0.0/10"
}