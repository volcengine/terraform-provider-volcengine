---
subcategory: "RABBITMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rabbitmq_instances"
sidebar_current: "docs-volcengine-datasource-rabbitmq_instances"
description: |-
  Use this data source to query detailed information of rabbitmq instances
---
# volcengine_rabbitmq_instances
Use this data source to query detailed information of rabbitmq instances
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

resource "volcengine_rabbitmq_instance" "foo" {
  zone_ids             = [data.volcengine_zones.foo.zones[0].id, data.volcengine_zones.foo.zones[1].id, data.volcengine_zones.foo.zones[2].id]
  subnet_id            = volcengine_subnet.foo.id
  version              = "3.8.18"
  user_name            = "acc-test-user"
  user_password        = "93f0cb0614Aab12"
  compute_spec         = "rabbitmq.n3.x2.small"
  storage_space        = 300
  instance_name        = "acc-test-rabbitmq"
  instance_description = "acc-test"
  charge_info {
    charge_type = "PostPaid"
  }
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

data "volcengine_rabbitmq_instances" "foo" {
  instance_id = volcengine_rabbitmq_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `charge_type` - (Optional) The charge type of rabbitmq instance.
* `instance_id` - (Optional) The id of rabbitmq instance. This field supports fuzzy query.
* `instance_name` - (Optional) The name of rabbitmq instance. This field supports fuzzy query.
* `instance_status` - (Optional) The status of rabbitmq instance.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of rabbitmq instance.
* `spec` - (Optional) The calculation specification of rabbitmq instance.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional) The vpc id of rabbitmq instance. This field supports fuzzy query.
* `zone_id` - (Optional) The zone id of rabbitmq instance. This field supports fuzzy query.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rabbitmq_instances` - The collection of query.
    * `account_id` - The account id of the rabbitmq instance.
    * `apply_private_dns_to_public` - Whether enable the public network parsing function of the rabbitmq instance.
    * `arch_type` - The type of the rabbitmq instance.
    * `charge_detail` - The charge detail information of the rabbitmq instance.
        * `auto_renew` - Whether to automatically renew in prepaid scenarios.
        * `charge_end_time` - The charge end time of the rabbitmq instance.
        * `charge_expire_time` - The charge expire time of the rabbitmq instance.
        * `charge_start_time` - The charge start time of the rabbitmq instance.
        * `charge_status` - The charge status of the rabbitmq instance.
        * `charge_type` - The charge type of the rabbitmq instance.
        * `overdue_reclaim_time` - The overdue reclaim time of the rabbitmq instance.
        * `overdue_time` - The overdue time of the rabbitmq instance.
    * `compute_spec` - The compute specification of the rabbitmq instance.
    * `create_time` - The create time of the rabbitmq instance.
    * `eip_id` - The eip id of the rabbitmq instance.
    * `endpoints` - The endpoint info of the rabbitmq instance.
        * `endpoint_type` - The endpoint type of the rabbitmq instance.
        * `internal_endpoint` - The internal endpoint of the rabbitmq instance.
        * `network_type` - The network type of the rabbitmq instance.
        * `public_endpoint` - The public endpoint of the rabbitmq instance.
    * `id` - The id of the rabbitmq instance.
    * `init_user_name` - The WebUI admin user name of the rabbitmq instance.
    * `instance_description` - The description of the rabbitmq instance.
    * `instance_id` - The id of the rabbitmq instance.
    * `instance_name` - The name of the rabbitmq instance.
    * `instance_status` - The status of the rabbitmq instance.
    * `project_name` - The project name of the rabbitmq instance.
    * `region_description` - The region description of the rabbitmq instance.
    * `region_id` - The region id of the rabbitmq instance.
    * `storage_space` - The total storage space of the rabbitmq instance. Unit: GiB.
    * `subnet_id` - The subnet id of the rabbitmq instance.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `used_storage_space` - The used storage space of the rabbitmq instance. Unit: GiB.
    * `version` - The version of the rabbitmq instance.
    * `vpc_id` - The vpc id of the rabbitmq instance.
    * `zone_description` - The zone description of the rabbitmq instance.
    * `zone_id` - The zone id of the rabbitmq instance.
* `total_count` - The total count of query.


