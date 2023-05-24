---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_service_permission"
sidebar_current: "docs-volcengine-resource-privatelink_vpc_endpoint_service_permission"
description: |-
  Provides a resource to manage privatelink vpc endpoint service permission
---
# volcengine_privatelink_vpc_endpoint_service_permission
Provides a resource to manage privatelink vpc endpoint service permission
## Example Usage
```hcl
resource "volcengine_privatelink_vpc_endpoint_service_permission" "foo" {
  service_id        = "epsvc-3rel73uf2ewao5zsk2j2l58ro"
  permit_account_id = "210000000"
}

resource "volcengine_privatelink_vpc_endpoint_service_permission" "foo1" {
  service_id        = "epsvc-3rel73uf2ewao5zsk2j2l58ro"
  permit_account_id = "210000001"
}
```
## Argument Reference
The following arguments are supported:
* `permit_account_id` - (Required, ForceNew) The id of account.
* `service_id` - (Required, ForceNew) The id of service.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VpcEndpointServicePermission can be imported using the serviceId:permitAccountId, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint_service_permission.default epsvc-2fe630gurkl37k5gfuy33****:2100000000
```

