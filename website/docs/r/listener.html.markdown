---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_listener"
sidebar_current: "docs-volcengine-resource-listener"
description: |-
  Provides a resource to manage listener
---
# volcengine_listener
Provides a resource to manage listener
## Example Usage
```hcl
resource "volcengine_listener" "foo" {
  load_balancer_id = "clb-274xltt3rfmyo7fap8sv1jq39"
  listener_name    = "Demo-HTTP-90"
  protocol         = "HTTP"
  port             = 90
  server_group_id  = "rsp-274xltv2sjoxs7fap8tlv3q3s"
  health_check {
    enabled              = "on"
    interval             = 10
    timeout              = 3
    healthy_threshold    = 5
    un_healthy_threshold = 2
    domain               = "volcengine.com"
    http_code            = "http_2xx"
    method               = "GET"
    uri                  = "/"
  }
  enabled = "on"
}

resource "volcengine_listener" "bar" {
  load_balancer_id = "clb-274xltt3rfmyo7fap8sv1jq39"
  listener_name    = "Demo-HTTP-91"
  protocol         = "HTTP"
  port             = 91
  server_group_id  = "rsp-274xltv2sjoxs7fap8tlv3q3s"
  health_check {
    enabled              = "on"
    interval             = 10
    timeout              = 3
    healthy_threshold    = 5
    un_healthy_threshold = 2
    domain               = "volcengine.com"
    http_code            = "http_2xx"
    method               = "GET"
    uri                  = "/"
  }
  enabled = "on"
}

resource "volcengine_listener" "demo" {
  load_balancer_id = "clb-274xltt3rfmyo7fap8sv1jq39"
  protocol         = "TCP"
  port             = 92
  server_group_id  = "rsp-274xltv2sjoxs7fap8tlv3q3s"
  health_check {
    enabled              = "on"
    interval             = 10
    timeout              = 3
    healthy_threshold    = 5
    un_healthy_threshold = 2
  }
  enabled             = "on"
  established_timeout = 10
}
```
## Argument Reference
The following arguments are supported:
* `load_balancer_id` - (Required, ForceNew) The region of the request.
* `port` - (Required, ForceNew) The port receiving request of the Listener, the value range in 1~65535.
* `protocol` - (Required, ForceNew) The protocol of the Listener. Optional choice contains `TCP`, `UDP`, `HTTP`, `HTTPS`.
* `server_group_id` - (Required) The server group id associated with the listener.
* `acl_ids` - (Optional) The id list of the Acl.
* `acl_status` - (Optional) The enable status of Acl. Optional choice contains `on`, `off`.
* `acl_type` - (Optional) The type of the Acl. Optional choice contains `white`, `black`.
* `certificate_id` - (Optional) The certificate id associated with the listener.
* `description` - (Optional) The description of the Listener.
* `enabled` - (Optional) The enable status of the Listener. Optional choice contains `on`, `off`.
* `established_timeout` - (Optional) The connection timeout of the Listener.
* `health_check` - (Optional) The config of health check.
* `listener_name` - (Optional) The name of the Listener.
* `scheduler` - (Optional) The scheduling algorithm of the Listener. Optional choice contains `wrr`, `wlc`, `sh`.

The `health_check` object supports the following:

* `domain` - (Optional) The domain of health check.
* `enabled` - (Optional) The enable status of health check function. Optional choice contains `on`, `off`.
* `healthy_threshold` - (Optional) The healthy threshold of health check, default 3, range in 2~10.
* `http_code` - (Optional) The normal http status code of health check, the value can be `http_2xx` or `http_3xx` or `http_4xx` or `http_5xx`.
* `interval` - (Optional) The interval executing health check, default 2, range in 1~300.
* `method` - (Optional) The method of health check, the value can be `GET` or `HEAD`.
* `timeout` - (Optional) The response timeout of health check, default 2, range in 1~60..
* `un_healthy_threshold` - (Optional) The unhealthy threshold of health check, default 3, range in 2~10.
* `uri` - (Optional) The uri of health check.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `listener_id` - The ID of the Listener.


## Import
Listener can be imported using the id, e.g.
```
$ terraform import volcengine_listener.default lsn-273yv0mhs5xj47fap8sehiiso
```

