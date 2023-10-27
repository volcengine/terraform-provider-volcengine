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
data "volcengine_zones" "foo" {
}
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_clb" "foo" {
  type               = "public"
  subnet_id          = volcengine_subnet.foo.id
  load_balancer_spec = "small_1"
  description        = "acc0Demo"
  load_balancer_name = "acc-test-create"
  eip_billing_config {
    isp              = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth        = 1
  }
}

resource "volcengine_server_group" "foo" {
  load_balancer_id  = volcengine_clb.foo.id
  server_group_name = "acc-test-create"
  description       = "hello demo11"
}

resource "volcengine_listener" "foo" {
  load_balancer_id = volcengine_clb.foo.id
  listener_name    = "acc-test-listener"
  protocol         = "HTTP"
  port             = 90
  server_group_id  = volcengine_server_group.foo.id
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

resource "volcengine_listener" "foo_tcp" {
  load_balancer_id         = volcengine_clb.foo.id
  listener_name            = "acc-test-listener"
  protocol                 = "TCP"
  port                     = 90
  server_group_id          = volcengine_server_group.foo.id
  enabled                  = "on"
  bandwidth                = 2
  proxy_protocol_type      = "standard"
  persistence_type         = "source_ip"
  persistence_timeout      = 100
  connection_drain_enabled = "on"
  connection_drain_timeout = 100
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
* `bandwidth` - (Optional) The bandwidth of the Listener. Unit: Mbps. Default is -1, indicating that the Listener does not specify a speed limit.
* `certificate_id` - (Optional) The certificate id associated with the listener.
* `connection_drain_enabled` - (Optional) Whether to enable connection drain of the Listener. Valid values: `off`, `on`. Default is `off`.
This filed is valid only when the value of field `protocol` is `TCP` or `UDP`.
* `connection_drain_timeout` - (Optional) The connection drain timeout of the Listener. Valid value range is `1-900`.
This filed is required when the value of field `connection_drain_enabled` is `on`.
* `description` - (Optional) The description of the Listener.
* `enabled` - (Optional) The enable status of the Listener. Optional choice contains `on`, `off`.
* `established_timeout` - (Optional) The connection timeout of the Listener.
* `health_check` - (Optional) The config of health check.
* `listener_name` - (Optional) The name of the Listener.
* `persistence_timeout` - (Optional) The persistence timeout of the Listener. Unit: second. Valid value range is `1-3600`. Default is `1000`.
This filed is valid only when the value of field `persistence_type` is `source_ip`.
* `persistence_type` - (Optional) The persistence type of the Listener. Valid values: `off`, `source_ip`. Default is `off`.
This filed is valid only when the value of field `protocol` is `TCP` or `UDP`.
* `proxy_protocol_type` - (Optional) Whether to enable proxy protocol. Valid values: `off`, `standard`. Default is `off`.
This filed is valid only when the value of field `protocol` is `TCP` or `UDP`.
* `scheduler` - (Optional) The scheduling algorithm of the Listener. Optional choice contains `wrr`, `wlc`, `sh`.

The `health_check` object supports the following:

* `domain` - (Optional) The domain of health check.
* `enabled` - (Optional) The enable status of health check function. Optional choice contains `on`, `off`.
* `healthy_threshold` - (Optional) The healthy threshold of health check, default 3, range in 2~10.
* `http_code` - (Optional) The normal http status code of health check, the value can be `http_2xx` or `http_3xx` or `http_4xx` or `http_5xx`.
* `interval` - (Optional) The interval executing health check, default 2, range in 1~300.
* `method` - (Optional) The method of health check, the value can be `GET` or `HEAD`.
* `timeout` - (Optional) The response timeout of health check, default 2, range in 1~60..
* `udp_expect` - (Optional) The UDP expect of health check. This field must be specified simultaneously with field `udp_request`.
* `udp_request` - (Optional) The UDP request of health check. This field must be specified simultaneously with field `udp_expect`.
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

