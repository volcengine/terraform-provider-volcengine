---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_health_check_log_topics"
sidebar_current: "docs-volcengine-datasource-health_check_log_topics"
description: |-
  Use this data source to query detailed information of health check log topics
---
# volcengine_health_check_log_topics
Use this data source to query detailed information of health check log topics
## Example Usage
```hcl
data "volcengine_health_check_log_topics" "example" {
  log_topic_id = "82fddbd8-4140-4527-****-b89d2aae4a61"
}
```
## Argument Reference
The following arguments are supported:
* `log_topic_id` - (Required) The ID of the log topic.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `health_check_log_topics` - The collection of query.
    * `load_balancer_ids` - The ID of the CLB instance.
* `total_count` - The total count of query.


