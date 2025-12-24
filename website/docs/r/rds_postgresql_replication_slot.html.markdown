---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_replication_slot"
sidebar_current: "docs-volcengine-resource-rds_postgresql_replication_slot"
description: |-
  Provides a resource to manage rds postgresql replication slot
---
# volcengine_rds_postgresql_replication_slot
Provides a resource to manage rds postgresql replication slot
## Example Usage
```hcl
resource "volcengine_rds_postgresql_replication_slot" "example" {
  instance_id = "postgres-72715e0d9f58"
  slot_name   = "my_standby_slot1"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of the PostgreSQL instance.
* `slot_name` - (Required, ForceNew) The name of the slot.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RdsPostgresqlReplicationSlot can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_replication_slot.default resource_id
```

