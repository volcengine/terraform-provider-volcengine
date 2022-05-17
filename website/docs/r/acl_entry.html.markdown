---
subcategory: "CLB"
layout: "vestack"
page_title: "Vestack: vestack_acl_entry"
sidebar_current: "docs-vestack-resource-acl_entry"
description: |-
  Provides a resource to manage acl entry
---
# vestack_acl_entry
Provides a resource to manage acl entry
## Example Usage
```hcl
resource "vestack_acl" "foo" {
  acl_name    = "tf-test-3"
  description = "tf-test"
}

resource "vestack_acl_entry" "foo" {
  acl_id      = vestack_acl.foo.id
  description = "tf acl entry desc demo"
  entry       = "192.2.2.1/32"
}
```
## Argument Reference
The following arguments are supported:
* `acl_id` - (Required, ForceNew) The ID of Acl.
* `entry` - (Required, ForceNew) The content of the AclEntry.
* `description` - (Optional, ForceNew) The description of the AclEntry.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
AclEntry can be imported using the id, e.g.
```
$ terraform import vestack_acl_entry.default ID is a string concatenated with colons(AclId:Entry)
```

