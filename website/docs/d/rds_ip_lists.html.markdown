---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_ip_lists"
sidebar_current: "docs-volcengine-datasource-rds_ip_lists"
description: |-
  Use this data source to query detailed information of rds ip lists
---
# volcengine_rds_ip_lists
(Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds ip lists
## Example Usage
```hcl
data "volcengine_rds_ip_lists" "default" {
  instance_id = "mysql-0fdd3bab2e7c"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of the RDS instance.
* `name_regex` - (Optional) A Name Regex of RDS ip list.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rds_ip_lists` - The collection of RDS ip list account query.
    * `group_name` - The name of the RDS ip list.
    * `id` - The ID of the RDS ip list.
    * `ip_list` - The list of IP address.
* `total_count` - The total count of RDS ip list query.


