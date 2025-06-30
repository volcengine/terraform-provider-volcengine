---
subcategory: "KAFKA"
layout: "volcengine"
page_title: "Volcengine: volcengine_kafka_allow_list_associate"
sidebar_current: "docs-volcengine-resource-kafka_allow_list_associate"
description: |-
  Provides a resource to manage kafka allow list associate
---
# volcengine_kafka_allow_list_associate
Provides a resource to manage kafka allow list associate
## Example Usage
```hcl
resource "volcengine_kafka_allow_list" "foo" {
  allow_list      = ["192.168.0.1", "10.32.55.66", "10.22.55.66"]
  allow_list_name = "tf-test"
}

resource "volcengine_kafka_allow_list_associate" "foo" {
  allow_list_id = volcengine_kafka_allow_list.foo.id
  instance_id   = "kafka-cnoex9j4un63uqjr"
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_id` - (Required, ForceNew) The id of the allow list.
* `instance_id` - (Required, ForceNew) The id of the kafka instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
KafkaAllowListAssociate can be imported using the id, e.g.
```
$ terraform import volcengine_kafka_allow_list_associate.default kafka-cnitzqgn**:acl-d1fd76693bd54e658912e7337d5b****
```

