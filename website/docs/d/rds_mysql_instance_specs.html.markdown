---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_instance_specs"
sidebar_current: "docs-volcengine-datasource-rds_mysql_instance_specs"
description: |-
  Use this data source to query detailed information of rds mysql instance specs
---
# volcengine_rds_mysql_instance_specs
Use this data source to query detailed information of rds mysql instance specs
## Example Usage
```hcl
data "volcengine_rds_mysql_instance_specs" "foo" {
  db_engine_version = ""
  instance_type     = ""
  spec_code         = ""
  zone_id           = ""
}
```
## Argument Reference
The following arguments are supported:
* `db_engine_version` - (Optional) Compatible version. Values:
MySQL_5_7: MySQL 5.7 version. Default value.
MySQL_8_0: MySQL 8.0 version.
* `instance_type` - (Optional) Instance type. The value is DoubleNode.
* `output_file` - (Optional) File name where to save data source results.
* `spec_code` - (Optional) Instance specification code.
* `zone_id` - (Optional) Availability zone ID.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instance_specs` - The collection of query.
    * `connection` - Default value of maximum number of connections.
    * `db_engine_version` - Compatible version. Values:
MySQL_5_7: MySQL 5.7 version. Default value.
MySQL_8_0: MySQL 8.0 version.
    * `instance_type` - Instance type. The value is DoubleNode.
    * `iops` - Maximum IOPS per second.
    * `memory` - Memory size, in GB.
    * `qps` - Queries Per Second (QPS).
    * `region_id` - The id of the region.
    * `spec_code` - Instance specification code.
    * `spec_family` - Instance specification type. Values:
General: Exclusive specification (formerly "General Purpose").
Shared: General specification (formerly "Shared Type").
    * `spec_status` - The status of the available zone where the specification is located includes the following statuses:
Normal: On sale.
Soldout: Sold out.
    * `storage_max` - Maximum storage space, in GB.
    * `storage_min` - Minimum storage space, in GB.
    * `storage_step` - Disk step size, in GB.
    * `vcpu` - Number of vCPUs.
    * `zone_id` - Availability zone ID.
* `total_count` - The total count of query.


