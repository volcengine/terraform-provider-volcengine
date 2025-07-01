---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_domain"
sidebar_current: "docs-volcengine-resource-waf_domain"
description: |-
  Provides a resource to manage waf domain
---
# volcengine_waf_domain
Provides a resource to manage waf domain
## Example Usage
```hcl
resource "volcengine_waf_domain" "foo" {
  domain      = "www.tf-test.com"
  access_mode = 10
  protocols   = ["HTTP"]
  protocol_ports {
    http = [80]
  }
  enable_ipv6          = 0
  proxy_config         = 1
  keep_alive_time_out  = 100
  keep_alive_request   = 200
  client_max_body_size = 1024
  lb_algorithm         = "wlc"
  public_real_server   = 0
  vpc_id               = "vpc-2d6485y7p95og58ozfcvxxxxx"
  backend_groups {
    access_port = [80]
    backends {
      protocol = "HTTP"
      ip       = "192.168.0.0"
      port     = 80
      weight   = 40
    }
    backends {
      protocol = "HTTP"
      ip       = "192.168.1.0"
      port     = 80
      weight   = 60
    }
    name = "default"
  }
  client_ip_location        = 0
  custom_header             = ["x-top-1", "x-top-2"]
  proxy_connect_time_out    = 10
  proxy_write_time_out      = 120
  proxy_read_time_out       = 200
  proxy_keep_alive          = 101
  proxy_retry               = 10
  proxy_keep_alive_time_out = 20
}
```
## Argument Reference
The following arguments are supported:
* `access_mode` - (Required, ForceNew) Access mode.
* `domain` - (Required, ForceNew) List of domain names that need to be protected by WAF.
* `api_enable` - (Optional) Whether to enable the API protection policy. Works only on modified scenes.
* `auto_cc_enable` - (Optional) Whether to enable the intelligent CC protection strategy. Works only on modified scenes.
* `backend_groups` - (Optional) The configuration of source station.
* `black_ip_enable` - (Optional) Whether to enable the access ban list policy. Works only on modified scenes.
* `black_lct_enable` - (Optional) Whether to enable the geographical location access control policy. Works only on modified scenes.
* `bot_dytoken_enable` - (Optional) Whether to enable the bot dynamic token. Works only on modified scenes.
* `bot_frequency_enable` - (Optional) Whether to enable the bot frequency limit policy. Works only on modified scenes.
* `bot_repeat_enable` - (Optional) Whether to enable the bot frequency limit policy. Works only on modified scenes.
* `bot_sequence_default_action` - (Optional) Set the default actions of the bot behavior map strategy. Works only on modified scenes.
* `bot_sequence_enable` - (Optional) Whether to enable the bot behavior map. Works only on modified scenes.
* `cc_enable` - (Optional) Whether to enable the CC protection policy. Works only on modified scenes.
* `certificate_id` - (Optional) When the protocol type is HTTPS, the bound certificate ID needs to be entered.
* `certificate_platform` - (Optional) Certificate custody platform. Works only on modified scenes.
* `client_ip_location` - (Optional) The method of obtaining the client IP.
* `client_max_body_size` - (Optional) The client requests the maximum value of body.
* `cloud_access_config` - (Optional) Access port information.If AccessMode is Alb/CLB, this field is required.
* `custom_bot_enable` - (Optional) Whether to enable the custom Bot classification strategy. Works only on modified scenes.
* `custom_header` - (Optional) Custom Header.
* `custom_rsp_enable` - (Optional) Whether to enable the custom response interception policy. Works only on modified scenes.
* `custom_sni` - (Optional) Custom SNI needs to be configured when EnableSNI=1. Works only on modified scenes.
* `defence_mode` - (Optional) The protection mode of the instance. Works only on modified scenes.
* `dlp_enable` - (Optional) Whether to activate the strategy for preventing the leakage of sensitive information. Works only on modified scenes.
* `enable_custom_redirect` - (Optional) Whether to enable user-defined redirection. Works only on modified scenes.
* `enable_http2` - (Optional) Whether to enable HTTP 2.0.
* `enable_ipv6` - (Optional) Whether it supports protecting IPv6 requests.
* `enable_sni` - (Optional) Whether to enable the SNI configuration. Works only on modified scenes.
* `extra_defence_mode_lb_instance` - (Optional) The protection mode of the exception instance. It takes effect when the access mode is accessed through an application load balancing (ALB) instance (AccessMode=20). Works only on modified scenes.
* `keep_alive_request` - (Optional) The number of long connection multiplexes.
* `keep_alive_time_out` - (Optional) Long connection retention time.
* `lb_algorithm` - (Optional) The types of load balancing algorithms.
* `llm_available` - (Optional) Is LLM available. Works only on modified scenes.
* `project_name` - (Optional) The name of project. Works only on modified scenes.
* `protocol_follow` - (Optional) Whether to enable protocol following.
* `protocol_ports` - (Optional) Access port information.
* `protocols` - (Optional) Access protocol types.
* `proxy_config` - (Optional) Whether to enable proxy configuration.
* `proxy_connect_time_out` - (Optional) The timeout period for establishing a connection between the WAF and the backend server.
* `proxy_keep_alive_time_out` - (Optional) Idle long connection timeout period.
* `proxy_keep_alive` - (Optional) The number of reusable WAF origin long connections.
* `proxy_read_time_out` - (Optional) The timeout period during which WAF reads the response from the backend server.
* `proxy_retry` - (Optional) The number of retries for WAF back to source.
* `proxy_write_time_out` - (Optional) The timeout period during which the WAF transmits the request to the backend server.
* `public_real_server` - (Optional) Connect to the source return mode.
* `redirect_https` - (Optional) When only the HTTPS protocol is enabled, whether to redirect HTTP requests to HTTPS. Works only on modified scenes.
* `ssl_ciphers` - (Optional) Encryption kit.
* `ssl_protocols` - (Optional) TLS protocol version.
* `system_bot_enable` - (Optional) Whether to enable the managed Bot classification strategy. Works only on modified scenes.
* `tamper_proof_enable` - (Optional) Whether to enable the page tamper-proof policy. Works only on modified scenes.
* `tls_enable` - (Optional) Whether to enable the log service.
* `tls_fields_config` - (Optional) Details of log field configuration. Works only on modified scenes.
* `volc_certificate_id` - (Optional) When the protocol type is HTTPS, the bound certificate ID needs to be entered. Works only on modified scenes.
* `vpc_id` - (Optional) The ID of vpc.
* `waf_enable` - (Optional) Whether to enable the vulnerability protection strategy. Works only on modified scenes.
* `waf_white_req_enable` - (Optional) Whether to enable the whitening strategy for vulnerability protection requests. Works only on modified scenes.
* `white_enable` - (Optional) Whether to enable the access list policy. Works only on modified scenes.
* `white_field_enable` - (Optional) Whether to enable the whitening strategy for vulnerability protection fields. Works only on modified scenes.

The `backend_groups` object supports the following:

* `access_port` - (Optional) Access port number.
* `backends` - (Optional) The details of the source station group.
* `name` - (Optional) Source station group name.

The `backends` object supports the following:

* `ip` - (Optional) Source station IP address.
* `port` - (Optional) Source station port number.
* `protocol` - (Optional) The agreement of Source Station.
* `weight` - (Optional) The weight of the source station rules.

The `cloud_access_config` object supports the following:

* `instance_id` - (Required) The ID of instance.
* `access_protocol` - (Optional) The access protocol needs to be consistent with the monitoring protocol.
* `instance_name` - (Optional) The name of instance. Works only on modified scenes.
* `listener_id` - (Optional) The ID of listener.
* `lost_association_from_alb` - (Optional) Whether the instance is unbound from the alb and is unbound on the ALB side. Works only on modified scenes.
* `port` - (Optional) The port number corresponding to the listener.
* `protocol` - (Optional) The type of Listener protocol.

The `extra_defence_mode_lb_instance` object supports the following:

* `defence_mode` - (Optional) Set the protection mode for exceptional ALB instances. Works only on modified scenes.
* `instance_id` - (Optional) The Id of ALB instance. Works only on modified scenes.

The `headers_config` object supports the following:

* `enable` - (Optional) Whether the log contains this field. Works only on modified scenes.
* `excluded_key_list` - (Optional) For the use of composite fields, exclude the fields in the keyword list from the JSON of the fields. Works only on modified scenes.
* `statistical_key_list` - (Optional) Create statistical indexes for the fields of the list. Works only on modified scenes.

The `protocol_ports` object supports the following:

* `http` - (Optional) Ports supported by the HTTP protocol.
* `https` - (Optional) Ports supported by the HTTPs protocol.

The `tls_fields_config` object supports the following:

* `headers_config` - (Optional) The configuration of Headers. Works only on modified scenes.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `advanced_defense_ip` - High-defense instance IP.
* `advanced_defense_ipv6` - High-defense instance IPv6.
* `attack_status` - The status of the attack.
* `certificate_name` - The name of the certificate.
* `cname` - The CNAME value generated by the WAF instance.
* `defence_mode_computed` - The protection mode of the instance.
* `server_ips` - The IP of the WAF protection instance.
* `src_ips` - WAF source IP.
* `status` - The status of access.
* `update_time` - The update time.


## Import
WafDomain can be imported using the id, e.g.
```
$ terraform import volcengine_waf_domain.default resource_id
```

