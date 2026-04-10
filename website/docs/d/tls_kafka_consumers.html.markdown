---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_kafka_consumers"
sidebar_current: "docs-volcengine-datasource-tls_kafka_consumers"
description: |-
  Use this data source to query detailed information of tls kafka consumers
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_tls_kafka_consumers
Use this data source to query detailed information of tls kafka consumers
## Example Usage
```hcl
data "volcengine_tls_kafka_consumers" "default" {
  ids = [
    "3c57a110-399a-43b3-bc3c-5d60e065239a"
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


