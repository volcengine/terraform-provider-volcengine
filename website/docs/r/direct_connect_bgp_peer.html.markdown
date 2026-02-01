---
subcategory: "DIRECT_CONNECT"
layout: "volcengine"
page_title: "Volcengine: volcengine_direct_connect_bgp_peer"
sidebar_current: "docs-volcengine-resource-direct_connect_bgp_peer"
description: |-
  Provides a resource to manage direct connect bgp peer
---
# volcengine_direct_connect_bgp_peer
Provides a resource to manage direct connect bgp peer
## Example Usage
```hcl
resource "volcengine_direct_connect_bgp_peer" "foo" {
  virtual_interface_id = "dcv-62vi13v131tsn3gd6il****"
  remote_asn           = 2000
  description          = "tf-test"
}
```
## Argument Reference
The following arguments are supported:
* `remote_asn` - (Required, ForceNew) The remote asn of bgp peer.
* `virtual_interface_id` - (Required, ForceNew) The id of virtual interface.
* `auth_key` - (Optional, ForceNew) The auth key of bgp peer.
* `bgp_peer_name` - (Optional) The name of bgp peer.
* `description` - (Optional) The description of bgp peer.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The id of account.
* `bgp_peer_id` - The id of bgp peer.
* `creation_time` - The create time of bgp peer.
* `local_asn` - The local asn of bgp peer.
* `session_status` - The session status of bgp peer.
* `status` - The status of bgp peer.
* `update_time` - The update time of bgp peer.


## Import
DirectConnectBgpPeer can be imported using the id, e.g.
```
$ terraform import volcengine_direct_connect_bgp_peer.default bgp-2752hz4teko3k7fap8u4c****
```

