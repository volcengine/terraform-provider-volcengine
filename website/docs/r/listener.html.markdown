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
  tags {
    key   = "k1"
    value = "v1"
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

resource "volcengine_listener" "foo_https" {
  load_balancer_id = volcengine_clb.foo.id
  listener_name    = "acc-test-listener-https"
  protocol         = "HTTPS"
  port             = 100
  server_group_id  = volcengine_server_group.foo.id
  health_check {
    enabled              = "on"
    interval             = 10
    timeout              = 3
    healthy_threshold    = 5
    un_healthy_threshold = 2
    domain               = "volcengine.com"
    http_code            = "http_2xx,http_3xx"
    method               = "GET"
    uri                  = "/"
  }
  enabled               = "on"
  client_header_timeout = 80
  client_body_timeout   = 80
  keepalive_timeout     = 80
  proxy_connect_timeout = 20
  proxy_send_timeout    = 1800
  proxy_read_timeout    = 1800
  certificate_source    = "clb"
  certificate_id        = "cert-mjpctunmog745smt1a******"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `load_balancer_id` - (Required, ForceNew) The region of the request.
* `port` - (Required, ForceNew) The port receiving request of the Listener, the value range in 0~65535. When `protocol` is `TCP` or `UDP`, 0 can be passed in, indicating that full port listening is enabled.
* `protocol` - (Required, ForceNew) The protocol of the Listener. Optional choice contains `TCP`, `UDP`, `HTTP`, `HTTPS`.
* `server_group_id` - (Required) The server group id associated with the listener.
* `acl_ids` - (Optional) The id list of the Acl.
* `acl_status` - (Optional) The enable status of Acl. Optional choice contains `on`, `off`.
* `acl_type` - (Optional) The type of the Acl. Optional choice contains `white`, `black`.
* `bandwidth` - (Optional) The bandwidth of the Listener. Unit: Mbps. Default is -1, indicating that the Listener does not specify a speed limit.
* `ca_certificate_id` - (Optional) The ID of the CA certificate which is associated with the listener. When `ca_enabled` is `on`, this parameter is required.
* `ca_enabled` - (Optional) Whether to enable CACertificate two-way authentication. Values: on, off.
* `cert_center_certificate_id` - (Optional) The ID of the certificate in Certificate Center. When `certificate_source` is `cert_center`, this parameter is required.
* `certificate_id` - (Optional) The certificate id associated with the listener.
* `certificate_source` - (Optional) The source of the certificate which is associated with the listener. Values: `clb`, `cert_center`.
* `client_body_timeout` - (Optional) The client body timeout of the Listener. Only HTTP/HTTPS listeners support this parameter. value range: 30-120.
* `client_header_timeout` - (Optional) The client header timeout of the Listener. Only HTTP/HTTPS listeners support this parameter, i.e., `protocol`=`HTTP` or `HTTPS`. value range: 30-120.
* `connection_drain_enabled` - (Optional) Whether to enable connection drain of the Listener. Valid values: `off`, `on`. Default is `off`.
This filed is valid only when the value of field `protocol` is `TCP` or `UDP`.
* `connection_drain_timeout` - (Optional) The connection drain timeout of the Listener. Valid value range is `0-900`.
This filed is required when the value of field `connection_drain_enabled` is `on`.
* `cookie` - (Optional) The name of the cookie for session persistence configured on the backend server. When PersistenceType is configured as `server`, this parameter is required. When PersistenceType is configured as any other value, this parameter is not effective.
* `cps` - (Optional) The maximum number of new connections per second allowed for the Listener. Default value: `-1`, no limit, which is the upper limit of new connections for the CLB instance.
* `description` - (Optional) The description of the Listener.
* `enabled` - (Optional) The enable status of the Listener. Optional choice contains `on`, `off`.
* `end_port` - (Optional, ForceNew) The end port for full port listening, with a value range of 1-65535. When `port` is 0, this parameter is required, and must be greater than `start_port`.
* `established_timeout` - (Optional) The connection timeout of the Listener.
* `health_check` - (Optional) The config of health check.
* `http2_enabled` - (Optional) Whether the HTTPS protocol listener enables the front-end HTTP 2.0 protocol. value range: `on`, `off`.
* `keepalive_timeout` - (Optional) The timeout period for the long connection between the client and the CLB. Only HTTP/HTTPS listeners support this parameter. value range: 0-900.
* `listener_name` - (Optional) The name of the Listener.
* `max_connections` - (Optional) The maximum number of connections allowed for the Listener. Default value: `-1`, no limit, which is the upper limit of new connections for the CLB instance.
* `persistence_timeout` - (Optional) The persistence timeout of the Listener. Unit: second. Default is `1000`. When PersistenceType is configured as source_ip, the value range is 1-3600. When PersistenceType is configured as insert, the value range is 1-86400. This filed is valid only when the value of field `persistence_type` is `source_ip` or `insert`.
* `persistence_type` - (Optional) The persistence type of the Listener. Valid values: `off`, `source_ip`, `insert`, `server`. Default is `off`.
`source_ip`: Represents the source IP address, only effective for TCP/UDP protocols. `insert`: means implanting a cookie, only effective for HTTP/HTTPS protocol and when the scheduler is `wrr`. `server`: Indicates rewriting cookies, only effective for HTTP/HTTPS protocols and when the scheduler is `wrr`.
* `proxy_connect_timeout` - (Optional) The timeout period for establishing a connection between the CLB and the backend server. Only HTTP/HTTPS listeners support this parameter. value range: 4-120.
* `proxy_protocol_type` - (Optional) Whether to enable proxy protocol. Valid values: `off`, `standard`. Default is `off`.
This filed is valid only when the value of field `protocol` is `TCP` or `UDP`.
* `proxy_read_timeout` - (Optional) The timeout period for CLB to read the response from the backend server. Only HTTP/HTTPS listeners support this parameter. value range: 30-3600.
* `proxy_send_timeout` - (Optional) The timeout period for CLB to transmit requests to backend servers. Only HTTP/HTTPS listeners support this parameter. value range: 30-3600.
* `scheduler` - (Optional) The scheduling algorithm of the Listener. Optional choice contains `wrr`, `wlc`, `sh`.
* `security_policy_id` - (Optional) The TLS security policy of the HTTPS listener. Only HTTPS listeners support this parameter. value range: `default_policy`, `tls_cipher_policy_1_0`, `tls_cipher_policy_1_1`, `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`.
* `send_timeout` - (Optional) The timeout period for CLB to send responses to the client. Only HTTP/HTTPS listeners support this parameter. value range: 1-3600.
* `start_port` - (Optional, ForceNew) The start port for full port listening, with a value range of 1-65535. When `port` is 0, this parameter is required.
* `tags` - (Optional) Tags.

The `health_check` object supports the following:

* `domain` - (Optional) The domain of health check.
* `enabled` - (Optional) The enable status of health check function. Optional choice contains `on`, `off`.
* `healthy_threshold` - (Optional) The healthy threshold of health check, default 3, range in 2~10.
* `http_code` - (Optional) The normal http status code of health check, the value can be `http_2xx` or `http_3xx` or `http_4xx` or `http_5xx`.
* `interval` - (Optional) The interval executing health check, default 2, range in 1~300.
* `method` - (Optional) The method of health check, the value can be `GET` or `HEAD`.
* `port` - (Optional) The port for health check, with a value range of 1-65535.
* `timeout` - (Optional) The response timeout of health check, default 2, range in 1~60..
* `udp_expect` - (Optional) The UDP expect of health check. This field must be specified simultaneously with field `udp_request`.
* `udp_request` - (Optional) The UDP request of health check. This field must be specified simultaneously with field `udp_expect`.
* `un_healthy_threshold` - (Optional) The unhealthy threshold of health check, default 3, range in 2~10.
* `uri` - (Optional) The uri of health check.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `listener_id` - The ID of the Listener.


## Import
Listener can be imported using the id, e.g.
```
$ terraform import volcengine_listener.default lsn-273yv0mhs5xj47fap8sehiiso
```

