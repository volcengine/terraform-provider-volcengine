---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_allowlists"
sidebar_current: "docs-volcengine-datasource-rds_mysql_allowlists"
description: |-
  Use this data source to query detailed information of rds mysql allowlists
---
# volcengine_rds_mysql_allowlists
Use this data source to query detailed information of rds mysql allowlists
## Example Usage
```hcl
data "volcengine_rds_mysql_allowlists" "default" {
  region_id = "cn-guilin-boe"
}
```
## Argument Reference
The following arguments are supported:
* `region_id` - (Required) The region of the allow lists.
* `instance_id` - (Optional) Instance ID. When an InstanceId is specified, the DescribeAllowLists interface will return the whitelist bound to the specified instance.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `allow_lists` - The list of allowed list.
    * `allow_list_desc` - The description of the allow list.
    * `allow_list_id` - The id of the allow list.
    * `allow_list_ip_num` - The total number of IP addresses (or address ranges) in the whitelist.
    * `allow_list_name` - The name of the allow list.
    * `allow_list_type` - The type of the allow list.
    * `associated_instance_num` - The total number of instances bound under the whitelist.
    * `associated_instances` - The list of instances.
        * `instance_id` - The id of the instance.
        * `instance_name` - The name of the instance.
        * `vpc` - The id of the vpc.
* `total_count` - The total count of Scaling Activity query.


