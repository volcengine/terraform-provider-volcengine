---
subcategory: "CLOUD_MONITOR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_monitor_contacts"
sidebar_current: "docs-volcengine-datasource-cloud_monitor_contacts"
description: |-
  Use this data source to query detailed information of cloud monitor contacts
---
# volcengine_cloud_monitor_contacts
Use this data source to query detailed information of cloud monitor contacts
## Example Usage
```hcl
data "volcengine_cloud_monitor_contacts" "foo" {
  ids = ["17******516", "1712**********0"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Required) A list of Contact IDs.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `contacts` - The collection of query.
    * `email` - The email of contact.
    * `id` - The ID of contact.
    * `name` - The name of contact.
    * `phone` - The phone of contact.
* `total_count` - The total count of query.


