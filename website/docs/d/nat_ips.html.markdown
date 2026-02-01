---
subcategory: "NAT"
layout: "volcengine"
page_title: "Volcengine: volcengine_nat_ips"
sidebar_current: "docs-volcengine-datasource-nat_ips"
description: |-
  Use this data source to query detailed information of nat ips
---
# volcengine_nat_ips
Use this data source to query detailed information of nat ips
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

data "volcengine_nat_ips" "foo" {
  nat_gateway_id = volcengine_nat_gateway.intranet_nat_gateway.id
}
```
## Argument Reference
The following arguments are supported:
* `nat_gateway_id` - (Required) The id of the Nat gateway.
* `ids` - (Optional) A list of Nat IP ids.
* `name_regex` - (Optional) The Name Regex of Nat ip.
* `nat_ip_name` - (Optional) The name of the Nat IP.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `nat_ips` - The collection of query.
    * `id` - The id of the Nat Ip.
    * `is_default` - Whether the Ip is the default Nat Ip.
    * `nat_gateway_id` - The id of the Nat gateway.
    * `nat_ip_description` - The description of the Nat Ip.
    * `nat_ip_id` - The id of the Nat Ip.
    * `nat_ip_name` - The name of the Nat Ip.
    * `nat_ip` - The ip address of the Nat Ip.
    * `status` - The status of the Nat Ip.
    * `using_status` - The using status of the Nat Ip.
* `total_count` - The total count of query.


