resource "volcengine_vefaas_function" "foo" {
  name = "test-tf-1"
  runtime = "golang/v1"
  project_name = "tf-test"
  description = "test"
  exclusive_mode = true
  request_timeout = 100
}

