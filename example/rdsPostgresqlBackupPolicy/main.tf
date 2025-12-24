resource "volcengine_rds_postgresql_backup_policy" "example" {
    instance_id                = "postgres-72715e0d9f58"
    backup_retention_period    = 7
    full_backup_period         = "Monday,Wednesday,Friday"
    full_backup_time           = "18:00Z-19:00Z"
    data_incr_backup_periods   = "Tuesday,Sunday"
    hourly_incr_backup_enable  = true
    increment_backup_frequency = 12
    wal_log_space_limit_enable = false
}