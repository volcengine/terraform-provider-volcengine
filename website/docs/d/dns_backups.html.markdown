---
subcategory: "DNS"
layout: "volcengine"
page_title: "Volcengine: volcengine_dns_backups"
sidebar_current: "docs-volcengine-datasource-dns_backups"
description: |-
  Use this data source to query detailed information of dns backups
---
# volcengine_dns_backups
Use this data source to query detailed information of dns backups
## Example Usage
```hcl
data "volcengine_dns_backups" "foo" {
  zid = 58846
}
```
## Argument Reference
The following arguments are supported:
* `zid` - (Required) The ID of the domain for which you want to get the backup schedule.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `backup_infos` - The collection of query.
    * `backup_id` - The ID of the backup.
    * `backup_time` - The time when the backup was created. The time zone is UTC + 8.
    * `record_count` - The number of DNS records in the backup.
* `total_count` - The total count of query.


