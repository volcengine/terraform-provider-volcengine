---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_listener"
sidebar_current: "docs-volcengine-resource-alb_listener"
description: |-
  Provides a resource to manage alb listener
---
# volcengine_alb_listener
Provides a resource to manage alb listener
## Example Usage
```hcl
resource "volcengine_alb_customized_cfg" "foo" {
  customized_cfg_name    = "acc-test-cfg1"
  description            = "This is a test modify"
  customized_cfg_content = "proxy_connect_timeout 4s;proxy_request_buffering on;"
  project_name           = "default"
}

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
  acl_ids           = ["acl-1g72w6z11ighs2zbhq4v3rvh4", "acl-1g72xvtt7kg002zbhq5diim3s"]
  description       = "acc test listener"
  customized_cfg_id = volcengine_alb_customized_cfg.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `load_balancer_id` - (Required, ForceNew) The Id of the load balancer.
* `port` - (Required, ForceNew) The port receiving request of the Listener, the value range in 1~65535.
* `protocol` - (Required, ForceNew) The protocol of the Listener. Optional choice contains `HTTP`, `HTTPS`.
* `server_group_id` - (Required) The server group id associated with the listener.
* `acl_ids` - (Optional) The id list of the Acl. When the AclStatus parameter is configured as on, AclType and AclIds.N are required.
* `acl_status` - (Optional) The enable status of Acl. Optional choice contains `on`, `off`. Default is `off`.
* `acl_type` - (Optional) The type of the Acl. Optional choice contains `white`, `black`. When the AclStatus parameter is configured as on, AclType and AclIds.N are required.
* `ca_certificate_id` - (Optional) The CA certificate id associated with the listener.
* `certificate_id` - (Optional) The certificate id associated with the listener.
* `customized_cfg_id` - (Optional) Personalized configuration ID, with a value of " " when not bound.
* `description` - (Optional) The description of the Listener.
* `enable_http2` - (Optional) The HTTP2 feature switch,valid value is on or off. Default is `off`.
* `enable_quic` - (Optional) The QUIC feature switch,valid value is on or off. Default is `off`.
* `enabled` - (Optional) The enable status of the Listener. Optional choice contains `on`, `off`. Default is `on`.
* `listener_name` - (Optional) The name of the Listener.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `listener_id` - The ID of the Listener.


## Import
AlbListener can be imported using the id, e.g.
```
$ terraform import volcengine_alb_listener.default lsn-273yv0mhs5xj47fap8sehiiso
```

