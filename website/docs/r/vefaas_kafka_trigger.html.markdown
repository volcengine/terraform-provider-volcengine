---
subcategory: "VEFAAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_vefaas_kafka_trigger"
sidebar_current: "docs-volcengine-resource-vefaas_kafka_trigger"
description: |-
  Provides a resource to manage vefaas kafka trigger
---
# volcengine_vefaas_kafka_trigger
Provides a resource to manage vefaas kafka trigger
## Example Usage
```hcl
resource "volcengine_vefaas_kafka_trigger" "foo" {
  function_id    = "35ybaxxx"
  name           = "tf-123"
  mq_instance_id = "kafka-cnngmbeq10mcxxxx"
  topic_name     = "topic"
  kafka_credentials {
    password  = "Waxxxxxx"
    username  = "test-1"
    mechanism = "PLAIN"
  }
  batch_size  = 100
  description = "modify"
  lifecycle {
    ignore_changes = [kafka_credentials]
  }
}
```
## Argument Reference
The following arguments are supported:
* `function_id` - (Required, ForceNew) The ID of Function.
* `kafka_credentials` - (Required, ForceNew) Kafka identity authentication. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `mq_instance_id` - (Required, ForceNew) The instance ID of Message queue Kafka.
* `name` - (Required, ForceNew) The name of the Kafka trigger.
* `topic_name` - (Required, ForceNew) The Topic name of the message queue Kafka instance.
* `batch_flush_duration_milliseconds` - (Optional) The maximum waiting time for batch consumption of triggers.
* `batch_size` - (Optional) The number of messages per batch consumed by the trigger in bulk.
* `description` - (Optional) The description of the Kafka trigger.
* `enabled` - (Optional) Whether to enable triggers at the same time as creating them.
* `maximum_retry_attempts` - (Optional) The maximum number of retries when a function has a runtime error.
* `starting_position` - (Optional, ForceNew) Specify the location where the messages in the Topic start to be consumed.

The `kafka_credentials` object supports the following:

* `mechanism` - (Required) Kafka authentication mechanism.
* `password` - (Required) The SASL/PLAIN user password set when creating a Kafka instance.
* `username` - (Required) The SASL/PLAIN user name set when creating a Kafka instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `consumer_group` - The consumer group name of the message queue Kafka instance.
* `creation_time` - The creation time of the Kafka trigger.
* `last_update_time` - The last update time of the Kafka trigger.
* `status` - The status of Kafka trigger.


## Import
VefaasKafkaTrigger can be imported using the id, e.g.
```
$ terraform import volcengine_vefaas_kafka_trigger.default resource_id
```

