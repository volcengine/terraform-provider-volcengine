---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_security_group"
sidebar_current: "docs-volcengine-resource-privatelink_security_group"
description: |-
  Provides a resource to manage privatelink security group
---
# volcengine_privatelink_security_group
Provides a resource to manage privatelink security group
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

resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-security-group"
  vpc_id              = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo1" {
  security_group_name = "acc-test-security-group-new"
  vpc_id              = volcengine_vpc.foo.id
}

resource "volcengine_clb" "foo" {
  type                       = "public"
  subnet_id                  = volcengine_subnet.foo.id
  load_balancer_spec         = "small_1"
  description                = "acc-test-demo"
  load_balancer_name         = "acc-test-clb"
  load_balancer_billing_type = "PostPaid"
  eip_billing_config {
    isp              = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth        = 1
  }
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_privatelink_vpc_endpoint_service" "foo" {
  resources {
    resource_id   = volcengine_clb.foo.id
    resource_type = "CLB"
  }
  description         = "acc-test"
  auto_accept_enabled = true
}

resource "volcengine_privatelink_vpc_endpoint" "foo" {
  security_group_ids = [volcengine_security_group.foo.id]
  service_id         = volcengine_privatelink_vpc_endpoint_service.foo.id
  endpoint_name      = "acc-test-ep"
  description        = "acc-test"
  lifecycle {
    ignore_changes = [security_group_ids]
  }
}

resource "volcengine_privatelink_security_group" "foo" {
  endpoint_id       = volcengine_privatelink_vpc_endpoint.foo.id
  security_group_id = volcengine_security_group.foo1.id
}
```
## Argument Reference
The following arguments are supported:
* `endpoint_id` - (Required, ForceNew) The id of the endpoint.
* `security_group_id` - (Required, ForceNew) The id of the security group. It is not recommended to use this resource for binding security groups, it is recommended to use the `security_group_id` field of `volcengine_privatelink_vpc_endpoint` for binding.
If using this resource and `volcengine_privatelink_vpc_endpoint` jointly for operations, use lifecycle ignore_changes to suppress changes to the `security_group_id` field in `volcengine_privatelink_vpc_endpoint`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
PrivateLink Security Group Service can be imported using the endpoint id and security group id, e.g.
```
$ terraform import volcengine_privatelink_security_group.default ep-2fe630gurkl37k5gfuy33****:sg-xxxxx
```

