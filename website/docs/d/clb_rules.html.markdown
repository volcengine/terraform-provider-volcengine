---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_clb_rules"
sidebar_current: "docs-volcengine-datasource-clb_rules"
description: |-
  Use this data source to query detailed information of clb rules
---
# volcengine_clb_rules
Use this data source to query detailed information of clb rules
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
resource "volcengine_clb_rule" "foo" {
  listener_id     = volcengine_listener.foo.id
  server_group_id = volcengine_server_group.foo.id
  domain          = "test-volc123.com"
  url             = "/yyyy"
}
data "volcengine_clb_rules" "foo" {
  ids         = [volcengine_clb_rule.foo.id]
  listener_id = volcengine_listener.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `listener_id` - (Required) The Id of listener.
* `ids` - (Optional) A list of Rule IDs.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rules` - The collection of Rule query.
    * `description` - The Description of Rule.
    * `domain` - The Domain of Rule.
    * `id` - The Id of Rule.
    * `rule_id` - The Id of Rule.
    * `server_group_id` - The Id of Server Group.
    * `url` - The Url of Rule.


