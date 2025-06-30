resource "volcengine_vefaas_function" "foo" {
  name = "tf-1"
  runtime = "golang/v1"
  description = "123131231"
  exclusive_mode = false
  request_timeout = 30
}
