---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_albs"
sidebar_current: "docs-volcengine-datasource-albs"
description: |-
  Use this data source to query detailed information of albs
---
# volcengine_albs
Use this data source to query detailed information of albs
## Example Usage
```hcl
data "volcengine_alb_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "subnet_1" {
  subnet_name = "acc-test-subnet-1"
  cidr_block  = "172.16.1.0/24"
  zone_id     = data.volcengine_alb_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_subnet" "subnet_2" {
  subnet_name = "acc-test-subnet-2"
  cidr_block  = "172.16.2.0/24"
  zone_id     = data.volcengine_alb_zones.foo.zones[1].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_alb" "foo" {
  address_ip_version = "IPv4"
  type               = "private"
  load_balancer_name = "acc-test-alb-private-${count.index}"
  description        = "acc-test"
  subnet_ids         = [volcengine_subnet.subnet_1.id, volcengine_subnet.subnet_2.id]
  project_name       = "default"
  delete_protection  = "off"
  tags {
    key   = "k1"
    value = "v1"
  }
  count = 3
}

data "volcengine_albs" "foo" {
  ids = volcengine_alb.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `eni_address` - (Optional) The private ip address of the Alb.
* `ids` - (Optional) A list of Alb IDs.
* `load_balancer_name` - (Optional) The name of the Alb.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project` - (Optional) The project of the Alb.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional) The vpc id which Alb belongs to.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `albs` - The collection of query.
    * `access_log` - The access log information of the Alb.
        * `bucket_name` - The bucket name where the logs are stored.
        * `enabled` - Whether the access log function of the Alb is enabled.
    * `address_ip_version` - The address ip version of the Alb, valid value: `IPv4`, `DualStack`.
    * `business_status` - The business status of the Alb, valid value:`Normal`, `FinancialLocked`.
    * `create_time` - The create time of the Alb.
    * `delete_protection` - The deletion protection function of the Alb instance is turned on or off.
    * `deleted_time` - The expected deleted time of the Alb. This parameter has a query value only when the status of the Alb instance is `FinancialLocked`.
    * `description` - The description of the Alb.
    * `dns_name` - The DNS name.
    * `health_log` - The health log information of the Alb.
        * `enabled` - Whether the health log function is enabled.
        * `project_id` - The TLS project id bound to the health check log.
        * `topic_id` - The TLS topic id bound to the health check log.
    * `id` - The ID of the Alb.
    * `listeners` - The listener information of the Alb.
        * `listener_id` - The listener id of the Alb.
        * `listener_name` - The listener name of the Alb.
    * `load_balancer_billing_type` - The billing type of the Alb.
    * `load_balancer_id` - The ID of the Alb.
    * `load_balancer_name` - The name of the Alb.
    * `local_addresses` - The local addresses of the Alb.
    * `lock_reason` - The reason why Alb is locked. This parameter has a query value only when the status of the Alb instance is `FinancialLocked`.
    * `overdue_time` - The overdue time of the Alb. This parameter has a query value only when the status of the Alb instance is `FinancialLocked`.
    * `project_name` - The project name of the Alb.
    * `status` - The status of the Alb.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `tls_access_log` - The tls access log information of the Alb.
        * `enabled` - Whether the tls access log function is enabled.
        * `project_id` - The TLS project id bound to the access log.
        * `topic_id` - The TLS topic id bound to the access log.
    * `type` - The type of the Alb, valid value: `public`, `private`.
    * `update_time` - The update time of the Alb.
    * `vpc_id` - The vpc id of the Alb.
    * `zone_mappings` - Configuration information of the Alb instance in different Availability Zones.
        * `load_balancer_addresses` - The IP address information of the Alb in this availability zone.
            * `eip_address` - The Eip address of the Alb in this availability zone.
            * `eip_id` - The Eip id of alb instance in this availability zone.
            * `eip` - The Eip information of the Alb in this availability zone.
                * `association_mode` - The association mode of the Alb. This parameter has a query value only when the type of the Eip is `anycast`.
                * `bandwidth` - The peek bandwidth of the Eip assigned to Alb. Units: Mbps.
                * `eip_address` - The Eip address of the Alb.
                * `eip_billing_type` - The billing type of the Eip assigned to Alb. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic`.
                * `eip_type` - The Eip type of the Alb.
                * `isp` - The ISP of the Eip assigned to Alb, the value can be `BGP`.
                * `pop_locations` - The pop locations of the Alb. This parameter has a query value only when the type of the Eip is `anycast`.
                    * `pop_id` - The pop id of the Anycast Eip.
                    * `pop_name` - The pop name of the Anycast Eip.
                * `security_protection_types` - The security protection types of the Alb.
            * `eni_address` - The Eni address of the Alb in this availability zone.
            * `eni_id` - The Eni id of the Alb in this availability zone.
            * `eni_ipv6_address` - The Eni Ipv6 address of the Alb in this availability zone.
            * `ipv6_eip_id` - The Ipv6 Eip id of alb instance in this availability zone.
            * `ipv6_eip` - The Ipv6 Eip information of the Alb in this availability zone.
                * `bandwidth` - The peek bandwidth of the Ipv6 Eip assigned to Alb. Units: Mbps.
                * `billing_type` - The billing type of the Ipv6 Eip assigned to Alb. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic`.
                * `isp` - The ISP of the Ipv6 Eip assigned to Alb, the value can be `BGP`.
        * `subnet_id` - The subnet id of the Alb in this availability zone.
        * `zone_id` - The availability zone id of the Alb.
* `total_count` - The total count of query.


