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
data "volcengine_rds_postgresql_allowlists" "default" {
  name_regex          = ".*allowlist.*"
  allow_list_id       = "acl-e7846436e1e741edbd385868fa657436"
  allow_list_category = "Ordinary"
  allow_list_desc     = "test allow list"
  allow_list_name     = "test"
  ip_address          = "100.64.0.0/10"
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_category` - (Optional) The category of the postgresql allow list. Valid values: Ordinary, Default.
* `allow_list_desc` - (Optional) The description of the postgresql allow list. Perform a fuzzy search based on the description information.
* `allow_list_id` - (Optional) The id of the postgresql allow list.
* `allow_list_name` - (Optional) The name of the postgresql allow list.
* `instance_id` - (Optional) The id of the postgresql Instance.
* `ip_address` - (Optional) The IP address to be added to the allow list.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `postgresql_allow_lists` - The list of postgresql allowed list.
    * `allow_list_category` - The category of the postgresql allow list.
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
    * `security_group_bind_infos` - The information of the security group bound by the allowlist.
        * `bind_mode` - The binding mode of the security group.
        * `ip_list` - IP addresses in the security group.
        * `security_group_id` - The ID of the security group.
        * `security_group_name` - The name of the security group.
    * `user_allow_list` - IP addresses outside the security group and added to the allowlist.
* `total_count` - The total count of query.


