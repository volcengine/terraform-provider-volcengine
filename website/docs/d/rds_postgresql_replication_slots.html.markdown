---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_replication_slots"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_replication_slots"
description: |-
  Use this data source to query detailed information of rds postgresql replication slots
---
# volcengine_rds_postgresql_replication_slots
Use this data source to query detailed information of rds postgresql replication slots
## Example Usage
```hcl
data "volcengine_rds_postgresql_replication_slots" "example" {
  instance_id = "postgres-72715e0d9f58"
  slot_name   = "my_standby_slot1"
  slot_status = "INACTIVE"
  slot_type   = "physical"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the PostgreSQL instance.
* `data_base` - (Optional) The database where the replication slot is located.
* `ip_address` - (Optional) The ip address.
* `output_file` - (Optional) File name where to save data source results.
* `plugin` - (Optional) The name of the plugin used by the logical replication slot to parse WAL logs.
* `slot_name` - (Optional) The name of the slot.
* `slot_status` - (Optional) The status of the replication slot: ACTIVE or INACTIVE.
* `slot_type` - (Optional) The type of the slot: physical or logical.
* `temporary` - (Optional) Whether the slot is temporary.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `replication_slots` - Replication slots under the specified query conditions in the instance.
    * `plugin` - The name of the plugin used by the logical replication slot to parse WAL logs.
    * `slot_name` - The name of the slot.
    * `slot_type` - The type of the slot: physical or logical.
* `total_count` - The total count of query.


