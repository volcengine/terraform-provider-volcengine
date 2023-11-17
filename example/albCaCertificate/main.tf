resource "volcengine_alb_ca_certificate" "foo" {
  ca_certificate_name = "acc-test-1"
  ca_certificate      = "-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----"
  description = "acc-test-1"
  project_name = "default"
}