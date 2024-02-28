resource "volcengine_cdn_domain" "foo" {
    domain = "tftest.1705568676000-509.byte-test.com"
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
                    CertId = "cert-69c97c462c3740218b204c45fe6db690"
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