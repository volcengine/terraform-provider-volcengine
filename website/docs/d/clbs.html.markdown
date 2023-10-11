---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_clbs"
sidebar_current: "docs-volcengine-datasource-clbs"
description: |-
  Use this data source to query detailed information of clbs
---
# volcengine_clbs
Use this data source to query detailed information of clbs
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
  load_balancer_name         = "acc-test-clb-${count.index}"
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
  count = 3
}

data "volcengine_clbs" "foo" {
  ids = volcengine_clb.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `eni_address` - (Optional) The private ip address of the Clb.
* `ids` - (Optional) A list of Clb IDs.
* `load_balancer_name` - (Optional) The name of the Clb.
* `name_regex` - (Optional) A Name Regex of Clb.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The ProjectName of Clb.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional) The id of the VPC.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `clbs` - The collection of Clb query.
    * `address_ip_version` - The address ip version of the Clb.
    * `business_status` - The business status of the Clb.
    * `create_time` - The create time of the Clb.
    * `deleted_time` - The expected recycle time of the Clb.
    * `description` - The description of the Clb.
    * `eip_address` - The Eip address of the Clb.
    * `eip_billing_config` - The eip billing config of the Clb.
        * `bandwidth` - The peek bandwidth of the EIP assigned to CLB. Units: Mbps.
        * `eip_billing_type` - The billing type of the EIP assigned to CLB. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
        * `isp` - The ISP of the EIP assigned to CLB, the value can be `BGP`.
    * `eip_id` - The Eip ID of the Clb.
    * `eni_address` - The Eni address of the Clb.
    * `eni_id` - The Eni ID of the Clb.
    * `eni_ipv6_address` - The eni ipv6 address of the Clb.
    * `expired_time` - The expired time of the CLB.
    * `id` - The ID of the Clb.
    * `instance_status` - The billing status of the CLB.
    * `ipv6_address_bandwidth` - The ipv6 address bandwidth information of the Clb.
        * `bandwidth_package_id` - The bandwidth package id of the Ipv6 EIP assigned to CLB.
        * `bandwidth` - The peek bandwidth of the Ipv6 EIP assigned to CLB. Units: Mbps.
        * `billing_type` - The billing type of the Ipv6 EIP assigned to CLB. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic`.
        * `isp` - The ISP of the Ipv6 EIP assigned to CLB, the value can be `BGP`.
        * `network_type` - The network type of the CLB Ipv6 address.
    * `ipv6_eip_id` - The Ipv6 Eip ID of the Clb.
    * `load_balancer_billing_type` - The billing type of the Clb.
    * `load_balancer_id` - The ID of the Clb.
    * `load_balancer_name` - The name of the Clb.
    * `load_balancer_spec` - The specifications of the Clb.
    * `lock_reason` - The reason why Clb is locked.
    * `master_zone_id` - The master zone ID of the CLB.
    * `modification_protection_reason` - The modification protection reason of the Clb.
    * `modification_protection_status` - The modification protection status of the Clb.
    * `overdue_reclaim_time` - The over reclaim time of the CLB.
    * `overdue_time` - The overdue time of the Clb.
    * `project_name` - The ProjectName of the Clb.
    * `reclaim_time` - The reclaim time of the CLB.
    * `remain_renew_times` - The remain renew times of the CLB. When the value of the renew_type is `AutoRenew`, the query returns this field.
    * `renew_period_times` - The renew period times of the CLB. When the value of the renew_type is `AutoRenew`, the query returns this field.
    * `renew_type` - The renew type of the CLB. When the value of the load_balancer_billing_type is `PrePaid`, the query returns this field.
    * `slave_zone_id` - The slave zone ID of the CLB.
    * `status` - The status of the Clb.
    * `subnet_id` - The subnet ID of the Clb.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `type` - The type of the Clb.
    * `update_time` - The update time of the Clb.
    * `vpc_id` - The vpc ID of the Clb.
* `total_count` - The total count of Clb query.


