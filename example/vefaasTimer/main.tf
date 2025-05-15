resource "volcengine_vefaas_timer" "foo" {
  function_id = "35ybaxxx"
  name = "test-1-tf"
  crontab = "*/10 * * * *"
}
