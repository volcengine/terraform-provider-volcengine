data "volcengine_alb_certificates" "default" {
  certificate_name = "tf-test"
  tags {
    key = "k1"
    value = "v1"
  }
}
