---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpc_prefix_list"
sidebar_current: "docs-volcengine-resource-vpc_prefix_list"
description: |-
  Provides a resource to manage vpc prefix list
---
# volcengine_vpc_prefix_list
Provides a resource to manage vpc prefix list
## Example Usage
```hcl
resource "volcengine_vpc_prefix_list" "foo" {
  prefix_list_name = "acc-test-prefix"
  max_entries      = 7
  description      = "acc test description"
  ip_version       = "IPv4"
  prefix_list_entries {
    cidr        = "192.168.4.0/28"
    description = "acc-test-1"
  }
  prefix_list_entries {
    cidr        = "192.168.9.0/28"
    description = "acc-test-4"
  }
  prefix_list_entries {
    cidr        = "192.168.8.0/28"
    description = "acc-test-5"
  }
  tags {
    key   = "tf-key1"
    value = "tf-value1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `max_entries` - (Required) Maximum number of entries, which is the maximum number of entries that can be added to the prefix list. The value range is 1 to 200.
* `description` - (Optional) The description of the prefix list.
* `ip_version` - (Optional, ForceNew) IP version type. Possible values:
IPv4 (default): IPv4 type.
IPv6: IPv6 type.
* `prefix_list_entries` - (Optional) Prefix list entry list.
* `prefix_list_name` - (Optional) The name of the prefix list.
* `tags` - (Optional, ForceNew) Tags.

The `prefix_list_entries` object supports the following:

* `cidr` - (Optional) CIDR of prefix list entries.
* `description` - (Optional) Description of prefix list entries.

The `tags` object supports the following:

* `key` - (Required, ForceNew) The Key of Tags.
* `value` - (Required, ForceNew) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `prefix_list_associations` - Collection of resources associated with VPC prefix list.
    * `resource_id` - Associated resource ID.
    * `resource_type` - Related resource types.


## Import
VpcPrefixList can be imported using the id, e.g.
```
$ terraform import volcengine_vpc_prefix_list.default resource_id
```

