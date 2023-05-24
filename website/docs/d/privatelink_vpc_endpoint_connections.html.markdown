---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_connections"
sidebar_current: "docs-volcengine-datasource-privatelink_vpc_endpoint_connections"
description: |-
  Use this data source to query detailed information of privatelink vpc endpoint connections
---
# volcengine_privatelink_vpc_endpoint_connections
Use this data source to query detailed information of privatelink vpc endpoint connections
## Example Usage
```hcl
data "volcengine_privatelink_vpc_endpoint_connections" "default" {
  endpoint_id = "ep-3rel74u229dz45zsk2i6l69qa"
  service_id  = "epsvc-2byz5mykk9y4g2dx0efs4aqz3"
}
```
## Argument Reference
The following arguments are supported:
* `service_id` - (Required) The id of the vpc endpoint service.
* `endpoint_id` - (Optional) The id of the vpc endpoint.
* `endpoint_owner_account_id` - (Optional) The account id of the vpc endpoint.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `connections` - The list of query.
    * `connection_status` - The status of the connection.
    * `creation_time` - The create time of the connection.
    * `endpoint_id` - The id of the vpc endpoint.
    * `endpoint_owner_account_id` - The account id of the vpc endpoint.
    * `endpoint_vpc_id` - The vpc id of the vpc endpoint.
    * `service_id` - The id of the vpc endpoint service.
    * `update_time` - The update time of the connection.
    * `zones` - The available zones.
        * `network_interface_id` - The id of the network interface.
        * `network_interface_ip` - The ip address of the network interface.
        * `resource_id` - The id of the resource.
        * `subnet_id` - The id of the subnet.
        * `zone_domain` - The domain of the zone.
        * `zone_id` - The id of the zone.
        * `zone_status` - The status of the zone.
* `total_count` - Returns the total amount of the data list.


