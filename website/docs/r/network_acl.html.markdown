---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_network_acl"
sidebar_current: "docs-volcengine-resource-network_acl"
description: |-
  Provides a resource to manage network acl
---
# volcengine_network_acl
Provides a resource to manage network acl
## Example Usage
```hcl
resource "volcengine_network_acl" "foo" {
  vpc_id           = "vpc-12bk4qjc69reo17q7y36shv6z"
  network_acl_name = "tf-test-acl"
  ingress_acl_entries {
    network_acl_entry_name = "ingress1"
    policy                 = "accept"
    protocol               = "all"
    source_cidr_ip         = "192.168.0.0/24"
  }
  egress_acl_entries {
    network_acl_entry_name = "egress2"
    policy                 = "accept"
    protocol               = "all"
    destination_cidr_ip    = "192.168.0.0/16"
  }
  project_name = "default"
}
```
## Argument Reference
The following arguments are supported:
* `vpc_id` - (Required, ForceNew) The vpc id of Network Acl.
* `description` - (Optional) The description of the Network Acl.
* `egress_acl_entries` - (Optional) The egress entries of Network Acl.
* `ingress_acl_entries` - (Optional) The ingress entries of Network Acl.
* `network_acl_name` - (Optional) The name of Network Acl.
* `project_name` - (Optional) The project name of the network acl.

The `egress_acl_entries` object supports the following:

* `description` - (Optional) The description of entry.
* `destination_cidr_ip` - (Optional) The DestinationCidrIp of entry.
* `network_acl_entry_name` - (Optional) The name of entry.
* `policy` - (Optional) The policy of entry.
* `port` - (Optional) The port of entry.
* `protocol` - (Optional) The protocol of entry.

The `ingress_acl_entries` object supports the following:

* `description` - (Optional) The description of entry.
* `network_acl_entry_name` - (Optional) The name of entry.
* `policy` - (Optional) The policy of entry.
* `port` - (Optional) The port of entry.
* `protocol` - (Optional) The protocol of entry.
* `source_cidr_ip` - (Optional) The SourceCidrIp of entry.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Network Acl can be imported using the id, e.g.
```
$ terraform import volcengine_network_acl.default nacl-172leak37mi9s4d1w33pswqkh
```

