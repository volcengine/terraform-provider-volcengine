---
subcategory: "VPN"
layout: "volcengine"
page_title: "Volcengine: volcengine_customer_gateway"
sidebar_current: "docs-volcengine-resource-customer_gateway"
description: |-
  Provides a resource to manage customer gateway
---
# volcengine_customer_gateway
Provides a resource to manage customer gateway
## Example Usage
```hcl
resource "volcengine_customer_gateway" "foo" {
  ip_address            = "192.0.1.3"
  customer_gateway_name = "acc-test"
  description           = "acc-test"
  project_name          = "default"
}
```
## Argument Reference
The following arguments are supported:
* `ip_address` - (Required, ForceNew) The IP address of the customer gateway.
* `customer_gateway_name` - (Optional) The name of the customer gateway.
* `description` - (Optional) The description of the customer gateway.
* `project_name` - (Optional) The project name of the VPN customer gateway.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account ID of the customer gateway.
* `connection_count` - The connection count of the customer gateway.
* `creation_time` - The create time of customer gateway.
* `customer_gateway_id` - The ID of the customer gateway.
* `status` - The status of the customer gateway.
* `update_time` - The update time of customer gateway.


## Import
CustomerGateway can be imported using the id, e.g.
```
$ terraform import volcengine_customer_gateway.default cgw-2byswc356dybk2dx0eed2****
```

