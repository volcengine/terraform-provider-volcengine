---
subcategory: "CDN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cdn_domains"
sidebar_current: "docs-volcengine-datasource-cdn_domains"
description: |-
  Use this data source to query detailed information of cdn domains
---
# volcengine_cdn_domains
Use this data source to query detailed information of cdn domains
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

data "volcengine_cdn_domains" "foo" {
  domain = volcengine_cdn_domain.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `domain` - (Optional) Search by specifying domain name keywords, with fuzzy matching.
* `https` - (Optional) Specify HTTPS configuration to filter accelerated domains. The optional values for this parameter are as follows: `true`: Indicates that the accelerated domain has enabled HTTPS function.`false`: Indicates that the accelerated domain has not enabled HTTPS function.
* `ipv6` - (Optional) Specify IPv6 configuration to filter accelerated domain names. The optional values for this parameter are as follows: `true`: Indicates that the accelerated domain name supports requests using IPv6 addresses.`false`: Indicates that the accelerated domain name does not support requests using IPv6 addresses.
* `origin_protocol` - (Optional) Configure the origin protocol for the accelerated domain.
* `output_file` - (Optional) File name where to save data source results.
* `primary_origin` - (Optional) Specify a primary origin server for filtering accelerated domains.
* `project` - (Optional) The project name of the domain.
* `service_type` - (Optional) The business type of the domain name is indicated by this parameter. The possible values are: `download`: for file downloads. `web`: for web pages. `video`: for audio and video on demand.
* `status` - (Optional) The status of the domain.
* `tags` - (Optional) Filter by specified domain name tags, up to 10 tags can be specified. Each tag is entered as a string in the format of key:value.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `domains` - The collection of query.
    * `backup_origin` - The list of backup origin servers for accelerating this domain name. If no backup origin server is configured for this acceleration domain name, the parameter value is null.
    * `cache_shared_target_host` - If CacheShared is cache_shared_on, it means the target domain name that shares cache with the accelerated domain name. If CacheShared is target_host or an empty value, the parameter value is empty.
    * `cache_shared` - Indicates the role of the accelerated domain in the shared cache configuration. This parameter can take the following values: `target_host`: Indicates that there is a shared cache configuration where the role of the accelerated domain is the target domain.`cache_shared_on`: Indicates that there is a shared cache configuration where the role of the accelerated domain is the configured domain.`""`: This parameter value is empty, indicating that the accelerated domain does not exist in any shared cache configuration.
    * `cname` - The CNAME address of the domain is automatically assigned when adding the domain.
    * `create_time` - The creation time of the domain.
    * `domain_lock` - Indicates the locked status of the accelerated domain.
        * `remark` - If the Status is on, this parameter value records the reason for the lock.
        * `status` - Indicates whether the domain name is locked.
    * `domain` - Search by specifying domain name keywords, with fuzzy matching.
    * `https` - Specify HTTPS configuration to filter accelerated domains. The optional values for this parameter are as follows: `true`: Indicates that the accelerated domain has enabled HTTPS function.`false`: Indicates that the accelerated domain has not enabled HTTPS function.
    * `ipv6` - Specify IPv6 configuration to filter accelerated domain names. The optional values for this parameter are as follows: `true`: Indicates that the accelerated domain name supports requests using IPv6 addresses.`false`: Indicates that the accelerated domain name does not support requests using IPv6 addresses.
    * `is_conflict_domain` - Indicates whether the accelerated domain name is a conflicting domain name. By default, each accelerated domain name is unique in the content delivery network. If you need to add an accelerated domain name that already exists in the content delivery network, you need to submit a ticket. If the domain name is added successfully, it becomes a conflicting domain name.
    * `origin_protocol` - Configure the origin protocol for the accelerated domain.
    * `primary_origin` - List of primary source servers to accelerate the domain name.
    * `project` - The project name of the domain.
    * `service_region` - Indicates the acceleration area. The parameter can take the following values: `chinese_mainland`: Indicates mainland China. `global`: Indicates global. `outside_chinese_mainland`: Indicates global (excluding mainland China).
    * `service_type` - The business type of the domain name is indicated by this parameter. The possible values are: `download`: for file downloads. `web`: for web pages. `video`: for audio and video on demand.
    * `status` - The status of the domain.
    * `tags` - Indicate the tags you have set for this domain name. You can set up to 10 tags.
        * `key` - The key of the tag.
        * `value` - The value of the tag.
    * `update_time` - The update time of the domain.
* `total_count` - The total count of query.


