---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb"
sidebar_current: "docs-volcengine-resource-alb"
description: |-
  Provides a resource to manage alb
---
# volcengine_alb
Provides a resource to manage alb
## Example Usage
```hcl
data "volcengine_alb_zones" "foo" {
}

resource "volcengine_vpc" "vpc_ipv6" {
  vpc_name    = "acc-test-vpc-ipv6"
  cidr_block  = "172.16.0.0/16"
  enable_ipv6 = true
}

resource "volcengine_subnet" "subnet_ipv6_1" {
  subnet_name     = "acc-test-subnet-ipv6-1"
  cidr_block      = "172.16.1.0/24"
  zone_id         = data.volcengine_alb_zones.foo.zones[0].id
  vpc_id          = volcengine_vpc.vpc_ipv6.id
  ipv6_cidr_block = 1
}

resource "volcengine_subnet" "subnet_ipv6_2" {
  subnet_name     = "acc-test-subnet-ipv6-2"
  cidr_block      = "172.16.2.0/24"
  zone_id         = data.volcengine_alb_zones.foo.zones[1].id
  vpc_id          = volcengine_vpc.vpc_ipv6.id
  ipv6_cidr_block = 2
}

resource "volcengine_vpc_ipv6_gateway" "ipv6_gateway" {
  vpc_id = volcengine_vpc.vpc_ipv6.id
  name   = "acc-test-ipv6-gateway"
}

resource "volcengine_alb" "alb-private" {
  address_ip_version = "IPv4"
  type               = "private"
  load_balancer_name = "acc-test-alb-private"
  description        = "acc-test"
  subnet_ids         = [volcengine_subnet.subnet_ipv6_1.id, volcengine_subnet.subnet_ipv6_2.id]
  project_name       = "default"
  delete_protection  = "off"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_alb" "alb-public" {
  address_ip_version             = "DualStack"
  type                           = "public"
  load_balancer_name             = "acc-test-alb-public"
  description                    = "acc-test"
  subnet_ids                     = [volcengine_subnet.subnet_ipv6_1.id, volcengine_subnet.subnet_ipv6_2.id]
  project_name                   = "default"
  delete_protection              = "off"
  modification_protection_status = "NonProtection"
  modification_protection_reason = "Test modification protection"
  load_balancer_edition          = "Basic"

  eip_billing_config {
    isp              = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth        = 1
  }
  ipv6_eip_billing_config {
    isp          = "BGP"
    billing_type = "PostPaidByBandwidth"
    bandwidth    = 1
  }

  tags {
    key   = "k1"
    value = "v1"
  }
  depends_on = [volcengine_vpc_ipv6_gateway.ipv6_gateway]
}

# CLone ALB instance
resource "volcengine_alb" "alb-cloned" {
  source_load_balancer_id = volcengine_alb.alb-private.id
  load_balancer_name      = "acc-test-alb-cloned"
  description             = "cloned from alb-private"
  subnet_ids              = [volcengine_subnet.subnet_ipv6_1.id]
  type                    = "private"
  project_name            = "default"
}

# Example of ALB network type change, private -> public
resource "volcengine_alb" "alb-type-change" {
  load_balancer_name = "acc-test-alb-type-change"
  description        = "will change to public type"
  subnet_ids         = [volcengine_subnet.subnet_ipv6_1.id, volcengine_subnet.subnet_ipv6_2.id]
  type               = "public"
  project_name       = "default"
  allocation_ids     = ["eip-iinpy4k1rytc74o8curgocd7", "eip-iinpy4k1rytc74o8curgocd8"]
}
```
## Argument Reference
The following arguments are supported:
* `subnet_ids` - (Required) The id of the Subnet.
* `type` - (Required) The type of the Alb. Valid values: `public`, `private`.
* `address_ip_version` - (Optional, ForceNew) The address ip version of the Alb. Valid values: `IPv4`, `DualStack`. Default is `ipv4`.
* `allocation_ids` - (Optional) The ID of the public IP. This field is only valid when the type field changes from private to public.
* `delete_protection` - (Optional) Whether to enable the delete protection function of the Alb. Valid values: `on`, `off`. Default is `off`.
* `description` - (Optional) The description of the Alb.
* `eip_billing_config` - (Optional, ForceNew) The billing configuration of the EIP which automatically associated to the Alb. This field is valid when the type of the Alb is `public`.When the type of the Alb is `private`, suggest using a combination of resource `volcengine_eip_address` and `volcengine_eip_associate` to achieve public network access function.
* `global_accelerator` - (Optional) The global accelerator configuration.
* `ipv6_eip_billing_config` - (Optional, ForceNew) The billing configuration of the Ipv6 EIP which automatically associated to the Alb. This field is required when the type of the Alb is `public`.When the type of the Alb is `private`, suggest using a combination of resource `volcengine_vpc_ipv6_gateway` and `volcengine_vpc_ipv6_address_bandwidth` to achieve ipv6 public network access function.
* `load_balancer_edition` - (Optional, ForceNew) The version of the ALB instance. Basic: Basic Edition. Standard: Standard Edition. Default is `Basic`.
* `load_balancer_name` - (Optional) The name of the Alb.
* `modification_protection_reason` - (Optional) The reason for enabling instance modification protection. This parameter is valid when the modification_protection_status is `ConsoleProtection`.
* `modification_protection_status` - (Optional) Whether to enable the modification protection function of the Alb. Valid values: `NonProtection`, `ConsoleProtection`. Default is `NonProtection`. NonProtection: Instance modification protection is not enabled. ConsoleProtection: Instance modification protection is enabled; you cannot modify the instance configuration through the ALB console, and can only modify the instance configuration by calling the API.
* `project_name` - (Optional) The ProjectName of the Alb.
* `proxy_protocol_enabled` - (Optional) ALB can support the Proxy Protocol and record the real IP of the client.
* `source_load_balancer_id` - (Optional, ForceNew) The source ALB instance ID for cloning. If specified, the ALB instance will be cloned from this source.
* `tags` - (Optional) Tags.
* `waf_instance_id` - (Optional) The ID of the WAF instance to be associated with the Alb. This field is valid when the value of the `waf_protection_enabled` is `on`.
* `waf_protected_domain` - (Optional) The domain name of the WAF protected Alb. This field is valid when the value of the `waf_protection_enabled` is `on`.
* `waf_protection_enabled` - (Optional) Whether to enable the WAF protection function of the Alb. Valid values: `off`, `on`. Default is `off`.

The `eip_billing_config` object supports the following:

* `bandwidth` - (Required, ForceNew) The peek bandwidth of the EIP which automatically assigned to the Alb. Unit: Mbps.
* `eip_billing_type` - (Required, ForceNew) The billing type of the EIP which automatically assigned to the Alb. Valid values: `PostPaidByBandwidth`, `PostPaidByTraffic`.
* `isp` - (Required, ForceNew) The ISP of the EIP which automatically associated to the Alb, the value can be `BGP`.

The `global_accelerator` object supports the following:

* `accelerator_id` - (Optional) The global accelerator id.
* `accelerator_listener_id` - (Optional) The global accelerator listener id.
* `endpoint_group_id` - (Optional) The global accelerator endpoint group id.
* `weight` - (Optional) The traffic distribution weight of the endpoint. The value range is: 1 - 100.

The `ipv6_eip_billing_config` object supports the following:

* `bandwidth` - (Required, ForceNew) The peek bandwidth of the Ipv6 EIP which automatically assigned to the Alb. Unit: Mbps.
* `billing_type` - (Required, ForceNew) The billing type of the Tpv6 EIP which automatically assigned to the Alb. Valid values: `PostPaidByBandwidth`, `PostPaidByTraffic`.
* `isp` - (Required, ForceNew) The ISP of the Ipv6 EIP which automatically associated to the Alb, the value can be `BGP`.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `dns_name` - The DNS name.
* `local_addresses` - The local addresses of the Alb.
* `status` - The status of the Alb.
* `vpc_id` - The vpc id of the Alb.
* `zone_mappings` - Configuration information of the Alb instance in different Availability Zones.
    * `load_balancer_addresses` - The IP address information of the Alb in this availability zone.
        * `eip_address` - The Eip address of the Alb in this availability zone.
        * `eip_id` - The Eip id of alb instance in this availability zone.
        * `eni_address` - The Eni address of the Alb in this availability zone.
        * `eni_id` - The Eni id of the Alb in this availability zone.
        * `eni_ipv6_address` - The Eni Ipv6 address of the Alb in this availability zone.
        * `ipv6_eip_id` - The Ipv6 Eip id of alb instance in this availability zone.
    * `subnet_id` - The subnet id of the Alb in this availability zone.
    * `zone_id` - The availability zone id of the Alb.


## Import
Alb can be imported using the id, e.g.
```
$ terraform import volcengine_alb.default resource_id
```

