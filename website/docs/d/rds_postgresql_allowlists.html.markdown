---
subcategory: "RDS_POSTGRESQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_postgresql_allowlists"
sidebar_current: "docs-volcengine-datasource-rds_postgresql_allowlists"
description: |-
  Use this data source to query detailed information of rds postgresql allowlists
---
# volcengine_rds_postgresql_allowlists
Use this data source to query detailed information of rds postgresql allowlists
## Example Usage
```hcl
data "volcengine_rds_postgresql_allowlists" "foo" {

}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Optional) The id of the postgresql Instance.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `postgresql_allow_lists` - The list of postgresql allowed list.
    * `allow_list_desc` - The description of the postgresql allow list.
    * `allow_list_id` - The id of the postgresql allow list.
    * `allow_list_ip_num` - The total number of IP addresses (or address ranges) in the whitelist.
    * `allow_list_name` - The name of the postgresql allow list.
    * `allow_list_type` - The type of the postgresql allow list.
    * `allow_list` - The IP address or a range of IP addresses in CIDR format.
    * `associated_instance_num` - The total number of instances bound under the whitelist.
    * `associated_instances` - The list of postgresql instances.
        * `instance_id` - The id of the postgresql instance.
        * `instance_name` - The name of the postgresql instance.
        * `vpc` - The id of the vpc.
    * `id` - The id of the postgresql allow list.
* `total_count` - The total count of query.


