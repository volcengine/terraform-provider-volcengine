resource "volcengine_rds_mysql_backup_policy" "foo" {
    instance_id = "mysql-c8c3f45c4b07"
    data_full_backup_periods = ["Monday", "Sunday"]
    binlog_file_counts_enable = true
    binlog_space_limit_enable = true
    lock_ddl_time = 80
}