data "volcengine_redis_backups" "default" {
  instance_id = "redis-cnlfvrv4qye6u4lpa"
  backup_strategy_list = ["ManualBackup"]
}