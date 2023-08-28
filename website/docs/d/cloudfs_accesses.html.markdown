---
subcategory: "CLOUDFS"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloudfs_accesses"
sidebar_current: "docs-volcengine-datasource-cloudfs_accesses"
description: |-
  Use this data source to query detailed information of cloudfs accesses
---
# volcengine_cloudfs_accesses
Use this data source to query detailed information of cloudfs accesses
## Example Usage
```hcl
data "volcengine_cloudfs_accesses" "default" {
  fs_name = "tftest2"
}
```
## Argument Reference
The following arguments are supported:
* `fs_name` - (Required) The name of file system.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `accesses` - The collection of query.
    * `access_account_id` - The account id of access.
    * `access_id` - The id of access.
    * `access_service_name` - The service name of access.
    * `created_time` - The creation time.
    * `fs_name` - The name of cloud fs.
    * `is_default` - Whether is default access.
    * `security_group_id` - The id of security group.
    * `status` - The status of access.
    * `subnet_id` - The id of subnet.
    * `vpc_id` - The id of vpc.
    * `vpc_route_enabled` - Whether to enable all vpc route.
* `total_count` - The total count of query.


