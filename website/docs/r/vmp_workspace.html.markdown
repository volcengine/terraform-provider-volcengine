---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_workspace"
sidebar_current: "docs-volcengine-resource-vmp_workspace"
description: |-
  Provides a resource to manage vmp workspace
---
# volcengine_vmp_workspace
Provides a resource to manage vmp workspace
## Example Usage
```hcl
resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-vmp-workspace"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test"
  username                  = "admin123"
  password                  = "Pass123456"
  project_name              = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `instance_type_id` - (Required, ForceNew) The instance type id of the workspace.
* `name` - (Required) The name of the workspace.
* `delete_protection_enabled` - (Optional) Whether enable delete protection.
* `description` - (Optional) The description of the workspace.
* `password` - (Optional) The password of the workspace.
* `project_name` - (Optional) The project name of the vmp workspace.
* `tags` - (Optional) Tags.
* `username` - (Optional) The username of the workspace.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of workspace.
* `overdue_reclaim_time` - The overdue reclaim time.
* `prometheus_push_intranet_endpoint` - The prometheus push intranet endpoint.
* `prometheus_query_intranet_endpoint` - The prometheus query intranet endpoint.
* `prometheus_write_intranet_endpoint` - The prometheus write intranet endpoint.
* `status` - The status of workspace.


## Import
Workspace can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_workspace.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
```

