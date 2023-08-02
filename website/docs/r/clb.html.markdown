---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_clb"
sidebar_current: "docs-volcengine-resource-clb"
description: |-
  Provides a resource to manage clb
---
# volcengine_clb
Provides a resource to manage clb
## Notice
When Destroy this resource,If the resource charge type is PrePaid,Please unsubscribe the resource 
in  [Volcengine Console](https://console.volcengine.com/finance/unsubscribe/),when complete console operation,yon can
use 'terraform state rm ${resourceId}' to remove.
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

resource "volcengine_clb" "public_clb" {
  type               = "public"
  subnet_id          = volcengine_subnet.foo.id
  load_balancer_name = "acc-test-clb-public"
  load_balancer_spec = "small_1"
  description        = "acc-test-demo"
  project_name       = "default"
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

resource "volcengine_clb" "private_clb" {
  type               = "private"
  subnet_id          = volcengine_subnet.foo.id
  load_balancer_name = "acc-test-clb-private"
  load_balancer_spec = "small_1"
  description        = "acc-test-demo"
  project_name       = "default"
}

resource "volcengine_eip_address" "eip" {
  billing_type = "PostPaidByBandwidth"
  bandwidth    = 1
  isp          = "BGP"
  name         = "tf-eip"
  description  = "tf-test"
  project_name = "default"
}

resource "volcengine_eip_associate" "associate" {
  allocation_id = volcengine_eip_address.eip.id
  instance_id   = volcengine_clb.private_clb.id
  instance_type = "ClbInstance"
}
```
## Argument Reference
The following arguments are supported:
* `load_balancer_spec` - (Required) The specification of the CLB, the value can be `small_1`, `small_2`, `medium_1`, `medium_2`, `large_1`, `large_2`.
* `subnet_id` - (Required, ForceNew) The id of the Subnet.
* `type` - (Required, ForceNew) The type of the CLB. And optional choice contains `public` or `private`.
* `description` - (Optional) The description of the CLB.
* `eip_billing_config` - (Optional, ForceNew) The billing configuration of the EIP which automatically associated to CLB. This field is valid when the type of CLB is `public`.When the type of the CLB is `private`, suggest using a combination of resource `volcengine_eip_address` and `volcengine_eip_associate` to achieve public network access function.
* `eni_address` - (Optional, ForceNew) The eni address of the CLB.
* `load_balancer_billing_type` - (Optional) The billing type of the CLB, the value can be `PostPaid` or `PrePaid`.
* `load_balancer_name` - (Optional) The name of the CLB.
* `master_zone_id` - (Optional) The master zone ID of the CLB.
* `modification_protection_reason` - (Optional) The reason of the console modification protection.
* `modification_protection_status` - (Optional) The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
* `period` - (Optional) The period of the NatGateway, the valid value range in 1~9 or 12 or 24 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid NatGateway. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `project_name` - (Optional) The ProjectName of the CLB.
* `region_id` - (Optional, ForceNew) The region of the request.
* `slave_zone_id` - (Optional) The slave zone ID of the CLB.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional, ForceNew) The id of the VPC.

The `eip_billing_config` object supports the following:

* `eip_billing_type` - (Required, ForceNew) The billing type of the EIP which automatically assigned to CLB. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.When creating a `PrePaid` public CLB, this field must be specified as `PrePaid` simultaneously.When the LoadBalancerBillingType changes from `PostPaid` to `PrePaid`, please manually modify the value of this field to `PrePaid` simultaneously.
* `isp` - (Required, ForceNew) The ISP of the EIP which automatically associated to CLB, the value can be `BGP`.
* `bandwidth` - (Optional) The peek bandwidth of the EIP which automatically assigned to CLB. The value range in 1~500 for PostPaidByBandwidth, and 1~200 for PostPaidByTraffic.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `eip_address` - The Eip address of the Clb.
* `eip_id` - The Eip ID of the Clb.
* `renew_type` - The renew type of the CLB. When the value of the load_balancer_billing_type is `PrePaid`, the query returns this field.


## Import
CLB can be imported using the id, e.g.
```
$ terraform import volcengine_clb.default clb-273y2ok6ets007fap8txvf6us
```

