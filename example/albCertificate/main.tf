resource "volcengine_alb_certificate" "foo" {
  description = "test123"
  public_key  = "public key"
  private_key = "private key"
}