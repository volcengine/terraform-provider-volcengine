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
* `address_ip_version` - (Optional) The address IP version of the CLB.
* `eip_address` - (Optional) The public ip address of the Clb.
* `eni_address` - (Optional) The private ip address of the Clb.
* `ids` - (Optional) A list of Clb IDs.
* `instance_ids` - (Optional) The IDs of the backend server of the CLB.
* `instance_ips` - (Optional) The IP address of the backend server of the CLB.
* `load_balancer_name` - (Optional) The name of the Clb.
* `master_zone_id` - (Optional) The master zone ID of the CLB.
* `name_regex` - (Optional) A Name Regex of Clb.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The ProjectName of Clb.
* `status` - (Optional) The status of the CLB.
* `tags` - (Optional) Tags.
* `type` - (Optional) The network type of the CLB.
* `vpc_id` - (Optional) The id of the VPC.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `clbs` - The collection of Clb query.
    * `access_log` - The access log configuration of the CLB instance.
        * `bucket_name` - The name of the bucket to which the access logs are delivered.
        * `enabled` - Whether to enable the function of delivering access logs (Layer 7) to Object Storage TOS.
        * `tls_enabled` - Whether to enable the function of delivering access logs (layer 7) to the log service TLS.
        * `tls_project_id` - The project ID of the log service TLS.
        * `tls_topic_id` - The topic ID of the log service TLS.
    * `address_ip_version` - The address ip version of the Clb.
    * `billing_type` - The billing type of the CLB instance.
    * `business_status` - The business status of the Clb.
    * `bypass_security_group_enabled` - Whether the CLB instance has enabled the "Allow Backend Security Groups" function.
    * `create_time` - The create time of the Clb.
    * `deleted_time` - The expected recycle time of the Clb.
    * `description` - The description of the Clb.
    * `eip_address` - The Eip address of the Clb.
    * `eip_billing_config` - The eip billing config of the Clb.
        * `bandwidth_package_id` - The bandwidth package id of the EIP assigned to CLB.
        * `bandwidth` - The peek bandwidth of the EIP assigned to CLB. Units: Mbps.
        * `eip_address` - The public IP address of the CLB instance.
        * `eip_billing_type` - The billing type of the EIP assigned to CLB. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
        * `isp` - The ISP of the EIP assigned to CLB, the value can be `BGP`.
        * `security_protection_types` - The security protection types of the EIP assigned to CLB.
    * `eip_id` - The Eip ID of the Clb.
    * `enabled` - Whether the CLB instance is enabled.
    * `eni_address_num` - The ENI address num of the CLB.
    * `eni_address` - The Eni address of the Clb.
    * `eni_addresses` - The ENI addresses of the CLB.
        * `eip_address` - The public IPv4 address bound to the private IPv4 address.
        * `eip_id` - The eip ID of the public IP bound to the private IPv4 address.
        * `eni_address` - The private IPv4 address of the CLB instance.
    * `eni_id` - The Eni ID of the Clb.
    * `eni_ipv6_address` - The eni ipv6 address of the Clb.
    * `exclusive_cluster_id` - The ID of the exclusive cluster to which the CLB instance belongs.
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
    * `listeners` - The information of the listeners in the CLB instance.
        * `listener_id` - The ID of the Listener.
        * `listener_name` - The name of the Listener.
    * `load_balancer_billing_type` - The billing type of the Clb.
    * `load_balancer_id` - The ID of the Clb.
    * `load_balancer_name` - The name of the Clb.
    * `load_balancer_spec` - The specifications of the Clb.
    * `lock_reason` - The reason why Clb is locked.
    * `log_topic_id` - The log topic ID of the Clb.
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
    * `server_groups` - The information of the server groups in the CLB instance.
        * `server_group_id` - The ID of the server group.
        * `server_group_name` - The name of the server group.
    * `service_managed` - Whether the CLB instance is a managed resource.
    * `slave_zone_id` - The slave zone ID of the CLB.
    * `status` - The status of the Clb.
    * `subnet_id` - The subnet ID of the Clb.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `timestamp_remove_enabled` - Whether to enable the function of clearing the timestamp of TCP/HTTP/HTTPS packets (i.e., time stamp).
    * `type` - The type of the Clb.
    * `update_time` - The update time of the Clb.
    * `vpc_id` - The vpc ID of the Clb.
* `total_count` - The total count of Clb query.


