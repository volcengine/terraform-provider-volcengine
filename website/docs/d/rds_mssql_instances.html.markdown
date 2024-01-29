---
subcategory: "RDS_MSSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mssql_instances"
sidebar_current: "docs-volcengine-datasource-rds_mssql_instances"
description: |-
  Use this data source to query detailed information of rds mssql instances
---
# volcengine_rds_mssql_instances
Use this data source to query detailed information of rds mssql instances
## Example Usage
```hcl
data "volcengine_rds_mssql_instances" "foo" {
  instance_id = "mssql-d2fc5abe****"
}
```
## Argument Reference
The following arguments are supported:
* `charge_type` - (Optional) The charge type. Valid values: `PostPaid`, `PrePaid`.
* `create_time_end` - (Optional) The end time of creating the instance, using UTC time format.
* `create_time_start` - (Optional) The start time of creating the instance, using UTC time format.
* `db_engine_version` - (Optional) Compatible version. Valid values: `SQLServer_2019_Std`, `SQLServer_2019_Web`, `SQLServer_2019_Ent`.
* `instance_id` - (Optional) Id of the instance.
* `instance_name` - (Optional) Name of the instance.
* `instance_status` - (Optional) Status of the instance.
* `instance_type` - (Optional) Instance type. Valid values: `HA`, `Basic`, `Cluster`.
* `name_regex` - (Optional) A Name Regex of RDS mssql instance.
* `output_file` - (Optional) File name where to save data source results.
* `tags` - (Optional) Tags.
* `zone_id` - (Optional) The id of the zone.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instances` - The collection of query.
    * `backup_use` - The used backup space of the instance. Unit: GiB.
    * `charge_detail` - The charge detail.
        * `auto_renew` - Whether to enable automatic renewal in the prepaid scenario. This parameter can be set when ChargeType is Prepaid.
        * `charge_end_time` - Charge end time.
        * `charge_start_time` - Charge start time.
        * `charge_status` - The charge status.
        * `charge_type` - The charge type.
        * `overdue_reclaim_time` - Expected release time when overdue fees are shut down.
        * `overdue_time` - Time for Disconnection due to Unpaid Fees.
        * `period_unit` - Purchase cycle in prepaid scenarios. This parameter can be set when ChargeType is Prepaid.
        * `period` - Purchase duration in a prepaid scenario.
    * `connection_info` - The connection info of the instance.
        * `address` - The address info.
            * `dns_visibility` - Whether to enable private to public network resolution.
            * `domain` - The domain.
            * `eip_id` - The eip id for public address.
            * `ip_address` - The ip address.
            * `network_type` - The network type.
            * `port` - The port.
            * `subnet_id` - The subnet id for private address.
        * `description` - The description.
        * `endpoint_id` - The endpoint id.
        * `endpoint_name` - The endpoint name.
        * `endpoint_type` - The endpoint type.
    * `create_time` - The creation time of the instance.
    * `db_engine_version` - The db engine version.
    * `id` - The id of the instance.
    * `inner_version` - The inner version of the instance.
    * `instance_category` - The instance category.
    * `instance_id` - The id of the instance.
    * `instance_name` - The name of the instance.
    * `instance_status` - The status of the instance.
    * `instance_type` - The type of the instance.
    * `memory` - The Memory of the instance. Unit: GiB.
    * `node_detail_info` - Node detail information.
        * `create_time` - Node creation time.
        * `instance_id` - Instance ID.
        * `memory` - The Memory.
        * `node_id` - The Node ID.
        * `node_ip` - The node ip.
        * `node_spec` - The node spec.
        * `node_status` - The node status.
        * `node_type` - The node type.
        * `region_id` - The region id.
        * `update_time` - The update time.
        * `v_cpu` - CPU size. For example: 1 represents 1U.
        * `zone_id` - The zone id.
    * `node_spec` - The node spec.
    * `parameter_count` - The count of instance parameters.
    * `parameters` - The list of instance parameters.
        * `checking_code` - The valid value range of the parameter.
        * `force_modify` - Indicates whether the parameter running value can be modified.
        * `force_restart` - Indicates whether the instance needs to be restarted to take effect after modifying the running value of the parameter.
        * `parameter_default_value` - The default value of the parameter.
        * `parameter_description` - The description of the parameter.
        * `parameter_name` - The name of the parameter.
        * `parameter_type` - The type of the parameter.
        * `parameter_value` - The value of the parameter.
    * `port` - The port of the instance.
    * `primary_instance_id` - The id of the primary instance.
    * `project_name` - The project name.
    * `read_only_number` - The number of read only instance.
    * `region_id` - The region id.
    * `server_collation` - Server sorting rules.
    * `slow_query_enable` - Whether to enable slow query function.
    * `slow_query_time` - The slow query time. Unit: second.
    * `storage_space` - The storage space.
    * `storage_type` - The storage type.
    * `storage_use` - The used storage space.
    * `subnet_id` - The subnet id.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `time_zone` - The time zone.
    * `update_time` - The update time of the instance.
    * `v_cpu` - The CPU size of the instance. For example: 1 represents 1U.
    * `vpc_id` - The vpc id.
    * `zone_id` - The zone id.
* `total_count` - The total count of query.


