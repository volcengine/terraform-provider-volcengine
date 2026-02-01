---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_acls"
sidebar_current: "docs-volcengine-datasource-acls"
description: |-
  Use this data source to query detailed information of acls
---
# volcengine_acls
Use this data source to query detailed information of acls
## Example Usage
```hcl
data "volcengine_acls" "default" {
  ids = ["acl-3ti8n0rurx4bwbh9jzdy"]
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `acl_name` - (Optional) The name of acl.
* `ids` - (Optional) A list of Acl IDs.
* `name_regex` - (Optional) A Name Regex of Acl.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The ProjectName of Acl.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `acls` - The collection of Acl query.
    * `acl_entries` - The acl entry list of the Acl.
        * `description` - The description of the AclEntry.
        * `entry` - The address range of the IP entry.
    * `acl_entry_count` - The count of acl entry.
    * `acl_id` - The ID of Acl.
    * `acl_name` - The Name of Acl.
    * `create_time` - Creation time of Acl.
    * `description` - The description of Acl.
    * `id` - The ID of Acl.
    * `listener_details` - The listener details of the Acl.
        * `acl_type` - The control method of the listener for this Acl. Valid values: `black`, `white`.
        * `listener_id` - The ID of the listener.
        * `listener_name` - The name of the listener.
        * `port` - The port receiving request of the listener.
        * `protocol` - The protocol of the listener.
    * `listeners` - The listeners of Acl.
    * `project_name` - The ProjectName of Acl.
    * `service_managed` - Whether the Acl is managed by service.
    * `status` - The status of the Acl.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `update_time` - Update time of Acl.
* `total_count` - The total count of Acl query.


