---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_consumer_group"
sidebar_current: "docs-volcengine-resource-tls_consumer_group"
description: |-
  Provides a resource to manage tls consumer group
---
# volcengine_tls_consumer_group
Provides a resource to manage tls consumer group
## Example Usage
```hcl
resource "volcengine_tls_consumer_group" "foo" {
  project_id          = "17ba378d-de43-495e-8906-03aexxxxxx"
  topic_id_list       = ["0ed72ac8-9531-4967-b216-ac30xxxxxx"]
  consumer_group_name = "tf-test-consumer-group"
  heartbeat_ttl       = 120
  ordered_consume     = false
}
```
## Argument Reference
The following arguments are supported:
* `consumer_group_name` - (Required, ForceNew) The name of the consumer group.
* `heartbeat_ttl` - (Required) The time of heart rate expiration, measured in seconds, has a value range of 1 to 300.
* `ordered_consume` - (Required) Whether to consume in sequence.
* `project_id` - (Required, ForceNew) The log project ID to which the consumption group belongs.
* `topic_id_list` - (Required) The list of log topic ids to be consumed by the consumer group.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ConsumerGroup can be imported using the id, e.g.
```
$ terraform import volcengine_consumer_group.default resource_id
```

