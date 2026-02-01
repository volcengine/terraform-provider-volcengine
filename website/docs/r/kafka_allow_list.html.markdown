---
subcategory: "KAFKA"
layout: "volcengine"
page_title: "Volcengine: volcengine_kafka_allow_list"
sidebar_current: "docs-volcengine-resource-kafka_allow_list"
description: |-
  Provides a resource to manage kafka allow list
---
# volcengine_kafka_allow_list
Provides a resource to manage kafka allow list
## Example Usage
```hcl
resource "volcengine_kafka_allow_list" "foo" {
  allow_list      = ["192.168.0.1", "10.32.55.66", "10.22.55.66"]
  allow_list_name = "tf-test"
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_name` - (Required) The name of the allow list.
* `allow_list` - (Required) Whitelist rule list. Supports specifying as IP addresses or IP network segments. Each whitelist can be configured with a maximum of 300 IP addresses or network segments.
* `allow_list_desc` - (Optional) The description of the allow list.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
KafkaAllowList can be imported using the id, e.g.
```
$ terraform import volcengine_kafka_allow_list.default resource_id
```

