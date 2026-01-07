resource "volcengine_tls_alarm_webhook_integration" "foo" {
  webhook_name   = "terraform-tf-webhook"
  webhook_url    = "http://zijie.com"
  webhook_type   = "lark"
  webhook_method = "PUT"
  webhook_secret = "****-****-***"
  webhook_headers {
    key   = "Content-Type"
    value = "application/json"
  }
}

