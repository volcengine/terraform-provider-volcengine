---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_service_permissions"
sidebar_current: "docs-volcengine-datasource-privatelink_vpc_endpoint_service_permissions"
description: |-
  Use this data source to query detailed information of privatelink vpc endpoint service permissions
---
# volcengine_privatelink_vpc_endpoint_service_permissions
Use this data source to query detailed information of privatelink vpc endpoint service permissions
## Example Usage
```hcl
data "volcengine_privatelink_vpc_endpoint_service_permissions" "default" {
  service_id = "epsvc-3rel73uf2ewao5zsk2j2l58ro"
}
```
## Argument Reference
The following arguments are supported:
* `service_id` - (Required) The Id of service.
* `output_file` - (Optional) File name where to save data source results.
* `permit_account_id` - (Optional) The Id of permit account.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `permissions` - The collection of query.
    * `permit_account_id` - The permit account id.
* `total_count` - Returns the total amount of the data list.


