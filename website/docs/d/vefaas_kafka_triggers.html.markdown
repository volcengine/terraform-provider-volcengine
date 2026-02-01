---
subcategory: "VEFAAS"
layout: "volcengine"
page_title: "Volcengine: volcengine_vefaas_kafka_triggers"
sidebar_current: "docs-volcengine-datasource-vefaas_kafka_triggers"
description: |-
  Use this data source to query detailed information of vefaas kafka triggers
---
# volcengine_vefaas_kafka_triggers
Use this data source to query detailed information of vefaas kafka triggers
## Example Usage
```hcl
data "volcengine_vefaas_kafka_triggers" "foo" {
  function_id = "f0zvcxxx"
}
```
## Argument Reference
The following arguments are supported:
* `function_id` - (Required) The ID of Function.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `items` - The list of timer trigger.
    * `creation_time` - The creation time of the Timer trigger.
    * `description` - The description of the Timer trigger.
    * `detailed_config` - The details of trigger configuration.
    * `enabled` - Whether the Timer trigger is enabled.
    * `function_id` - The ID of Function.
    * `id` - The ID of the Timer trigger.
    * `image_version` - The image version of the Timer trigger.
    * `last_update_time` - The last update time of the Timer trigger.
    * `name` - The name of the Timer trigger.
    * `type` - The category of the Timer trigger.
* `total_count` - The total count of query.


