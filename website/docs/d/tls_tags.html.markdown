---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_tags"
sidebar_current: "docs-volcengine-datasource-tls_tags"
description: |-
  Use this data source to query detailed information of tls tags
---
# volcengine_tls_tags
Use this data source to query detailed information of tls tags
## Example Usage
```hcl
# Basic example - query tags for specific resources
data "volcengine_tls_tags" "basic" {
  resource_type = "project"
  resource_ids  = ["b01a99c0-cf7b-482f-b317-6563865111c6"]
  max_results   = 10
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


