---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance_state"
sidebar_current: "docs-volcengine-resource-rds_postgresql_instance_state"
description: |-
  Provides a resource to manage rds postgresql instance state
---
# volcengine_rds_postgresql_instance_state
Provides a resource to manage rds postgresql instance state
## Example Usage
```hcl
resource "volcengine_rds_postgresql_instance_state" "example" {
  instance_id = "postgres-72715e0d9f58"
  action      = "Restart"
}
```
## Argument Reference
The following arguments are supported:
* `action` - (Required, ForceNew) The action to perform on the instance. Valid value: Restart.
* `instance_id` - (Required, ForceNew) The ID of the RDS PostgreSQL instance to perform the action on.
* `apply_scope` - (Optional) The scope of the action. Valid values: AllNode, CustomNode. Default value: AllNode.
* `custom_node_ids` - (Optional) The ID of the read-only node(s) to restart. Required if apply_scope is CustomNode.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RdsPostgresqlInstanceState can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_instance_state.default resource_id
```

