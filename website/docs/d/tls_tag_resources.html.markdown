---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_tag_resources"
sidebar_current: "docs-volcengine-datasource-tls_tag_resources"
description: |-
  Use this data source to query detailed information of tls tag resources
---
# volcengine_tls_tag_resources
Use this data source to query detailed information of tls tag resources
## Example Usage
```hcl
# Basic example - query tags for specific resources
data "volcengine_tls_tag_resources" "basic" {
  resource_type = "project"
  resource_ids  = ["6e6ea17f-ee1d-494f-83f7-c3ecc5c351ea"]
  max_results   = 10
}

# Example with tag filters and max_results
data "volcengine_tls_tag_resources" "with_filters" {
  resource_type = "project"
  resource_ids  = ["project-123456", "project-789012"]
  max_results   = 50
  tag_filters {
    key    = "environment"
    values = ["production", "development"]
  }
  tag_filters {
    key    = "department"
    values = ["devops"]
  }
}


# Example with pagination using max_results
data "volcengine_tls_tag_resources" "first_page" {
  resource_type = "topic"
  resource_ids  = ["topic-123456"]
  max_results   = 20
}
```
## Argument Reference
The following arguments are supported:
* `resource_ids` - (Required) The IDs of the resources.
* `resource_type` - (Required) The type of the resource. Valid values: project, topic, shipper, host_group, host, consumer_group, rule, alarm, alarm_notify_group, etl_task, import_task, schedule_sql_task, download_task, trace_instance.
* `max_results` - (Optional) The number of results returned per page. Default value: 20. Maximum value: 100.
* `next_token` - (Optional) The token to get the next page of results. If this parameter is left empty, it means to get the first page of results.
* `output_file` - (Optional) File name where to save data source results.
* `tag_filters` - (Optional) The tag filters.

The `tag_filters` object supports the following:

* `key` - (Required) The key of the tag filter.
* `values` - (Required) The values of the tag filter.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `tags` - The list of tags.
    * `key` - The key of the tag.
    * `resource_id` - The ID of the resource.
    * `resource_type` - The type of the resource.
    * `value` - The value of the tag.


