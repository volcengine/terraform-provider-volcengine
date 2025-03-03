---
subcategory: "CLOUD_FIREWALL"
layout: "volcengine"
page_title: "Volcengine: volcengine_cfw_address_book"
sidebar_current: "docs-volcengine-resource-cfw_address_book"
description: |-
  Provides a resource to manage cfw address book
---
# volcengine_cfw_address_book
Provides a resource to manage cfw address book
## Example Usage
```hcl
resource "volcengine_cfw_address_book" "foo" {
  group_name   = "acc-test-address-book"
  description  = "acc-test"
  group_type   = "ip"
  address_list = ["192.168.1.1", "192.168.2.2"]
}
```
## Argument Reference
The following arguments are supported:
* `address_list` - (Required) The address list of the address book.
 When group_type is `ip`, fill in IPv4/CIDRV4 addresses in the address list.
 When group_type is `port`, fill in the port information in the address list, supporting two formats: 22 and 100/200.
 When group_type is `domain`, fill in the domain name information in the address list.
* `group_name` - (Required) The name of the address book.
* `group_type` - (Required, ForceNew) The type of the address book. Valid values: `ip`, `port`, `domain`.
* `description` - (Optional) The description of the address book.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `ref_cnt` - The reference count of the address book.


## Import
AddressBook can be imported using the id, e.g.
```
$ terraform import volcengine_address_book.default resource_id
```

