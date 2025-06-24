---
subcategory: "KAFKA"
layout: "volcengine"
page_title: "Volcengine: volcengine_kafka_allow_lists"
sidebar_current: "docs-volcengine-datasource-kafka_allow_lists"
description: |-
  Use this data source to query detailed information of kafka allow lists
---
# volcengine_kafka_allow_lists
Use this data source to query detailed information of kafka allow lists
## Example Usage
```hcl
data "volcengine_kafka_allow_lists" "foo" {
  instance_id = "kafka-xxx"
  region_id   = "cn-beijing"
}
```
## Argument Reference
The following arguments are supported:
* `region_id` - (Required) The region ID.
* `instance_id` - (Optional) The instance ID to query.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `allow_lists` - The collection of query.
    * `allow_list_desc` - The description of the allow list.
    * `allow_list_id` - The id of the allow list.
    * `allow_list_ip_num` - The number of rules specified in the whitelist.
    * `allow_list_name` - The name of the allow list.
    * `allow_list` - Whitelist rule list.
    * `associated_instance_num` - The number of instances bound to the whitelist.
    * `associated_instances` - The list of associated instances.
        * `instance_id` - The id of the instance.
        * `instance_name` - The name of the instance.
* `total_count` - The total count of query.


