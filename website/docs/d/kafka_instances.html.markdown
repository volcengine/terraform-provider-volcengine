---
subcategory: "KAFKA"
layout: "volcengine"
page_title: "Volcengine: volcengine_kafka_instances"
sidebar_current: "docs-volcengine-datasource-kafka_instances"
description: |-
  Use this data source to query detailed information of kafka instances
---
# volcengine_kafka_instances
Use this data source to query detailed information of kafka instances
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_kafka_instance" "foo" {
  instance_name        = "acc-test-kafka"
  instance_description = "tf-test"
  version              = "2.2.2"
  compute_spec         = "kafka.20xrate.hw"
  subnet_id            = volcengine_subnet.foo.id
  user_name            = "tf-user"
  user_password        = "tf-pass!@q1"
  charge_type          = "PostPaid"
  storage_space        = 300
  partition_number     = 350
  project_name         = "default"
  tags {
    key   = "k1"
    value = "v1"
  }

  parameters {
    parameter_name  = "MessageMaxByte"
    parameter_value = "12"
  }
  parameters {
    parameter_name  = "LogRetentionHours"
    parameter_value = "70"
  }
}

data "volcengine_kafka_instances" "default" {
  instance_id = volcengine_kafka_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Optional) The id of instance.
* `instance_name` - (Optional) The name of instance.
* `instance_status` - (Optional) The status of instance.
* `output_file` - (Optional) File name where to save data source results.
* `tags` - (Optional) The tags of instance.
* `zone_id` - (Optional) The zone id of instance.

The `tags` object supports the following:

* `key` - (Required) The key of tag.
* `value` - (Required) The value of tag.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instances` - The collection of query.
    * `account_id` - The id of account.
    * `auto_renew` - The auto renew status of instance.
    * `charge_expire_time` - The charge expire time of instance.
    * `charge_start_time` - The charge start time of instance.
    * `charge_status` - The charge status of instance.
    * `charge_type` - The charge type of instance.
    * `compute_spec` - The compute spec of instance.
    * `connection_info` - Connection info of the instance.
        * `endpoint_type` - The endpoint type of instance.
        * `internal_endpoint` - The internal endpoint of instance.
        * `network_type` - The network type of instance.
        * `public_endpoint` - The public endpoint of instance.
    * `create_time` - The create time of instance.
    * `eip_id` - The id of eip.
    * `id` - The id of instance.
    * `instance_description` - The description of instance.
    * `instance_id` - The id of instance.
    * `instance_name` - The name of instance.
    * `instance_status` - The status of instance.
    * `overdue_reclaim_time` - The overdue reclaim time of instance.
    * `overdue_time` - The overdue time of instance.
    * `parameters` - Parameters of the instance.
        * `parameter_name` - Parameter name.
        * `parameter_value` - Parameter value.
    * `period_unit` - The period unit of instance.
    * `private_domain_on_public` - Whether enable private domain on public.
    * `project_name` - The name of project.
    * `region_id` - The id of region.
    * `storage_space` - The storage space of instance.
    * `storage_type` - The storage type of instance.
    * `subnet_id` - The id of subnet.
    * `tags` - The Tags of instance.
        * `key` - The key of tags.
        * `value` - The value of tags.
    * `usable_partition_number` - The usable partition number of instance.
    * `used_group_number` - The used group number of instance.
    * `used_partition_number` - The used partition number of instance.
    * `used_storage_space` - The used storage space of instance.
    * `used_topic_number` - The used topic number of instance.
    * `version` - The version of instance.
    * `vpc_id` - The id of vpc.
    * `zone_id` - The id of zone.
* `total_count` - The total count of query.


