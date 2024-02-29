resource "volcengine_cdn_certificate" "foo" {
    certificate {
        certificate = ""
        private_key = ""
    }
    source = "cdn_cert_hosting"
}

data "volcengine_cdn_certificates" "foo"{
    source = volcengine_cdn_certificate.foo.source
}