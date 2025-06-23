resource "volcengine_rds_mysql_backup_policy" "foo" {
    instance_id = "mysql-b51d37110dd1"
    data_full_backup_periods = ["Monday", "Sunday", "Tuesday"]
    binlog_file_counts_enable = true
    binlog_space_limit_enable = true
    lock_ddl_time = 80
    cross_backup_policy {
        backup_enabled = true
        cross_backup_region = "cn-beijing"
        log_backup_enabled = true
        retention = 10
    }
}