---
subcategory: "ROCKETMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rocketmq_instances"
sidebar_current: "docs-volcengine-datasource-rocketmq_instances"
description: |-
  Use this data source to query detailed information of rocketmq instances
---
# volcengine_rocketmq_instances
Use this data source to query detailed information of rocketmq instances
## Example Usage
```hcl
data "volcengine_rocketmq_instances" "foo" {
  instance_id = "rocketmq-cnoeea6b32118fc2"
}
```
## Argument Reference
The following arguments are supported:
* `charge_type` - (Optional) The charge type of rocketmq instance. Valid values: `PostPaid`, `PrePaid`.
* `instance_id` - (Optional) The id of rocketmq instance.
* `instance_name` - (Optional) The name of rocketmq instance. This field support fuzzy query.
* `instance_status` - (Optional) The status of rocketmq instance.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of rocketmq instance.
* `spec` - (Optional) The spec of rocketmq instance.
* `tags` - (Optional) Tags.
* `version` - (Optional) The version of rocketmq instance. Valid values: `4.8`.
* `vpc_id` - (Optional) The vpc id of rocketmq instance.
* `zone_id` - (Optional) The zone id of rocketmq instance.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rocketmq_instances` - The collection of query.
    * `account_id` - The account id of the rocketmq instance.
    * `apply_private_dns_to_public` - Whether the private dns to public function is enabled for the rocketmq instance.
    * `available_queue_number` - The available queue number of the rocketmq instance.
    * `charge_detail` - The charge detail information of the rocketmq instance.
        * `auto_renew` - Whether to enable automatic renewal.
        * `charge_expire_time` - The charge expire time of the rocketmq instance.
        * `charge_start_time` - The charge start time of the rocketmq instance.
        * `charge_status` - The charge status of the rocketmq instance.
        * `charge_type` - The charge type of the rocketmq instance.
        * `overdue_reclaim_time` - The overdue reclaim time of the rocketmq instance.
        * `overdue_time` - The overdue time of the rocketmq instance.
        * `period_unit` - The period unit of the rocketmq instance.
    * `compute_spec` - The compute spec of the rocketmq instance.
    * `connection_info` - The connection information of the rocketmq.
        * `endpoint_address_ip` - The endpoint address ip of the rocketmq.
        * `endpoint_type` - The endpoint type of the rocketmq.
        * `internal_endpoint` - The internal endpoint of the rocketmq.
        * `network_type` - The network type of the rocketmq.
        * `public_endpoint` - The public endpoint of the rocketmq.
    * `create_time` - The create time of the rocketmq instance.
    * `eip_id` - The eip id of the rocketmq instance.
    * `enable_ssl` - Whether the ssl authentication is enabled for the rocketmq instance.
    * `file_reserved_time` - The reserved time of messages on the RocketMQ server of the message queue. Messages that exceed the reserved time will be cleared after expiration. The unit is in hours.
    * `id` - The id of the rocketmq instance.
    * `instance_description` - The description of the rocketmq instance.
    * `instance_id` - The id of the rocketmq instance.
    * `instance_name` - The name of the rocketmq instance.
    * `instance_status` - The status of the rocketmq instance.
    * `project_name` - The project name of the rocketmq instance.
    * `region_id` - The region id of the rocketmq instance.
    * `ssl_mode` - The ssl mode of the rocketmq instance.
    * `storage_space` - The total storage space of the rocketmq instance.
    * `subnet_id` - The subnet id of the rocketmq instance.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `used_group_number` - The used group number of the rocketmq instance.
    * `used_queue_number` - The used queue number of the rocketmq instance.
    * `used_storage_space` - The used storage space of the rocketmq instance.
    * `used_topic_number` - The used topic number of the rocketmq instance.
    * `version` - The version of the rocketmq instance.
    * `vpc_id` - The vpc id of the rocketmq instance.
    * `zone_id` - The zone id of the rocketmq instance.
* `total_count` - The total count of query.


