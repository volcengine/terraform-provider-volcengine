variable "certificate" {
  type = object({
    public_key = string
    private_key = string
  })
}

resource "vestack_certificate" "foo" {
  certificate_name = "demo-certificate"
  description = "This is a clb certificate"
  public_key = var.certificate.public_key
  private_key = var.certificate.private_key
}