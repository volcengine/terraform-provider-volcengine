resource "volcengine_cdn_certificate" "foo" {
    certificate = ""
    private_key = ""
    desc = "tftest"
    source = "cdn_cert_hosting"
}

data "volcengine_cdn_certificates" "foo"{
    source = volcengine_cdn_certificate.foo.source
}