---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_workspaces"
sidebar_current: "docs-volcengine-datasource-vmp_workspaces"
description: |-
  Use this data source to query detailed information of vmp workspaces
---
# volcengine_vmp_workspaces
Use this data source to query detailed information of vmp workspaces
## Example Usage
```hcl
resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-1"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-1"
  username                  = "admin123"
  password                  = "*******"
}

data "volcengine_vmp_workspaces" "foo" {
  ids = [volcengine_vmp_workspace.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Workspace IDs.
* `instance_type_ids` - (Optional) A list of Instance Type IDs.
* `name` - (Optional) The name of workspace.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of vmp workspace.
* `statuses` - (Optional) A list of Workspace status.
* `tags` - (Optional) The tags of vmp workspace.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `values` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `total_count` - The total count of query.
* `workspaces` - The collection of query.
    * `create_time` - The create time of workspace.
    * `delete_protection_enabled` - Whether enable delete protection.
    * `description` - The description of workspace.
    * `id` - The ID of workspace.
    * `instance_type_id` - The id of instance type.
    * `name` - The name of workspace.
    * `overdue_reclaim_time` - The overdue reclaim time.
    * `project_name` - The project name of vmp workspace.
    * `prometheus_query_intranet_endpoint` - The prometheus query intranet endpoint.
    * `prometheus_write_intranet_endpoint` - The prometheus write intranet endpoint.
    * `status` - The status of workspace.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `username` - The username of workspace.


