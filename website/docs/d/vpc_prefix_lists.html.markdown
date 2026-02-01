---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpc_prefix_lists"
sidebar_current: "docs-volcengine-datasource-vpc_prefix_lists"
description: |-
  Use this data source to query detailed information of vpc prefix lists
---
# volcengine_vpc_prefix_lists
Use this data source to query detailed information of vpc prefix lists
## Example Usage
```hcl
resource "volcengine_vpc_prefix_list" "foo" {
  prefix_list_name = "acc-test-prefix"
  max_entries      = 3
  description      = "acc test description"
  ip_version       = "IPv4"
  prefix_list_entries {
    cidr        = "192.168.4.0/28"
    description = "acc-test-1"
  }
  prefix_list_entries {
    cidr        = "192.168.5.0/28"
    description = "acc-test-2"
  }
  tags {
    key   = "tf-key1"
    value = "tf-value1"
  }
}

data "volcengine_vpc_prefix_lists" "foo" {
  ids = [volcengine_vpc_prefix_list.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of prefix list ids.
* `ip_version` - (Optional) IP version of prefix list.
* `output_file` - (Optional) File name where to save data source results.
* `prefix_list_name` - (Optional) A Name of prefix list.
* `tag_filters` - (Optional) List of tag filters.

The `tag_filters` object supports the following:

* `key` - (Optional) The key of the tag.
* `values` - (Optional) The values of the tag.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `prefix_lists` - The collection of query.
    * `association_count` - Number of associated resources for prefix list.
    * `cidrs` - CIDR address block information for prefix list.
    * `creation_time` - The creation time of the prefix list.
    * `description` - Description.
    * `id` - The id of the prefix list.
    * `ip_version` - The ip version of the prefix list.
    * `max_entries` - Maximum number of entries, which is the maximum number of items that can be added to the prefix list.
    * `prefix_list_associations` - Collection of resources associated with VPC prefix list.
        * `resource_id` - Associated resource ID.
        * `resource_type` - Related resource types.
    * `prefix_list_entries` - The prefix list entries.
        * `cidr` - CIDR address blocks for prefix list entries.
        * `description` - Description.
        * `prefix_list_id` - The prefix list id.
    * `prefix_list_id` - The prefix list id.
    * `prefix_list_name` - The prefix list name.
    * `status` - The status of the prefix list.
    * `update_time` - The update time of the prefix list.
* `total_count` - The total count of query.


