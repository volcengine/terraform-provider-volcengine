---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance_specs"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_instance_specs"
description: |-
  Use this data source to query detailed information of rds postgresql instance specs
---
# volcengine_rds_postgresql_instance_specs
Use this data source to query detailed information of rds postgresql instance specs
## Example Usage
```hcl
data "volcengine_rds_postgresql_instance_specs" "example" {
  zone_id           = "cn-chongqing-a"
  db_engine_version = "PostgreSQL_12"
  spec_code         = "rds.postgres.32c128g"
  storage_type      = "LocalSSD"
}
```
## Argument Reference
The following arguments are supported:
* `db_engine_version` - (Optional) The version of the RDS PostgreSQL instance.
* `output_file` - (Optional) File name where to save data source results.
* `spec_code` - (Optional) Instance specification code.
* `storage_type` - (Optional) Storage type, fixed to LocalSSD.
* `zone_id` - (Optional) Primary availability zone ID.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instance_specs` - Available instance specs.
    * `connection` - The maximum number of connections supported by the instance.
    * `db_engine_version` - The version of the RDS PostgreSQL instance.
    * `memory` - The memory size of the instance. Unit: GB.
    * `region_id` - The ID of the region.
    * `spec_code` - Instance specification code.
    * `storage_type` - Storage type, fixed to LocalSSD.
    * `v_cpu` - The number of vCPUs of the instance.
    * `zone_id` - Supported availability zone ID.
* `total_count` - The total count of query.


