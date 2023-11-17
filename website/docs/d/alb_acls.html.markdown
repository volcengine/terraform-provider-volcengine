---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_acls"
sidebar_current: "docs-volcengine-datasource-alb_acls"
description: |-
  Use this data source to query detailed information of alb acls
---
# volcengine_alb_acls
Use this data source to query detailed information of alb acls
## Example Usage
```hcl
data "volcengine_alb_acls" "default" {
  project_name = "default"
  ids          = ["acl-1g72w6z11ighs2zbhq4v3rvh4"]
}
```
## Argument Reference
The following arguments are supported:
* `acl_name` - (Optional) The name of acl.
* `ids` - (Optional) A list of Acl IDs.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The name of project.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `acls` - The collection of Acl query.
    * `acl_entries` - The entries info of acl.
        * `description` - The description of entry.
        * `entry` - The info of entry.
    * `acl_entry_count` - The count of acl entry.
    * `acl_id` - The ID of Acl.
    * `acl_name` - The Name of Acl.
    * `create_time` - Creation time of Acl.
    * `description` - The description of Acl.
    * `id` - The ID of Acl.
    * `listeners` - The listeners of acl.
        * `acl_type` - The type of acl.
        * `listener_id` - The ID of Listener.
        * `listener_name` - The Name of Listener.
        * `port` - The port info of listener.
        * `protocol` - The protocol info of listener.
    * `project_name` - The project name of Acl.
    * `update_time` - Update time of Acl.
* `total_count` - The total count of Acl query.


