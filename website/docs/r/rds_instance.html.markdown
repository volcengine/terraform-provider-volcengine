---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_instance"
sidebar_current: "docs-volcengine-resource-rds_instance"
description: |-
  Provides a resource to manage rds instance
---
# volcengine_rds_instance
Provides a resource to manage rds instance
## Example Usage
```hcl
resource "volcengine_rds_instance" "foo" {
  region             = "cn-north-4"
  zone               = "cn-langfang-b"
  instance_name      = "tf-test"
  db_engine          = "MySQL"
  db_engine_version  = "MySQL_Community_5_7"
  vpc_id             = "vpc-3cj17x7u9bzeo6c6rrtzfpaeb"
  instance_type      = "HA"
  charge_type        = "PostPaid"
  storage_type       = "LocalSSD"
  storage_space_gb   = 100
  instance_spec_name = "rds.mysql.1c2g"
  subnet_id          = "subnet-1g0d4fkh1nabk8ibuxx1jtvss"
}
```
## Argument Reference
The following arguments are supported:
* `charge_type` - (Required, ForceNew) Billing type. Value:
PostPaid: Postpaid (pay-as-you-go).
Prepaid: Prepaid (yearly and monthly).
* `db_engine_version` - (Required, ForceNew) Instance type. Value:
MySQL_Community_5_7
MySQL_8_0.
* `instance_spec_name` - (Required, ForceNew) Instance specification name, you can specify the specification name of the instance to be created. Value:
rds.mysql.1c2g
rds.mysql.2c4g
rds.mysql.4c8g
rds.mysql.4c16g
rds.mysql.8c32g
rds.mysql.16c64g
rds.mysql.16c128g
rds.mysql.32c128g
rds.mysql.32c256g.
* `instance_type` - (Required, ForceNew) Instance type. Value:
HA: High availability version.
* `storage_space_gb` - (Required, ForceNew) The storage space(GB) of the RDS instance.
* `storage_type` - (Required, ForceNew) Instance storage type. Value:
LocalSSD: Local SSD disk.
* `subnet_id` - (Required, ForceNew) Subnet ID. The subnet must belong to the selected Availability Zone.
* `vpc_id` - (Required, ForceNew) The vpc ID of the RDS instance.
* `zone` - (Required, ForceNew) The available zone of the RDS instance.
* `auto_renew` - (Optional, ForceNew) Whether to automatically renew. Default: false. Value:
true: yes.
false: no.
* `db_engine` - (Optional, ForceNew) Database type. Value:
MySQL (default).
* `instance_name` - (Optional, ForceNew) Set the name of the instance. The naming rules are as follows:

Cannot start with a number, a dash (-).
It can only contain Chinese characters, letters, numbers, underscores (_) and underscores (-).
The length needs to be within 1~128 characters.
* `prepaid_period` - (Optional, ForceNew) The purchase cycle in the prepaid scenario. Value:
Month: monthly subscription.
Year: yearly subscription.
* `project_name` - (Optional, ForceNew) Select the project to which the instance belongs. If this parameter is left blank, the new instance will not be added to any project.
* `region` - (Optional, ForceNew) The region of the RDS instance.
* `super_account_name` - (Optional, ForceNew) Fill in the high-privileged user account name. The naming rules are as follows:
Unique name.
Start with a letter and end with a letter or number.
Consists of lowercase letters, numbers, or underscores (_).
The length is 2~32 characters.
[Keywords](https://www.volcengine.com/docs/6313/66162) are not allowed for account names.
* `supper_account_password` - (Optional, ForceNew) Set a high-privilege account password. The rules are as follows:
Only uppercase and lowercase letters, numbers and the following special characters _#!@$%^*()+=-.
The length needs to be within 8~32 characters.
Contains at least 3 of uppercase letters, lowercase letters, numbers or special characters.
* `used_time` - (Optional, ForceNew) The purchase time of RDS instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `connection_info` - The connection info ot the RDS instance.
    * `enable_read_only` - Whether global read-only is enabled.
    * `enable_read_write_splitting` - Whether read-write separation is enabled.
    * `internal_domain` - The internal domain of the RDS instance.
    * `internal_port` - The interval port of the RDS instance.
    * `public_domain` - The public domain of the RDS instance.
    * `public_port` - The public port of the RDS instance.


## Import
RDS Instance can be imported using the id, e.g.
```
$ terraform import volcengine_rds_instance.default mysql-42b38c769c4b
```

