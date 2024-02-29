resource "volcengine_cdn_certificate" "foo" {
    certificate {
        certificate = ""
        private_key = ""
    }
    source = "cdn_cert_hosting"
}

resource "volcengine_cdn_domain" "foo" {
    domain = "tftest.byte-test.com"
    service_type = "web"
    origin_protocol = "https"
    domain_config = jsonencode(
        {
            Origin = [
                {
                    OriginAction = {
                        OriginLines = [
                            {
                                Address = "1.1.1.1",
                                HttpPort = "80",
                                HttpsPort = "443",
                                InstanceType = "ip",
                                OriginType = "primary",
                                PrivateBucketAccess = false,
                                Weight = "2"
                            }
                        ]
                    }
                }
            ]
            HTTPS = {
                CertInfo = {
                    CertId = volcengine_cdn_certificate.foo.id
                }
                DisableHttp = false,
                HTTP2 = true,
                Switch = true,
                Ocsp = false,
                TlsVersion = [
                    "tlsv1.1",
                    "tlsv1.2"
                ],
            }
        }
    )
}