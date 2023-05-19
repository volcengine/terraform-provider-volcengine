resource "volcengine_rds_mysql_allowlist" "foo" {
    allow_list_name = "tf-test-opt"
    allow_list_desc = "terraform test zzm"
    allow_list = [
        "127.0.0.1"
    ]
}