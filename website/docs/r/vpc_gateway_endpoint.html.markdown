---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_vpc_gateway_endpoint"
sidebar_current: "docs-volcengine-resource-vpc_gateway_endpoint"
description: |-
  Provides a resource to manage vpc gateway endpoint
---
# volcengine_vpc_gateway_endpoint
Provides a resource to manage vpc gateway endpoint
## Example Usage
```hcl

```
## Argument Reference
The following arguments are supported:
* `service_id` - (Required, ForceNew) The id of the gateway endpoint service.
* `vpc_id` - (Required, ForceNew) The id of the vpc.
* `description` - (Optional) The description of the gateway endpoint.
* `endpoint_name` - (Optional) The name of the gateway endpoint.
* `project_name` - (Optional) The project name of the gateway endpoint.
* `tags` - (Optional) Tags.
* `vpc_policy` - (Optional) The vpc policy of the gateway endpoint.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `endpoint_id` - The id of the gateway endpoint.


## Import
VpcGatewayEndpoint can be imported using the id, e.g.
```
$ terraform import volcengine_vpc_gateway_endpoint.default gwep-273yuq6q7bgn47fap8squ****
```

