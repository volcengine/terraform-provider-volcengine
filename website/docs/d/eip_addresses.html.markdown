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
data "volcengine_eip_addresses" "default" {
  ids = ["eip-2748mbpjqzhfk7fap8teu0k1a"]
}
```
## Argument Reference
The following arguments are supported:
* `associated_instance_id` - (Optional) An id of associated instance.
* `associated_instance_type` - (Optional) A type of associated instance, the value can be `Nat`, `NetworkInterface`, `ClbInstance` or `EcsInstance`.
* `eip_addresses` - (Optional) A list of EIP ip address that you want to query.
* `ids` - (Optional) A list of EIP allocation ids.
* `isp` - (Optional) An ISP of EIP Address, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom`.
* `name` - (Optional) A name of EIP.
* `output_file` - (Optional) File name where to save data source results.
* `status` - (Optional) A status of EIP, the value can be `Attaching` or `Detaching` or `Attached` or `Available`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `addresses` - The collection of EIP addresses.
    * `allocation_id` - The id of the EIP address.
    * `allocation_time` - The allocation time of the EIP.
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
    * `status` - The status of the EIP.
    * `updated_at` - The last update time of the EIP.
* `total_count` - The total count of EIP addresses query.


