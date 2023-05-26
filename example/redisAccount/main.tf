resource "volcengine_redis_account" "foo" {
  instance_id = "redis-cn0398aizj8cwmopx"
  account_name = "test"
  password = "1qaz!QAZ"
  role_name = "ReadOnly"
  description = "test12345"
}