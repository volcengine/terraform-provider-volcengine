---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_instance_price_details"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_instance_price_details"
description: |-
  Use this data source to query detailed information of rds postgresql instance price details
---
# volcengine_rds_postgresql_instance_price_details
Use this data source to query detailed information of rds postgresql instance price details
## Example Usage
```hcl
data "volcengine_rds_postgresql_instance_price_details" "example" {
  node_info {
    zone_id           = "cn-beijing-a"
    node_type         = "Primary"
    node_spec         = "rds.postgres.1c2g"
    node_operate_type = "Create"
  }
  node_info {
    zone_id           = "cn-beijing-a"
    node_type         = "Secondary"
    node_spec         = "rds.postgres.1c2g"
    node_operate_type = "Create"
  }
  node_info {
    zone_id           = "cn-beijing-a"
    node_type         = "ReadOnly"
    node_spec         = "rds.postgres.2c8g"
    node_operate_type = "Create"
  }
  storage_type  = "LocalSSD"
  storage_space = 100

  charge_info {
    charge_type = "PrePaid"
    period_unit = "Month"
    period      = 2
    number      = 4
  }
}
```
## Argument Reference
The following arguments are supported:
* `charge_info` - (Required) The charge information of the instance.
* `node_info` - (Required) Instance specification configuration. An instance must have only one primary node, only one secondary node, and 0~10 read-only nodes.
* `storage_space` - (Required) The storage space of the instance. Value range: [20, 3000], unit: GB, step 10GB.
* `storage_type` - (Required) The type of the storage. Valid values: LocalSSD.
* `output_file` - (Optional) File name where to save data source results.

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
* `node_id` - (Optional) The id of the node.
* `node_operate_type` - (Optional) The operate type of the node. Valid values: Create.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instances_price` - The collection of query.
    * `charge_item_prices` - Price of each charge item.
        * `charge_item_key` - If charge_item_key is Primary, Secondary, or ReadOnly, this parameter returns the instance specification, such as rds.pg.d1.1c2g. If charge_item_key is Storage, this parameter returns the stored key, such as rds.pg.d1.localssd.
        * `charge_item_type` - Billing item name. Values:Primary, Secondary, ReadOnly, Storage.
        * `charge_item_value` - If charge_item_key is Primary, Secondary, or ReadOnly, this parameter returns the number of nodes, with a value of "1". If charge_item_key is Storage, his parameter returns the storage size in GB.
        * `discount_price` - Discount price of each charge item.
        * `node_num_per_instance` - Number of nodes of each instance.
        * `original_price` - Original price of each charge item.
        * `payable_price` - Payable price of each charge item.
    * `currency` - Currency unit.
    * `discount_price` - Instance price after discount.
    * `instance_quantity` - Number of purchased instances.
    * `original_price` - Instance price before discount.
    * `payable_price` - Price payable of instance.
* `total_count` - The total count of query.


