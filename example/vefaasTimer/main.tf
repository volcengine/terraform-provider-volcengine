resource "volcengine_vefaas_timer" "foo" {
  function_id = "g79asjxxx"
  name = "tf-test"
  description = "test-tf-modify"
  crontab = "*/5 * * * *"
  payload = "hello world"
  enabled = true
  enable_concurrency = true
}
