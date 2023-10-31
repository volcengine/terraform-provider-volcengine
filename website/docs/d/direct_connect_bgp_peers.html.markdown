---
subcategory: "DIRECT_CONNECT"
layout: "volcengine"
page_title: "Volcengine: volcengine_direct_connect_bgp_peers"
sidebar_current: "docs-volcengine-datasource-direct_connect_bgp_peers"
description: |-
  Use this data source to query detailed information of direct connect bgp peers
---
# volcengine_direct_connect_bgp_peers
Use this data source to query detailed information of direct connect bgp peers
## Example Usage
```hcl
data "volcengine_direct_connect_bgp_peers" "foo" {
  ids = ["bgp-171w6pn39ruo04d1w33iq****"]
}
```
## Argument Reference
The following arguments are supported:
* `bgp_peer_name` - (Optional) The name of bgp peer.
* `direct_connect_gateway_id` - (Optional) The id of direct connect gateway.
* `ids` - (Optional) A list of IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `remote_asn` - (Optional) The remote asn of bgp peer.
* `virtual_interface_id` - (Optional) The id of virtual interface.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `bgp_peers` - The collection of query.
    * `account_id` - The id of account.
    * `auth_key` - The key of auth.
    * `bgp_peer_id` - The id of bgp peer.
    * `bgp_peer_name` - The name of bgp peer.
    * `creation_time` - The create time of bgp peer.
    * `description` - The Description of bgp peer.
    * `local_asn` - The local asn of bgp peer.
    * `remote_asn` - The remote asn of bgp peer.
    * `session_status` - The session status of bgp peer.
    * `status` - The status of bgp peer.
    * `update_time` - The update time of bgp peer.
    * `virtual_interface_id` - The id of virtual interface.
* `total_count` - The total count of query.


