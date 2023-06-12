---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_connection"
sidebar_current: "docs-volcengine-resource-privatelink_vpc_endpoint_connection"
description: |-
  Provides a resource to manage privatelink vpc endpoint connection
---
# volcengine_privatelink_vpc_endpoint_connection
Provides a resource to manage privatelink vpc endpoint connection
## Example Usage
```hcl
resource "volcengine_privatelink_vpc_endpoint_connection" "foo" {
  endpoint_id = "ep-3rel74u229dz45zsk2i6l69qa"
  service_id  = "epsvc-2byz5mykk9y4g2dx0efs4aqz3"
}
```
## Argument Reference
The following arguments are supported:
* `endpoint_id` - (Required, ForceNew) The id of the endpoint.
* `service_id` - (Required, ForceNew) The id of the security group.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `connection_status` - The status of the connection.
* `creation_time` - The create time of the connection.
* `endpoint_owner_account_id` - The account id of the vpc endpoint.
* `endpoint_vpc_id` - The vpc id of the vpc endpoint.
* `update_time` - The update time of the connection.
* `zones` - The available zones.
    * `network_interface_id` - The id of the network interface.
    * `network_interface_ip` - The ip address of the network interface.
    * `resource_id` - The id of the resource.
    * `subnet_id` - The id of the subnet.
    * `zone_domain` - The domain of the zone.
    * `zone_id` - The id of the zone.
    * `zone_status` - The status of the zone.


## Import
PrivateLink Vpc Endpoint Connection Service can be imported using the endpoint id and service id, e.g.
```
$ terraform import volcengine_privatelink_vpc_endpoint_connection.default ep-3rel74u229dz45zsk2i6l69qa:epsvc-2byz5mykk9y4g2dx0efs4aqz3
```

