---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_kafka_consumer"
sidebar_current: "docs-volcengine-resource-tls_kafka_consumer"
description: |-
  Provides a resource to manage tls kafka consumer
---
# volcengine_tls_kafka_consumer
Provides a resource to manage tls kafka consumer
## Example Usage
```hcl
resource "volcengine_tls_kafka_consumer" "foo" {
  topic_id = "cfb5c08b-0c7a-44fa-8971-8afc12f1b123"
}
```
## Argument Reference
The following arguments are supported:
* `topic_id` - (Required, ForceNew) The id of topic.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `allow_consume` - Whether allow consume.
* `consume_topic` - The topic of consume.


## Import
Tls Kafka Consumer can be imported using the kafka:topic_id, e.g.
```
$ terraform import volcengine_tls_kafka_consumer.default kafka:edf051ed-3c46-49ba-9339-bea628fedc15
```

