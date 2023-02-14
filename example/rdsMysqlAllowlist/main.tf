resource "volcengine_rds_mysql_allowlist" "foo" {
<<<<<<< HEAD
    allow_list_name = "tf-test-opt"
    allow_list_desc = "terraform test zzm"
    allow_list = [
        "127.0.0.1"
    ]
=======
    allow_list_name = "tf-test"
    allow_list_desc = "terraform test zzm"
    allow_list = "127.0.0.1"
>>>>>>> 7bde4df (feat: add rds mysql allowlist and allowlist_associate)
}