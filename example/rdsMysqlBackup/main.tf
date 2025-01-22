resource "volcengine_rds_mysql_backup" "foo" {
    instance_id = "mysql-c8c3f45c4b07"
    #backup_type = "Full"
    backup_method = "Logical"
    backup_meta {
        db_name = "order"
    }
}