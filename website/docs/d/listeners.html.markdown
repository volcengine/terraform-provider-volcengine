---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_listeners"
sidebar_current: "docs-volcengine-datasource-listeners"
description: |-
  Use this data source to query detailed information of listeners
---
# volcengine_listeners
Use this data source to query detailed information of listeners
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


data "volcengine_listeners" "foo" {
  ids = [volcengine_listener.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Listener IDs.
* `listener_name` - (Optional) The name of the Listener.
* `load_balancer_id` - (Optional) The id of the Clb.
* `name_regex` - (Optional) A Name Regex of Listener.
* `output_file` - (Optional) File name where to save data source results.
* `protocol` - (Optional) The protocol of the Listener. Values: `TCP`, `UDP`, `HTTP`, `HTTPS`.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `listeners` - The collection of Listener query.
    * `acl_ids` - The acl ID list to which the Listener is bound.
    * `acl_status` - The acl status of the Listener.
    * `acl_type` - The acl type of the Listener.
    * `bandwidth` - The bandwidth of the Listener. Unit: Mbps.
    * `ca_certificate_id` - The ID of the CA certificate which is associated with the Listener. When `ca_enabled` is `true`, this parameter is returned.
    * `ca_enabled` - Whether to enable CACertificate two-way authentication.
    * `cert_center_certificate_id` - The ID of the certificate in Certificate Center. When `certificate_source` is `cert_center`, this parameter is returned.
    * `certificate_id` - The ID of the certificate which is associated with the Listener.
    * `certificate_source` - The source of the certificate which is associated with the Listener. Values: `clb`, `cert_center`.
    * `client_body_timeout` - The client body timeout of the Listener. Only HTTP/HTTPS listeners return this parameter.
    * `client_header_timeout` - The client header timeout of the Listener. Only HTTP/HTTPS listeners return this parameter.
    * `connection_drain_enabled` - Whether to enable connection drain of the Listener.
    * `connection_drain_timeout` - The connection drain timeout of the Listener.
    * `cookie` - The name of the cookie for session persistence configured on the backend server.
    * `cps` - The maximum number of new connections for Lsistener.
    * `create_time` - The create time of the Listener.
    * `description` - The description of the Listener.
    * `enabled` - The enable status of the Listener.
    * `end_port` - The end port of the Listener. This parameter is returned only when full-port listening is enabled.
    * `established_timeout` - The established timeout of the Listener.
    * `health_check_domain` - The domain of health check.
    * `health_check_enabled` - The enable status of health check function.
    * `health_check_healthy_threshold` - The healthy threshold of health check.
    * `health_check_http_code` - The normal http status code of health check.
    * `health_check_interval` - The interval executing health check.
    * `health_check_method` - The method of health check.
    * `health_check_timeout` - The response timeout of health check.
    * `health_check_udp_expect` - The expected response string for the health check.
    * `health_check_udp_request` - A request string to perform a health check.
    * `health_check_un_healthy_threshold` - The unhealthy threshold of health check.
    * `health_check_uri` - The uri of health check.
    * `helth_check_port` - The backend server port for health checks. When full-port listening is enabled, this parameter is returned to indicate the port used for health checks. When full-port listening is not enabled, this parameter is not returned, and the health check uses the service port of the backend server.
    * `http2_enabled` - Whether the HTTPS protocol listener enables the front-end HTTP 2.0 protocol.
    * `id` - The ID of the Listener.
    * `keepalive_timeout` - The timeout period for the long connection between the client and the CLB. Only HTTP/HTTPS listeners return this parameter.
    * `listener_id` - The ID of the Listener.
    * `listener_name` - The name of the Listener.
    * `load_balancer_id` - The id of the Clb.
    * `max_connections` - The maximum number of connections for the Listener.
    * `persistence_timeout` - The persistence timeout of the Listener.
    * `persistence_type` - The persistence type of the Listener.
    * `port` - The port receiving request of the Listener.
    * `protocol` - The protocol of the Listener.
    * `proxy_connect_timeout` - The timeout period for establishing a connection between the CLB and the backend server. Only HTTP/HTTPS listeners return this parameter.
    * `proxy_protocol_type` - Whether to enable proxy protocol.
    * `proxy_read_timeout` - The timeout period for CLB to read the response from the backend server. Only HTTP/HTTPS listeners return this parameter.
    * `proxy_send_timeout` - The timeout period for CLB to transmit requests to backend servers. Only HTTP/HTTPS listeners return this parameter.
    * `scheduler` - The scheduling algorithm of the Listener. Values: `wrr`, `wlc`, `sh`.
    * `security_policy_id` - The TLS security policy of the HTTPS listener. Only HTTPS listeners return this parameter.
    * `send_timeout` - The timeout period for CLB to send responses to the client. Only HTTP/HTTPS listeners return this parameter.
    * `server_group_id` - The ID of the backend server group which is associated with the Listener.
    * `start_port` - The start port of the Listener. This parameter is returned only when full-port listening is enabled.
    * `status` - The status of the Listener.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `update_time` - The update time of the Listener.
    * `waf_protection_enabled` - Whether to enable WAF protection.
* `total_count` - The total count of Listener query.


