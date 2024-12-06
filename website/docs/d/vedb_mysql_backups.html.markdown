---
subcategory: "VEDB_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_vedb_mysql_backups"
sidebar_current: "docs-volcengine-datasource-vedb_mysql_backups"
description: |-
  Use this data source to query detailed information of vedb mysql backups
---
# volcengine_vedb_mysql_backups
Use this data source to query detailed information of vedb mysql backups
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

resource "volcengine_vedb_mysql_backup" "foo" {
  instance_id = volcengine_vedb_mysql_instance.foo.id
  backup_policy {
    backup_time             = "18:00Z-20:00Z"
    full_backup_period      = "Monday,Tuesday,Wednesday"
    backup_retention_period = 8
  }
}

data "volcengine_vedb_mysql_backups" "foo" {
  instance_id = volcengine_vedb_mysql_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the instance.
* `backup_end_time` - (Optional) The end time of the backup.
* `backup_method` - (Optional) Backup method. Currently, only physical backup is supported. The value is Physical.
* `backup_start_time` - (Optional) The start time of the backup.
* `backup_status` - (Optional) The status of the backup.
* `backup_type` - (Optional) The type of the backup.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `backups` - The collection of query.
    * `backup_end_time` - The end time of the backup.
    * `backup_file_size` - The size of the backup file.
    * `backup_id` - The id of the backup.
    * `backup_method` - The name of the backup method.
    * `backup_policy` - Data backup strategy for instances.
        * `backup_retention_period` - Data backup retention period, value: 7 to 30 days.
        * `backup_time` - The time for executing the backup task. The interval window is two hours. Format: HH:mmZ-HH:mmZ (UTC time).
        * `continue_backup` - Whether to enable continuous backup. The value is fixed as true.
        * `full_backup_period` - Full backup period. Multiple values are separated by English commas (,). Values:
Monday: Monday.
Tuesday: Tuesday.
Wednesday: Wednesday.
Thursday: Thursday.
Friday: Friday.
Saturday: Saturday.
Sunday: Sunday.
        * `instance_id` - The id of the instance.
    * `backup_start_time` - The start time of the backup.
    * `backup_status` - The status of the backup.
    * `backup_type` - The type of the backup.
    * `consistent_time` - The time point of consistent backup, in the format: yyyy-MM-ddTHH:mm:ssZ (UTC time).
    * `create_type` - The type of the backup create.
    * `id` - The id of the backup.
* `total_count` - The total count of query.


