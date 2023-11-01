---
subcategory: "DIRECT_CONNECT"
layout: "volcengine"
page_title: "Volcengine: volcengine_direct_connect_virtual_interface"
sidebar_current: "docs-volcengine-resource-direct_connect_virtual_interface"
description: |-
  Provides a resource to manage direct connect virtual interface
---
# volcengine_direct_connect_virtual_interface
Provides a resource to manage direct connect virtual interface
## Example Usage
```hcl
resource "volcengine_direct_connect_virtual_interface" "foo" {
  virtual_interface_name       = "tf-test-vi"
  description                  = "tf-test"
  direct_connect_connection_id = "dcc-rtkzeotzst1cu3numzi****"
  direct_connect_gateway_id    = "dcg-638x4bjvjawwn3gd5xw****"
  vlan_id                      = 2
  local_ip                     = "**.**.**.**/**"
  peer_ip                      = "**.**.**.**/**"
  route_type                   = "Static"
  enable_bfd                   = false
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `direct_connect_connection_id` - (Required, ForceNew) The direct connect connection ID which associated with.
* `direct_connect_gateway_id` - (Required, ForceNew) The direct connect gateway ID which associated with.
* `local_ip` - (Required, ForceNew) The local IP that associated with.
* `peer_ip` - (Required, ForceNew) The peer IP that associated with.
* `vlan_id` - (Required, ForceNew) The VLAN ID used to connect to the local IDC, please ensure that this VLAN ID is not occupied, the value range: 0 ~ 2999.
* `bandwidth` - (Optional) The band width limit of virtual interface,in Mbps.
* `bfd_detect_interval` - (Optional) The BFD detect interval.
* `bfd_detect_multiplier` - (Optional) The BFD detect times.
* `description` - (Optional) The description of virtual interface.
* `enable_bfd` - (Optional) Whether enable BFD detect.
* `enable_nqa` - (Optional) Whether enable NQA detect.
* `nqa_detect_interval` - (Optional) The NQA detect interval.
* `nqa_detect_multiplier` - (Optional) The NAQ detect times.
* `route_type` - (Optional, ForceNew) The route type of virtual interface,valid value contains `Static`,`BGP`.
* `tags` - (Optional) The tags that direct connect gateway added.
* `virtual_interface_name` - (Optional) The name of virtual interface.

The `tags` object supports the following:

* `key` - (Optional) The tag key.
* `value` - (Optional) The tag value.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
DirectConnectVirtualInterface can be imported using the id, e.g.
```
$ terraform import volcengine_direct_connect_virtual_interface.default resource_id
```

