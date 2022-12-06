---
subcategory: "NAT"
layout: "volcengine"
page_title: "Volcengine: volcengine_dnat_entry"
sidebar_current: "docs-volcengine-resource-dnat_entry"
description: |-
  Provides a resource to manage dnat entry
---
# volcengine_dnat_entry
Provides a resource to manage dnat entry
## Example Usage
```hcl
resource "volcengine_dnat_entry" "foo" {
  nat_gateway_id  = "ngw-imw3aej7e96o8gbssxkfbybv"
  external_ip     = "10.249.186.68"
  external_port   = "23"
  internal_ip     = "193.168.1.1"
  internal_port   = "24"
  protocol        = "tcp"
  dnat_entry_name = "terraform-test2"
}
```
## Argument Reference
The following arguments are supported:
* `external_ip` - (Required) Provides the public IP address for public network access.
* `external_port` - (Required) The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
* `internal_ip` - (Required) Provides the internal IP address.
* `internal_port` - (Required) The port or port segment on which the cloud server instance provides services to the public network.
* `nat_gateway_id` - (Required, ForceNew) The id of the nat gateway to which the entry belongs.
* `protocol` - (Required) The network protocol.
* `dnat_entry_name` - (Optional) The name of the DNAT rule.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `dnat_entry_id` - The id of the DNAT rule.


## Import
Dnat entry can be imported using the id, e.g.
```
$ terraform import volcengine_dnat_entry.default dnat-3fvhk47kf56****
```

