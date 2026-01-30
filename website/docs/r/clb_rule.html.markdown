---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_clb_rule"
sidebar_current: "docs-volcengine-resource-clb_rule"
description: |-
  Provides a resource to manage clb rule
---
# volcengine_clb_rule
Provides a resource to manage clb rule
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
  url             = "/tftest"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_clb_rule" "foo_redirect" {
  listener_id = volcengine_listener.foo.id
  action_type = "Redirect"
  description = "Redirect rule"
  domain      = "example1.com"
  redirect_config {
    protocol    = "HTTP"
    host        = "example3.com"
    path        = "/test"
    port        = "443"
    status_code = "301"
  }
}
```
## Argument Reference
The following arguments are supported:
* `listener_id` - (Required, ForceNew) The ID of listener.
* `action_type` - (Optional) The action type of Rule, valid values: `Forward`, `Redirect`.
* `description` - (Optional) The description of the Rule.
* `domain` - (Optional, ForceNew) The domain of Rule.
* `redirect_config` - (Optional) The redirect configuration. Required when action_type is `Redirect`.
* `server_group_id` - (Optional) Server Group Id. Required when action_type is Forward.
* `tags` - (Optional) Tags.
* `url` - (Optional, ForceNew) The Url of Rule.

The `redirect_config` object supports the following:

* `host` - (Optional) The redirect host, i.e. the domain name redirected by the rule.
* `path` - (Optional) The redirect path.
* `port` - (Optional) The redirect port, valid range: 1~65535.
* `protocol` - (Optional) The redirect protocol. Valid values: `HTTP`, `HTTPS`.
* `status_code` - (Optional) The redirect status code. Valid values: 301, 302, 307, 308.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Rule can be imported using the id, e.g.
Notice: resourceId is ruleId, due to the lack of describeRuleAttributes in openapi, for import resources, please use ruleId:listenerId to import.
we will fix this problem later.
```
$ terraform import volcengine_clb_rule.foo rule-273zb9hzi1gqo7fap8u1k3utb:lsn-273ywvnmiu70g7fap8u2xzg9d
```

