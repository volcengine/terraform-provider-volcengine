resource "volcengine_cdn_domain" "foo" {
    domain = "www.tftest.com"
    service_type = "web"
    origin_protocol = "http"
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
        }
    )
}