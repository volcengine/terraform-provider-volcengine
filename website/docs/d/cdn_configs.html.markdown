---
subcategory: "CDN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cdn_configs"
sidebar_current: "docs-volcengine-datasource-cdn_configs"
description: |-
  Use this data source to query detailed information of cdn configs
---
# volcengine_cdn_configs
Use this data source to query detailed information of cdn configs
## Example Usage
```hcl
resource "volcengine_cdn_certificate" "foo" {
  certificate = ""
  private_key = ""
  desc        = "tftest"
  source      = "cdn_cert_hosting"
}

resource "volcengine_cdn_domain" "foo" {
  domain       = "tftest.byte-test.com"
  service_type = "web"
  tags {
    key   = "tfkey1"
    value = "tfvalue1"
  }
  tags {
    key   = "tfkey2"
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
                Address             = "1.1.1.1",
                HttpPort            = "80",
                HttpsPort           = "443",
                InstanceType        = "ip",
                OriginType          = "primary",
                PrivateBucketAccess = false,
                Weight              = "2"
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
        HTTP2       = true,
        Switch      = true,
        Ocsp        = false,
        TlsVersion = [
          "tlsv1.1",
          "tlsv1.2"
        ],
      }
    }
  )
}

data "volcengine_cdn_configs" "foo" {
  domain = volcengine_cdn_domain.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `domain` - (Required) The domain name.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `domain_config` - The collection of query.
    * `cname` - The cname of the domain.
    * `create_time` - The create time of the domain.
    * `domain` - The domain name.
    * `lock_status` - Indicates whether the configuration of this domain name is allowed to be changed.
    * `project` - The project name.
    * `service_region` - The service region of the domain.
    * `service_type` - The service type of the domain.
    * `status` - The status of the domain.
    * `update_time` - The update time of the domain.
* `total_count` - The total count of query.


