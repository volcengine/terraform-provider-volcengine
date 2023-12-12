---
subcategory: "EIP"
layout: "volcengine"
page_title: "Volcengine: volcengine_eip_addresses"
sidebar_current: "docs-volcengine-datasource-eip_addresses"
description: |-
  Use this data source to query detailed information of eip addresses
---
# volcengine_eip_addresses
Use this data source to query detailed information of eip addresses
## Example Usage
```hcl
resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByTraffic"
}
data "volcengine_eip_addresses" "foo" {
  ids = [volcengine_eip_address.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `associated_instance_id` - (Optional) An id of associated instance.
* `associated_instance_type` - (Optional) A type of associated instance, the value can be `Nat`, `NetworkInterface`, `ClbInstance`, `AlbInstance`, `HaVip` or `EcsInstance`.
* `eip_addresses` - (Optional) A list of EIP ip address that you want to query.
* `ids` - (Optional) A list of EIP allocation ids.
* `isp` - (Optional) An ISP of EIP Address, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom`.
* `name` - (Optional) A name of EIP.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The ProjectName of EIP.
* `security_protection_enabled` - (Optional) Security protection. The values are as follows: `true`: Enhanced protection type for public IP (can be added to DDoS native protection (Enterprise Edition) instance). `false`: Default protection type for public IP.
* `status` - (Optional) A status of EIP, the value can be `Attaching` or `Detaching` or `Attached` or `Available`.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `addresses` - The collection of EIP addresses.
    * `allocation_id` - The id of the EIP address.
    * `allocation_time` - The allocation time of the EIP.
    * `bandwidth_package_id` - The id of the bandwidth package.
    * `bandwidth` - The peek bandwidth of the EIP.
    * `billing_type` - The billing type of the EIP.
    * `business_status` - The business status of the EIP.
    * `deleted_time` - The deleted time of the EIP.
    * `description` - The description of the EIP.
    * `eip_address` - The EIP ip address of the EIP.
    * `expired_time` - The expired time of the EIP.
    * `id` - The id of the EIP address.
    * `instance_id` - The instance id which be associated to the EIP.
    * `instance_type` - The type of the associated instance.
    * `isp` - The ISP of EIP Address.
    * `lock_reason` - The lock reason of the EIP.
    * `name` - The name of the EIP.
    * `overdue_time` - The overdue time of the EIP.
    * `project_name` - The ProjectName of the EIP.
    * `security_protection_types` - Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced`.
    * `status` - The status of the EIP.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `updated_at` - The last update time of the EIP.
* `total_count` - The total count of EIP addresses query.


