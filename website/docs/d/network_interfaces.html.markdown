---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_network_interfaces"
sidebar_current: "docs-volcengine-datasource-network_interfaces"
description: |-
  Use this data source to query detailed information of network interfaces
---
# volcengine_network_interfaces
Use this data source to query detailed information of network interfaces
## Example Usage
```hcl
data "volcengine_network_interfaces" "default" {
  ids = ["eni-2744htx2w0j5s7fap8t3ivwze"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of ENI ids.
* `instance_id` - (Optional) An id of the instance to which the ENI is bound.
* `network_interface_ids` - (Optional) A list of network interface ids.
* `network_interface_name` - (Optional) A name of ENI.
* `output_file` - (Optional) File name where to save data source results.
* `primary_ip_addresses` - (Optional) A list of primary IP address of ENI.
* `private_ip_addresses` - (Optional) A list of private IP addresses.
* `project_name` - (Optional) The ProjectName of the ENI.
* `security_group_id` - (Optional) An id of the security group to which the secondary ENI belongs.
* `status` - (Optional) A status of ENI, Optional choice contains `Creating`, `Available`, `Attaching`, `InUse`, `Detaching`, `Deleting`.
* `subnet_id` - (Optional) An id of the subnet to which the ENI is connected.
* `tags` - (Optional) Tags.
* `type` - (Optional) A type of ENI, Optional choice contains `primary`, `secondary`.
* `vpc_id` - (Optional) An id of the virtual private cloud (VPC) to which the ENI belongs.
* `zone_id` - (Optional) The zone ID.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `network_interfaces` - The collection of ENI.
    * `account_id` - The account id of the ENI creator.
    * `associated_elastic_ip_address` - The IP address of the EIP to which the ENI associates.
    * `associated_elastic_ip_id` - The allocation id of the EIP to which the ENI associates.
    * `created_at` - The create time of the ENI.
    * `description` - The description of the ENI.
    * `device_id` - The id of the device to which the ENI is bound.
    * `id` - The id of the ENI.
    * `mac_address` - The mac address of the ENI.
    * `network_interface_id` - The id of the ENI.
    * `network_interface_name` - The name of the ENI.
    * `port_security_enabled` - The enable of port security.
    * `primary_ip_address` - The primary IP address of the ENI.
    * `private_ip_sets` - The IP address of secondary private network interface.
        * `associated_elastic_ip` - The public IP that secondary private network IP associated with.
            * `allocation_id` - The public IP ID.
            * `eip_address` - The public IP address.
        * `primary` - Whether the network interface is primary IP address.
        * `private_ip_address` - The secondary private network IP address of the network interface card.
    * `project_name` - The ProjectName of the ENI.
    * `security_group_ids` - The list of the security group id to which the secondary ENI belongs.
    * `service_managed` - Whether the network card has been authorized to be used by other account services.
    * `status` - The status of the ENI.
    * `subnet_id` - The id of the subnet to which the ENI is connected.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `type` - The type of the ENI.
    * `updated_at` - The last update time of the ENI.
    * `vpc_id` - The id of the virtual private cloud (VPC) to which the ENI belongs.
    * `vpc_name` - The name of the virtual private cloud (VPC) to which the ENI belongs.
    * `zone_id` - The zone id of the ENI.
* `total_count` - The total count of ENI query.


