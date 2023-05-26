---
subcategory: "REDIS"
layout: "volcengine"
page_title: "Volcengine: volcengine_redis_allow_lists"
sidebar_current: "docs-volcengine-datasource-redis_allow_lists"
description: |-
  Use this data source to query detailed information of redis allow lists
---
# volcengine_redis_allow_lists
Use this data source to query detailed information of redis allow lists
## Example Usage
```hcl
data "volcengine_redis_allow_lists" "default" {
  region_id = "cn-beijing"
}
```
## Argument Reference
The following arguments are supported:
* `region_id` - (Required) The Id of region.
* `instance_id` - (Optional) The Id of instance.
* `name_regex` - (Optional) A Name Regex of Allow List.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `allow_lists` - Information of list of allow list.
    * `allow_list_desc` - Description of allow list.
    * `allow_list_id` - Id of allow list.
    * `allow_list_ip_num` - The IP number of allow list.
    * `allow_list_name` - Name of allow list.
    * `allow_list_type` - Type of allow list.
    * `allow_list` - Ip list of allow list.
    * `associated_instance_num` - The number of instance that associated to allow list.
    * `associated_instances` - Instances associated by this allow list.
        * `instance_id` - Id of instance.
        * `instance_name` - Name of instance.
        * `vpc` - Id of virtual private cloud.
* `total_count` - The total count of allow list query.


