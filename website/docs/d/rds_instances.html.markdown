---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_instances"
sidebar_current: "docs-volcengine-datasource-rds_instances"
description: |-
  Use this data source to query detailed information of rds instances
---
# volcengine_rds_instances
Use this data source to query detailed information of rds instances
## Example Usage
```hcl
data "volcengine_rds_instances" "default" {
  instance_id = "mysql-0fdd3bab2e7c"
}
```
## Argument Reference
The following arguments are supported:
* `create_end_time` - (Optional) The end time of creating RDS instance.
* `create_start_time` - (Optional) The start time of creating RDS instance.
* `instance_id` - (Optional) The id of the RDS instance.
* `instance_status` - (Optional) The status of the RDS instance.
* `instance_type` - (Optional) The type of the RDS instance.
* `name_regex` - (Optional) A Name Regex of RDS instance.
* `output_file` - (Optional) File name where to save data source results.
* `region` - (Optional) The region of the RDS instance.
* `zone` - (Optional) The available zone of the RDS instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rds_instances` - The collection of RDS instance query.
    * `charge_status` - The charge status of the RDS instance.
    * `charge_type` - The charge type of the RDS instance.
    * `connection_info` - The connection info ot the RDS instance.
        * `enable_read_only` - Whether global read-only is enabled.
        * `enable_read_write_splitting` - Whether read-write separation is enabled.
        * `internal_domain` - The internal domain of the RDS instance.
        * `internal_port` - The interval port of the RDS instance.
        * `public_domain` - The public domain of the RDS instance.
        * `public_port` - The public port of the RDS instance.
    * `create_time` - The create time of the RDS instance.
    * `db_engine_version` - The engine version of the RDS instance.
    * `db_engine` - The engine of the RDS instance.
    * `id` - The ID of the RDS instance.
    * `instance_id` - The ID of the RDS instance.
    * `instance_name` - The name of the RDS instance.
    * `instance_spec` - The spec type detail of RDS instance.
        * `cpu_num` - The cpu core count of spec type.
        * `mem_in_gb` - The memory size(GB) of spec type.
        * `spec_name` - The name of spec type.
    * `instance_status` - The status of the RDS instance.
    * `instance_type` - The type of the RDS instance.
    * `region` - The region of the RDS instance.
    * `storage_space_gb` - The total storage GB of the RDS instance.
    * `update_time` - The update time of the RDS instance.
    * `vpc_id` - The vpc ID of the RDS instance.
    * `zone` - The available zone of the RDS instance.
* `total_count` - The total count of RDS instance query.


