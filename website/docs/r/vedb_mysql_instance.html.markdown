---
subcategory: "VEDB_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_vedb_mysql_instance"
sidebar_current: "docs-volcengine-resource-vedb_mysql_instance"
description: |-
  Provides a resource to manage vedb mysql instance
---
# volcengine_vedb_mysql_instance
Provides a resource to manage vedb mysql instance
## Notice
When Destroy this resource,If the resource charge type is PrePaid,Please unsubscribe the resource 
in  [Volcengine Console](https://console.volcengine.com/finance/unsubscribe/),when complete console operation,yon can
use 'terraform state rm ${resourceId}' to remove.
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
  project_name        = "default"
  tags {
    key   = "tftest"
    value = "tftest"
  }
}
```
## Argument Reference
The following arguments are supported:
* `charge_type` - (Required) Calculate the billing type. When calculating the billing type during instance creation, the possible values are as follows:
PostPaid: Pay-as-you-go (postpaid).
PrePaid: Monthly or yearly subscription (prepaid).
* `db_engine_version` - (Required, ForceNew) Database engine version, with a fixed value of MySQL_8_0.
* `node_number` - (Required) Number of instance nodes. The value range is from 2 to 16.
* `node_spec` - (Required) Node specification code of an instance.
* `subnet_id` - (Required, ForceNew) Subnet ID of the veDB Mysql instance.
* `auto_renew` - (Optional) Whether to automatically renew under the prepaid scenario. Values:
true: Automatically renew.
false: Do not automatically renew.
Description:
When the value of ChargeType (billing type) is PrePaid (monthly/yearly package), this parameter is required.
* `db_minor_version` - (Optional, ForceNew) veDB MySQL minor version. For detailed instructions on version numbers, please refer to Version Number Management.
 3.0 (default): veDB MySQL stable version, 100% compatible with MySQL 8.0.
 3.1: Natively supports HTAP application scenarios and accelerates complex queries.
 3.2: Natively supports HTAP application scenarios and accelerates complex queries. In addition, it has built-in cold data archiving capabilities. It can archive data with low-frequency access to object storage TOS to reduce storage costs.
* `db_time_zone` - (Optional, ForceNew) Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `instance_name` - (Optional) Instance name. Naming rules:
It cannot start with a number or a hyphen (-).
It can only contain Chinese characters, letters, numbers, underscores (_), and hyphens (-).
The length must be within 1 to 128 characters.
Description
If the instance name is not filled in, the instance ID will be used as the instance name.
When creating instances in batches, if an instance name is passed in, a serial number will be automatically added after the instance name.
* `lower_case_table_names` - (Optional, ForceNew) Whether table names are case-sensitive. The default value is 1. Value range:
0: Table names are case-sensitive. The backend stores them according to the actual table name.
1: (default) Table names are not case-sensitive. The backend stores them by converting table names to lowercase letters.
Description:
This rule cannot be modified after creating an instance. Please set it reasonably according to business requirements.
* `period_unit` - (Optional) Purchase cycle in prepaid scenarios.
Month: Monthly package.
Year: Annual package.
Description:
When the value of ChargeType (computing billing type) is PrePaid (monthly or annual package), this parameter is required.
* `period` - (Optional) Purchase duration in prepaid scenarios.
Description:
When the value of ChargeType (computing billing type) is PrePaid (monthly/yearly package), this parameter is required.
* `port` - (Optional, ForceNew) Specify the private network port number for the connection terminal created by default for the instance. The default value is 3306, and the value range is 1000 to 65534.
Note:
This configuration item is only effective for the primary node terminal, default terminal, and HTAP cluster terminal. That is, after the instance is created successfully, for the newly created custom terminal, the port number is still 3306 by default.
After the instance is created successfully, you can also modify the port number at any time. Currently, only modification through the console is supported.
* `pre_paid_storage_in_gb` - (Optional) Storage size in prepaid scenarios.
Description: When the value of StorageChargeType (storage billing type) is PrePaid (monthly/yearly prepaid), this parameter is required.
* `project_name` - (Optional) Project name of the instance. When this parameter is left blank, the newly created instance is added to the default project by default.
* `storage_charge_type` - (Optional) Storage billing type. When this parameter is not passed, the storage billing type defaults to be the same as the computing billing type. The values are as follows:
PostPaid: Pay-as-you-go (postpaid).
PrePaid: Monthly or yearly subscription (prepaid).
Note
When the computing billing type is PostPaid, the storage billing type can only be PostPaid.
When the computing billing type is PrePaid, the storage billing type can be PrePaid or PostPaid.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VedbMysqlInstance can be imported using the id, e.g.
```
$ terraform import volcengine_vedb_mysql_instance.default resource_id
```

