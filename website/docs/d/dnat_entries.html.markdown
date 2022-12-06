---
subcategory: "NAT"
layout: "volcengine"
page_title: "Volcengine: volcengine_dnat_entries"
sidebar_current: "docs-volcengine-datasource-dnat_entries"
description: |-
  Use this data source to query detailed information of dnat entries
---
# volcengine_dnat_entries
Use this data source to query detailed information of dnat entries
## Example Usage
```hcl
data "volcengine_dnat_entries" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `dnat_entry_name` - (Optional) The name of the DNAT entry.
* `external_ip` - (Optional) Provides the public IP address for public network access.
* `external_port` - (Optional) The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
* `ids` - (Optional) A list of DNAT entry ids.
* `internal_ip` - (Optional) Provides the internal IP address.
* `internal_port` - (Optional) The port or port segment on which the cloud server instance provides services to the public network.
* `nat_gateway_id` - (Optional) The id of the NAT gateway.
* `output_file` - (Optional) File name where to save data source results.
* `protocol` - (Optional) The network protocol.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `dnat_entries` - List of DNAT entries.
    * `dnat_entry_id` - The ID of the DNAT entry.
    * `dnat_entry_name` - The name of the DNAT entry.
    * `external_ip` - Provides the public IP address for public network access.
    * `external_port` - The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
    * `internal_ip` - Provides the internal IP address.
    * `internal_port` - The port or port segment on which the cloud server instance provides services to the public network.
    * `nat_gateway_id` - The ID of the NAT gateway.
    * `protocol` - The network protocol.
    * `status` - The network status.
* `total_count` - The total count of snat entries query.


