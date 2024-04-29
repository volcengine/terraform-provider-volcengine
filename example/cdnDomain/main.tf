resource "volcengine_cdn_certificate" "foo" {
    certificate = ""
    private_key = ""
    desc = "tftest"
    source = "cdn_cert_hosting"
}

resource "volcengine_cdn_domain" "foo" {
    domain = "tftest.byte-test.com"
    service_type = "web"
    tags {
        key = "tfkey1"
        value = "tfvalue1"
    }
    tags {
        key = "tfkey2"
        value = "tfvalue2"
    }
    domain_config = jsonencode(
        {
            OriginProtocol = "https"
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