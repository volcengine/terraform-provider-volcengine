---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_schema"
sidebar_current: "docs-volcengine-resource-rds_postgresql_schema"
description: |-
  Provides a resource to manage rds postgresql schema
---
# volcengine_rds_postgresql_schema
Provides a resource to manage rds postgresql schema
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-project1"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-subnet-test-2"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}


resource "volcengine_rds_postgresql_instance" "foo" {
  db_engine_version = "PostgreSQL_12"
  node_spec         = "rds.postgres.1c2g"
  primary_zone_id   = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id = data.volcengine_zones.foo.zones[0].id
  storage_space     = 40
  subnet_id         = volcengine_subnet.foo.id
  instance_name     = "acc-test-1"
  charge_info {
    charge_type = "PostPaid"
  }
  project_name = "default"
  tags {
    key   = "tfk1"
    value = "tfv1"
  }
  parameters {
    name  = "auto_explain.log_analyze"
    value = "off"
  }
  parameters {
    name  = "auto_explain.log_format"
    value = "text"
  }
}

resource "volcengine_rds_postgresql_database" "foo" {
  db_name     = "acc-test"
  instance_id = volcengine_rds_postgresql_instance.foo.id
  c_type      = "C"
  collate     = "zh_CN.utf8"
}

resource "volcengine_rds_postgresql_account" "foo" {
  account_name       = "acc-test-account"
  account_password   = "9wc@********12"
  account_type       = "Normal"
  instance_id        = volcengine_rds_postgresql_instance.foo.id
  account_privileges = "Inherit,Login,CreateRole,CreateDB"
}

resource "volcengine_rds_postgresql_account" "foo1" {
  account_name       = "acc-test-account1"
  account_password   = "9wc@*******12"
  account_type       = "Normal"
  instance_id        = volcengine_rds_postgresql_instance.foo.id
  account_privileges = "Inherit,Login,CreateRole,CreateDB"
}

resource "volcengine_rds_postgresql_schema" "foo" {
  db_name     = volcengine_rds_postgresql_database.foo.db_name
  instance_id = volcengine_rds_postgresql_instance.foo.id
  owner       = volcengine_rds_postgresql_account.foo.account_name
  schema_name = "acc-test-schema"
}
```
## Argument Reference
The following arguments are supported:
* `db_name` - (Required, ForceNew) The name of the database.
* `instance_id` - (Required, ForceNew) The id of the postgresql instance.
* `owner` - (Required) The owner of the schema.The instance read-only account, a high-privilege account with DDL permissions disabled, or a normal account with DDL permissions disabled cannot be used as the owner of the schema.
* `schema_name` - (Required, ForceNew) The name of the schema.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
RdsPostgresqlSchema can be imported using the instance id, database name and schema name, e.g.
```
$ terraform import volcengine_rds_postgresql_schema.default instance_id:db_name:schema_name
```

