---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_service"
sidebar_current: "docs-volcengine-resource-privatelink_vpc_endpoint_service"
description: |-
  Provides a resource to manage privatelink vpc endpoint service
---
# volcengine_privatelink_vpc_endpoint_service
Provides a resource to manage privatelink vpc endpoint service
## Example Usage
```hcl
resource "volcengine_privatelink_vpc_endpoint_service" "foo" {
  resources {
    resource_id   = "clb-2bzxccdjo9uyo2dx0eg0orzla"
    resource_type = "CLB"
  }
  description         = "tftest"
  auto_accept_enabled = true
}
```
## Argument Reference
The following arguments are supported:
* `resources` - (Required) The resources info. When create vpc endpoint service, the resource must exist. It is recommended to bind resources using the resources' field in this resource instead of using vpc_endpoint_service_resource. For operations that jointly use this resource and vpc_endpoint_service_resource, use lifecycle ignore_changes to suppress changes to the resources field.
* `auto_accept_enabled` - (Optional) Whether auto accept node connect.
* `description` - (Optional) The description of service.

The `resources` object supports the following:

* `resource_id` - (Required) The id of resource.
* `resource_type` - (Required) The type of resource.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The create time of service.
* `service_domain` - The domain of service.
* `service_id` - The Id of service.
* `service_name` - The name of service.
* `service_resource_type` - The resource type of service.
* `service_type` - The type of service.
* `status` - The status of service.
* `update_time` - The update time of service.
* `zone_ids` - The IDs of zones.


## Import
VpcEndpointService can be imported using the id, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint_service.default epsvc-2fe630gurkl37k5gfuy33****
```
It is recommended to bind resources using the resources' field in this resource instead of using vpc_endpoint_service_resource.
For operations that jointly use this resource and vpc_endpoint_service_resource, use lifecycle ignore_changes to suppress changes to the resources field.

