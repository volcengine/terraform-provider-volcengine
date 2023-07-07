resource "volcengine_certificate" "foo" {
  certificate_name = "demo-certificate"
  description = "This is a clb certificate"
  public_key = "public-key"
  private_key = "private-key"
  tags  {
    key = "k1"
    value = "v1"
  }
}