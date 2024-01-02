---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_service_resource"
sidebar_current: "docs-volcengine-resource-privatelink_vpc_endpoint_service_resource"
description: |-
  Provides a resource to manage privatelink vpc endpoint service resource
---
# volcengine_privatelink_vpc_endpoint_service_resource
Provides a resource to manage privatelink vpc endpoint service resource
## Example Usage
```hcl
resource "volcengine_privatelink_vpc_endpoint_service_resource" "foo" {
  service_id  = "epsvc-3rel73uf2ewao5zsk2j2l58ro"
  resource_id = "clb-3reii8qfbp7gg5zsk2hsrbe3c"
}

resource "volcengine_privatelink_vpc_endpoint_service_resource" "foo1" {
  service_id  = "epsvc-3rel73uf2ewao5zsk2j2l58ro"
  resource_id = "clb-2d6sfye98rzls58ozfducee1o"
}

resource "volcengine_privatelink_vpc_endpoint_service_resource" "foo2" {
  service_id  = "epsvc-3rel73uf2ewao5zsk2j2l58ro"
  resource_id = "clb-3refkvae02gow5zsk2ilaev5y"
}
```
## Argument Reference
The following arguments are supported:
* `resource_id` - (Required, ForceNew) The id of resource. It is not recommended to use this resource for binding resources, it is recommended to use the resources field of vpc_endpoint_service for binding. If using this resource and vpc_endpoint_service jointly for operations, use lifecycle ignore_changes to suppress changes to the resources field in vpc_endpoint_service.
* `service_id` - (Required, ForceNew) The id of service.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VpcEndpointServiceResource can be imported using the serviceId:resourceId, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint_service_resource.default epsvc-2fe630gurkl37k5gfuy33****:clb-bp1o94dp5i6ea****
```
It is not recommended to use this resource for binding resources, it is recommended to use the resources field of vpc_endpoint_service for binding.
If using this resource and vpc_endpoint_service jointly for operations, use lifecycle ignore_changes to suppress changes to the resources field in vpc_endpoint_service.

