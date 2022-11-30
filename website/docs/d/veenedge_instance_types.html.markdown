---
subcategory: "VEENEDGE"
layout: "volcengine"
page_title: "Volcengine: volcengine_veenedge_instance_types"
sidebar_current: "docs-volcengine-datasource-veenedge_instance_types"
description: |-
  Use this data source to query detailed information of veenedge instance types
---
# volcengine_veenedge_instance_types
Use this data source to query detailed information of veenedge instance types
## Example Usage
```hcl
data "volcengine_veenedge_instance_types" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `instance_type_configs` - The collection of instance types query.
    * `cpu` - The cpu of instance type.
    * `gpu_spec` - The gpu spec of instance.
    * `gpu` - The gpu of instance type.
    * `instance_type_family_name` - The name of instance type family.
    * `instance_type_family` - The type family of instance.
    * `instance_type` - The type of instance.
    * `memory` - The memory of instance type.
    * `storage` - The config of storage.
        * `local_storage_amount` - The amount of local storage.
        * `local_storage_capacity` - The capacity of local storage.
        * `local_storage_category` - The local storage category.
        * `local_storage_unit` - The unit of local storage.
* `total_count` - The total count of instance types query.


