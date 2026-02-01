---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_kafka_consumers"
sidebar_current: "docs-volcengine-datasource-tls_kafka_consumers"
description: |-
  Use this data source to query detailed information of tls kafka consumers
---
# volcengine_tls_kafka_consumers
Use this data source to query detailed information of tls kafka consumers
## Example Usage
```hcl
data "volcengine_tls_kafka_consumers" "default" {
  ids = [
    "65d67d34-c5b4-4ec8-b3a9-175d33668b45", "cfb5c08b-0c7a-44fa-8971-8afc12f1b123",
    "edf051ed-3c46-49ba-9339-bea628fedc15"
  ]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of topic IDs.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `data` - The collection of query.
    * `allow_consume` - Whether allow consume.
    * `consume_topic` - The topic of consume.
    * `topic_id` - The ID of Topic.
* `total_count` - The total count of query.


