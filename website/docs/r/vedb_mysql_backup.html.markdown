---
subcategory: "VEDB_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_vedb_mysql_backup"
sidebar_current: "docs-volcengine-resource-vedb_mysql_backup"
description: |-
  Provides a resource to manage vedb mysql backup
---
# volcengine_vedb_mysql_backup
Provides a resource to manage vedb mysql backup
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
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The id of the instance.
* `backup_method` - (Optional, ForceNew) Backup method. Currently, only physical backup is supported. The value is Physical.
* `backup_policy` - (Optional) Data backup strategy for instances.
* `backup_type` - (Optional, ForceNew) Backup type. Currently, only full backup is supported. The value is Full.

The `backup_policy` object supports the following:

* `backup_retention_period` - (Required) Data backup retention period, value: 7 to 30 days.
* `backup_time` - (Required) The time for executing the backup task has an interval window of 2 hours and must be an even-hour time. Format: HH:mmZ-HH:mmZ (UTC time).
* `full_backup_period` - (Required) Full backup period. It is recommended to select at least 2 days per week for full backup. Multiple values are separated by English commas (,). Values: Monday: Monday. Tuesday: Tuesday. Wednesday: Wednesday. Thursday: Thursday. Friday: Friday. Saturday: Saturday. Sunday: Sunday.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `backup_id` - The id of the backup.


## Import
VedbMysqlBackup can be imported using the instance id and backup id, e.g.
```
$ terraform import volcengine_vedb_mysql_backup.default instanceID:backupId
```

