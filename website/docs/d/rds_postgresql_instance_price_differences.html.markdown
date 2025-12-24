---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance_price_differences"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_instance_price_differences"
description: |-
  Use this data source to query detailed information of rds postgresql instance price differences
---
# volcengine_rds_postgresql_instance_price_differences
Use this data source to query detailed information of rds postgresql instance price differences
## Example Usage
```hcl
data "volcengine_rds_postgresql_instance_price_differences" "example" {
  instance_id = "postgres-72715e0d9f58"

  modify_type = "Usually"

  node_info {
    node_id   = "postgres-72715e0d9f58"
    zone_id   = "cn-beijing-a"
    node_type = "Primary"
    node_spec = "rds.postgres.2c4g"
  }
  node_info {
    node_id   = "postgres-72715e0d9f58-iyys"
    zone_id   = "cn-beijing-a"
    node_type = "Secondary"
    node_spec = "rds.postgres.2c4g"
  }

  storage_type  = "LocalSSD"
  storage_space = 100

  charge_info {
    charge_type = "PostPaid"
    number      = 1
  }
}
```
## Argument Reference
The following arguments are supported:
* `charge_info` - (Required) Charge info of the instance.
* `instance_id` - (Required) Instance ID.
* `node_info` - (Required) Instance spec nodes. Primary=1, Secondary=1, ReadOnly=0~10.
* `storage_space` - (Required) The storage space of the instance. Value range: [20, 3000], unit: GB, step 10GB.
* `storage_type` - (Required) The type of the storage. Valid values: LocalSSD.
* `modify_type` - (Optional) Spec change type. Usually or Temporary. Default value: Usually. This parameter can only take the value Temporary when the billing type of the instance is a yearly/monthly subscription instance.
* `output_file` - (Optional) File name where to save data source results.
* `rollback_time` - (Optional) Rollback time for Temporary change, UTC format yyyy-MM-ddTHH:mm:ss.sssZ. This parameter is required when the modify_type is set to Temporary.

The `charge_info` object supports the following:

* `charge_type` - (Required) The charge type of the instance. Valid values: PostPaid, PrePaid.
* `auto_renew` - (Optional) Whether to auto renew the subscription in a pre-paid scenario.
* `number` - (Optional) Number of purchased instances. Can be an integer between 1 and 20. Default value:1.
* `period_unit` - (Optional) Purchase cycle in a pre-paid scenario. Valid values: Month, Year.
* `period` - (Optional) Subscription duration in a pre-paid scenario.Default value:1.

The `node_info` object supports the following:

* `node_spec` - (Required) The specification of the node.
* `node_type` - (Required) The type of the node. Valid values: Primary, Secondary, ReadOnly.
* `zone_id` - (Required) The AZ of the node.
* `node_id` - (Optional) The id of the node.When the modify_type is set to Temporary, this parameter is required.
* `node_operate_type` - (Optional) The operate type of the node. Valid values: Create, Modify.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instances_price` - The collection of query.
    * `currency` - Currency unit.
    * `discount_price` - Instance price after discount.
    * `original_price` - Instance price before discount.
    * `payable_price` - Price payable of instance.
* `total_count` - The total count of query.


