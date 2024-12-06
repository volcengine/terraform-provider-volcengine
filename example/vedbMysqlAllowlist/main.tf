resource "volcengine_vedb_mysql_allowlist" "foo" {
    allow_list_name = "acc-test-allowlist"
    allow_list_desc = "acc-test"
    allow_list_type = "IPv4"
    allow_list = ["192.168.0.0/24", "192.168.1.0/24", "192.168.2.0/24"]
}