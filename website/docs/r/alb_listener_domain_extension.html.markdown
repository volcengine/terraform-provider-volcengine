---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_listener_domain_extension"
sidebar_current: "docs-volcengine-resource-alb_listener_domain_extension"
description: |-
  Provides a resource to manage alb listener domain extension
---
# volcengine_alb_listener_domain_extension
Provides a resource to manage alb listener domain extension
## Example Usage
```hcl
resource "volcengine_alb_listener" "foo" {
  load_balancer_id  = "alb-1iidd17v3klj474adhfrunyz9"
  listener_name     = "acc-test-listener-1"
  protocol          = "HTTPS"
  port              = 6666
  enabled           = "on"
  certificate_id    = "cert-1iidd2pahdyio74adhfr9ajwg"
  ca_certificate_id = "cert-1iidd2r9ii0hs74adhfeodxo1"
  server_group_id   = "rsp-1g72w74y4umf42zbhq4k4hnln"
  enable_http2      = "on"
  enable_quic       = "off"
  acl_status        = "on"
  acl_type          = "white"
  acl_ids           = ["acl-1g72w6z11ighs2zbhq4v3rvh4"]
  description       = "acc test listener"
}

resource "volcengine_alb_listener_domain_extension" "foo" {
  listener_id    = volcengine_alb_listener.foo.id
  domain         = "test-modify.com"
  certificate_id = "cert-1iidd2pahdyio74adhfr9ajwg"
}
```
## Argument Reference
The following arguments are supported:
* `domain` - (Required) The domain name. The maximum number of associated domain names for an HTTPS listener is 20, with a value range of 1 to 20.
* `listener_id` - (Required, ForceNew) The listener id. Only HTTPS listener is effective.
* `cert_center_certificate_id` - (Optional) The server certificate ID used by the domain name. Valid when the certificate_source is `cert_center`.
* `certificate_id` - (Optional) Server certificate used for the domain name. Valid when the certificate_source is `alb`.
* `certificate_source` - (Optional) The source of the certificate. Valid values: `alb`, `cert_center`, `pca_leaf`. Default is `alb`.
* `pca_leaf_certificate_id` - (Optional) The server certificate ID used by the domain name. Valid when the certificate source is `pca_leaf`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `domain_extension_id` - The id of the domain extension.


## Import
AlbListenerDomainExtension can be imported using the listener id and domain extension id, e.g.
```
$ terraform import volcengine_alb_listener_domain_extension.default listenerId:extensionId
```

