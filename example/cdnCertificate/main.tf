resource "volcengine_cdn_certificate" "foo" {
  certificate {
    certificate = "-----BEGIN CERTIFICATE----------END CERTIFICATE-----"
    private_key = "-----BEGIN RSA PRIVATE KEY----------END RSA PRIVATE KEY-----"
  }
  source = "cdn_cert_hosting"
}