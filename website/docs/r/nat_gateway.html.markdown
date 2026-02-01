---
subcategory: "NAT"
layout: "volcengine"
page_title: "Volcengine: volcengine_nat_gateway"
sidebar_current: "docs-volcengine-resource-nat_gateway"
description: |-
  Provides a resource to manage nat gateway
---
# volcengine_nat_gateway
Provides a resource to manage nat gateway
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

// create internet nat gateway and snat entry and dnat entry
resource "volcengine_nat_gateway" "internet_nat_gateway" {
  vpc_id           = volcengine_vpc.foo.id
  subnet_id        = volcengine_subnet.foo.id
  spec             = "Small"
  nat_gateway_name = "acc-test-internet_ng"
  description      = "acc-test"
  billing_type     = "PostPaid"
  project_name     = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_eip_address" "foo" {
  name         = "acc-test-eip"
  description  = "acc-test"
  bandwidth    = 1
  billing_type = "PostPaidByBandwidth"
  isp          = "BGP"
}

resource "volcengine_eip_associate" "foo" {
  allocation_id = volcengine_eip_address.foo.id
  instance_id   = volcengine_nat_gateway.internet_nat_gateway.id
  instance_type = "Nat"
}

resource "volcengine_snat_entry" "foo" {
  snat_entry_name = "acc-test-snat-entry"
  nat_gateway_id  = volcengine_nat_gateway.internet_nat_gateway.id
  eip_id          = volcengine_eip_address.foo.id
  source_cidr     = "172.16.0.0/24"
  depends_on      = [volcengine_eip_associate.foo]
}

resource "volcengine_dnat_entry" "foo" {
  dnat_entry_name = "acc-test-dnat-entry"
  external_ip     = volcengine_eip_address.foo.eip_address
  external_port   = 80
  internal_ip     = "172.16.0.10"
  internal_port   = 80
  nat_gateway_id  = volcengine_nat_gateway.internet_nat_gateway.id
  protocol        = "tcp"
  depends_on      = [volcengine_eip_associate.foo]
}

// create intranet nat gateway and snat entry and dnat entry
resource "volcengine_nat_gateway" "intranet_nat_gateway" {
  vpc_id           = volcengine_vpc.foo.id
  subnet_id        = volcengine_subnet.foo.id
  nat_gateway_name = "acc-test-intranet_ng"
  description      = "acc-test"
  network_type     = "intranet"
  billing_type     = "PostPaidByUsage"
  project_name     = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_nat_ip" "foo" {
  nat_gateway_id     = volcengine_nat_gateway.intranet_nat_gateway.id
  nat_ip_name        = "acc-test-nat-ip"
  nat_ip_description = "acc-test"
  nat_ip             = "172.16.0.3"
}

resource "volcengine_snat_entry" "foo-intranet" {
  snat_entry_name = "acc-test-snat-entry-intranet"
  nat_gateway_id  = volcengine_nat_gateway.intranet_nat_gateway.id
  nat_ip_id       = volcengine_nat_ip.foo.id
  source_cidr     = "172.16.0.0/24"
}

resource "volcengine_dnat_entry" "foo-intranet" {
  nat_gateway_id  = volcengine_nat_gateway.intranet_nat_gateway.id
  dnat_entry_name = "acc-test-dnat-entry-intranet"
  protocol        = "tcp"
  internal_ip     = "172.16.0.5"
  internal_port   = "82"
  external_ip     = volcengine_nat_ip.foo.nat_ip
  external_port   = "87"
}
```
## Argument Reference
The following arguments are supported:
* `subnet_id` - (Required, ForceNew) The ID of the Subnet.
* `vpc_id` - (Required, ForceNew) The ID of the VPC.
* `billing_type` - (Optional, ForceNew) The billing type of the NatGateway, the value is `PostPaid` or `PrePaid` or `PostPaidByUsage`. Default value is `PostPaid`.
When the `network_type` is `intranet`, the billing type must be `PostPaidByUsage`.
* `description` - (Optional) The description of the NatGateway.
* `nat_gateway_name` - (Optional) The name of the NatGateway.
* `network_type` - (Optional, ForceNew) The network type of the NatGateway. Valid values are `internet` and `intranet`. Default value is `internet`.
* `period` - (Optional, ForceNew) The period of the NatGateway, the valid value range in 1~9 or 12 or 24 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid NatGateway. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `project_name` - (Optional) The ProjectName of the NatGateway.
* `spec` - (Optional) The specification of the NatGateway. Optional choice contains `Small`(default), `Medium`, `Large` or leave blank.
When the `billing_type` is `PostPaidByUsage`, this field should not be specified.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
NatGateway can be imported using the id, e.g.
```
$ terraform import volcengine_nat_gateway.default ngw-vv3t043k05sm****
```

