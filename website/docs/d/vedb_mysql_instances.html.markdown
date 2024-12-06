---
subcategory: "VEDB_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_vedb_mysql_instances"
sidebar_current: "docs-volcengine-datasource-vedb_mysql_instances"
description: |-
  Use this data source to query detailed information of vedb mysql instances
---
# volcengine_vedb_mysql_instances
Use this data source to query detailed information of vedb mysql instances
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
  zone_id     = data.volcengine_zones.foo.zones[2].id
  vpc_id      = volcengine_vpc.foo.id
}


resource "volcengine_vedb_mysql_instance" "foo" {
  charge_type         = "PostPaid"
  storage_charge_type = "PostPaid"
  db_engine_version   = "MySQL_8_0"
  db_minor_version    = "3.0"
  node_number         = 2
  node_spec           = "vedb.mysql.x4.large"
  subnet_id           = volcengine_subnet.foo.id
  instance_name       = "tf-test"
  project_name        = "testA"
  tags {
    key   = "tftest"
    value = "tftest"
  }
  tags {
    key   = "tftest2"
    value = "tftest2"
  }
}

data "volcengine_vedb_mysql_instances" "foo" {
  instance_id = volcengine_vedb_mysql_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `charge_type` - (Optional) The charge type of the veDB Mysql instance.
* `create_time_end` - (Optional) The end time of creating veDB Mysql instance.
* `create_time_start` - (Optional) The start time of creating veDB Mysql instance.
* `db_engine_version` - (Optional) The version of the veDB Mysql instance.
* `instance_id` - (Optional) The id of the veDB Mysql instance.
* `instance_name` - (Optional) The name of the veDB Mysql instance.
* `instance_status` - (Optional) The status of the veDB Mysql instance.
* `name_regex` - (Optional) A Name Regex of veDB mysql instance.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of the veDB Mysql instance.
* `tags` - (Optional) Tags.
* `zone_id` - (Optional) The available zone of the veDB Mysql instance.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instances` - The collection of query.
    * `auto_renew` - Whether auto-renewal is enabled in the prepaid scenario. Values:
true: Auto-renewal is enabled.
false: Auto-renewal is not enabled.
    * `charge_end_time` - The billing expiration time in the prepaid scenario, in the format: yyyy-MM-ddTHH:mm:ssZ (UTC time).
    * `charge_start_time` - The time when billing starts. Format: yyyy-MM-ddTHH:mm:ssZ (UTC time).
    * `charge_status` - Payment status:
Normal: Normal.
Overdue: In arrears.
Shutdown: Shut down.
    * `charge_type` - Calculate the billing type. Values:
PostPaid: Pay-as-you-go (postpaid).
PrePaid: Monthly/yearly subscription (prepaid).
    * `create_time` - The create time of the veDB Mysql instance.
    * `db_engine_version` - The engine version of the veDB Mysql instance.
    * `id` - The ID of the veDB Mysql instance.
    * `instance_id` - The ID of the veDB Mysql instance.
    * `instance_name` - The name of the veDB Mysql instance.
    * `instance_status` - The status of the veDB Mysql instance.
    * `lower_case_table_names` - Whether the table name is case sensitive, the default value is 1.
Ranges:
0: Table names are stored as fixed and table names are case-sensitive.
1: Table names will be stored in lowercase and table names are not case sensitive.
    * `nodes` - Detailed information of instance nodes.
        * `memory` - Memory size, in GiB.
        * `node_id` - The id of the node.
        * `node_spec` - Node specification of an instance.
        * `node_type` - Node type. Values:
Primary: Primary node.
ReadOnly: Read-only node.
        * `v_cpu` - CPU size. For example, when the value is 1, it means the CPU size is 1U.
        * `zone_id` - The zone id.
    * `overdue_reclaim_time` - Expected release time when shut down due to arrears. Format: yyyy-MM-ddTHH:mm:ssZ (UTC time).
    * `overdue_time` - Overdue shutdown time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC time).
    * `pre_paid_storage_in_gb` - Total storage capacity in GiB for prepaid services.
    * `project_name` - The project name of the veDB Mysql instance.
    * `region_id` - The region id.
    * `storage_charge_type` - Storage billing type. Values:
PostPaid: Pay-as-you-go (postpaid).
PrePaid: Monthly/yearly subscription (prepaid).
    * `storage_used_gib` - Used storage size, unit: GiB.
    * `subnet_id` - The subnet ID of the veDB Mysql instance.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `time_zone` - Time zone.
    * `vpc_id` - The vpc ID of the veDB Mysql instance.
    * `zone_ids` - The available zone of the veDB Mysql instance.
* `total_count` - The total count of query.


