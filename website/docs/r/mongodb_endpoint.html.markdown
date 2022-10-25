---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_endpoint"
sidebar_current: "docs-volcengine-resource-mongodb_endpoint"
description: |-
  Provides a resource to manage mongodb endpoint
---
# volcengine_mongodb_endpoint
Provides a resource to manage mongodb endpoint
## Example Usage
```hcl
resource "volcengine_mongodb_endpoint" "foo" {
  instance_id  = "mongo-shard-xxx"
  object_id    = "mongo-shard-xxx-s1"
  network_type = "Public"
  eip_ids      = ["eip-xx", "eip-xx"]
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required, ForceNew) The instance where the endpoint resides.
* `eip_ids` - (Optional) A list of EIP IDs that need to be bound when applying for endpoint.
* `mongos_node_ids` - (Optional) A list of the Mongos node that needs to apply for the endpoint.
* `network_type` - (Optional) The network type of endpoint.
* `object_id` - (Optional) The object ID corresponding to the endpoint.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `endpoint_id` - The id of endpoint.


## Import
mongodb endpoint can be imported using the instanceId:endpointId
`instanceId`: represents the instance that endpoint related to.
`endpointId`: the id of endpoint.
e.g.
```
$ terraform import volcengine_mongodb_endpoint.default mongo-replica-e405f8e2****:BRhFA0pDAk0XXkxCZQ
```

