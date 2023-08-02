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

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `listeners` - The collection of Listener query.
    * `acl_ids` - The acl ID list to which the Listener is bound.
    * `acl_status` - The acl status of the Listener.
    * `acl_type` - The acl type of the Listener.
    * `certificate_id` - The ID of the certificate which is associated with the Listener.
    * `create_time` - The create time of the Listener.
    * `enabled` - The enable status of the Listener.
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
    * `id` - The ID of the Listener.
    * `listener_id` - The ID of the Listener.
    * `listener_name` - The name of the Listener.
    * `port` - The port receiving request of the Listener.
    * `protocol` - The protocol of the Listener.
    * `server_group_id` - The ID of the backend server group which is associated with the Listener.
    * `status` - The status of the Listener.
    * `update_time` - The update time of the Listener.
* `total_count` - The total count of Listener query.


