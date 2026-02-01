---
subcategory: "CDN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cdn_domain"
sidebar_current: "docs-volcengine-resource-cdn_domain"
description: |-
  Provides a resource to manage cdn domain
---
# volcengine_cdn_domain
Provides a resource to manage cdn domain
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
```
## Argument Reference
The following arguments are supported:
* `domain_config` - (Required) Accelerate domain configuration. Please convert the configuration module structure into json and pass it into a string. You must specify the Origin module. The OriginProtocol parameter, OriginHost parameter, and other domain configuration modules are optional.
* `domain` - (Required, ForceNew) You need to add a domain. The main account can add up to 200 accelerated domains.
* `service_type` - (Required, ForceNew) The business type of the domain name is indicated by this parameter. The possible values are: `download`: for file downloads. `web`: for web pages. `video`: for audio and video on demand.
* `project` - (Optional, ForceNew) The project to which this domain name belongs. Default is `default`.
* `service_region` - (Optional, ForceNew) Indicates the acceleration area. The parameter can take the following values: `chinese_mainland`: Indicates mainland China. `global`: Indicates global. `outside_chinese_mainland`: Indicates global (excluding mainland China).
* `shared_cname` - (Optional, ForceNew) Configuration for sharing CNAME.
* `tags` - (Optional) Indicate the tags you have set for this domain name. You can set up to 10 tags.

The `shared_cname` object supports the following:

* `cname` - (Required, ForceNew) Assign a CNAME to the accelerated domain.
* `switch` - (Required, ForceNew) Specify whether to enable shared CNAME.

The `tags` object supports the following:

* `key` - (Required) The key of the tag.
* `value` - (Required) The value of the tag.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `status` - The status of the domain.


## Import
CdnDomain can be imported using the domain, e.g.
```
$ terraform import volcengine_cdn_domain.default www.volcengine.com
```
Please note that when you execute destroy, we will first take the domain name offline and then delete it.

