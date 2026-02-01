---
subcategory: "DIRECT_CONNECT"
layout: "volcengine"
page_title: "Volcengine: volcengine_direct_connect_virtual_interfaces"
sidebar_current: "docs-volcengine-datasource-direct_connect_virtual_interfaces"
description: |-
  Use this data source to query detailed information of direct connect virtual interfaces
---
# volcengine_direct_connect_virtual_interfaces
Use this data source to query detailed information of direct connect virtual interfaces
## Example Usage
```hcl
data "volcengine_direct_connect_virtual_interfaces" "foo" {
  virtual_interface_name = "tf-test"
}
```
## Argument Reference
The following arguments are supported:
* `direct_connect_connection_id` - (Optional) The direct connect connection ID that associated with this virtual interface.
* `direct_connect_gateway_id` - (Optional) The direct connect gateway ID that associated with this virtual interface.
* `ids` - (Optional) A list of IDs.
* `local_ip` - (Optional) The local IP that associated with this virtual interface.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `peer_ip` - (Optional) The peer IP that associated with this virtual interface.
* `route_type` - (Optional) The route type of virtual interface.
* `tag_filters` - (Optional) The filter tag of direct connect virtual interface.
* `virtual_interface_name` - (Optional) The name of virtual interface.
* `vlan_id` - (Optional) The VLAN ID of virtual interface.

The `tag_filters` object supports the following:

* `key` - (Optional) The tag key of cloud resource instance.
* `value` - (Optional) The tag value of cloud resource instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `virtual_interfaces` - The collection of query.
    * `account_id` - The account ID which this virtual interface belongs.
    * `bandwidth` - The band width limit of virtual interface,in Mbps.
    * `bfd_detect_interval` - The BFD detect interval.
    * `bfd_detect_multiplier` - The BFD detect times.
    * `creation_time` - The creation time of virtual interface.
    * `description` - The description of the virtual interface.
    * `direct_connect_connection_id` - The direct connect connection ID which associated with this virtual interface.
    * `direct_connect_gateway_id` - The direct connect gateway ID which associated with this virtual interface.
    * `enable_bfd` - Whether enable BFD detect.
    * `enable_nqa` - Whether enable NQA detect.
    * `local_ip` - The local IP that associated with this virtual interface.
    * `nqa_detect_interval` - The NQA detect interval.
    * `nqa_detect_multiplier` - The NAQ detect times.
    * `peer_ip` - The peer IP that associated with this virtual interface.
    * `route_type` - The route type of this virtual interface.
    * `status` - The status of virtaul interface.
    * `tags` - The tags that direct connect gateway added.
        * `key` - The tag key.
        * `value` - The tag value.
    * `update_time` - The update time of virtual interface.
    * `virtual_interface_id` - The virtual interface ID.
    * `virtual_interface_name` - The name of virtual interface.
    * `vlan_id` - The VLAN ID of virtual interface.


