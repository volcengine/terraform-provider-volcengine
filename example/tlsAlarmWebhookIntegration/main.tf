resource "volcengine_tls_alarm_webhook_integration" "foo" {
  webhook_name   = "terraform-tf-webhook-modify"
  webhook_url    = "http://tencent.com"
  webhook_type   = "wechat"
  webhook_method = "PUT"
  webhook_secret = "your secret"
  webhook_headers {
    key   = "Content-Type"
    value = "application/json"
  }
}

