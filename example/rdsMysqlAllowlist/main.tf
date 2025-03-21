resource "volcengine_rds_mysql_allowlist" "foo" {
    allow_list_name = "acc-test-allowlist"
    allow_list_desc = "acc-test"
    allow_list_type = "IPv4"
    user_allow_list = ["192.168.0.0/24", "192.168.1.0/24"]
    //user_allow_list = ["192.168.0.0/24", "192.168.1.0/24"]
    security_group_bind_infos {
        bind_mode = "IngressDirectionIp"
        security_group_id = "sg-mjoa9qfyzg1s5smt1a6dmc1l"
    }
}