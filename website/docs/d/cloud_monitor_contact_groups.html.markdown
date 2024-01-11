---
subcategory: "CLOUD_MONITOR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_monitor_contact_groups"
sidebar_current: "docs-volcengine-datasource-cloud_monitor_contact_groups"
description: |-
  Use this data source to query detailed information of cloud monitor contact groups
---
# volcengine_cloud_monitor_contact_groups
Use this data source to query detailed information of cloud monitor contact groups
## Example Usage
```hcl
data "volcengine_cloud_monitor_contact_groups" "foo" {
  name = "tftest"
}
```
## Argument Reference
The following arguments are supported:
* `name` - (Optional) Search for keywords in contact group names, supports fuzzy search.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `groups` - The collection of query.
    * `account_id` - The id of the account.
    * `contacts` - Contact information in the contact group.
        * `email` - The email of contact.
        * `id` - The id of the contact.
        * `name` - The name of contact.
        * `phone` - The phone of contact.
    * `created_at` - The create time.
    * `description` - The description of the contact group.
    * `id` - The id of the contact group.
    * `name` - The name of the contact group.
    * `updated_at` - The update time.
* `total_count` - The total count of query.


