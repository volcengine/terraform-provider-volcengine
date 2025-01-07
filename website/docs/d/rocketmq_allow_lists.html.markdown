---
subcategory: "ROCKETMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rocketmq_allow_lists"
sidebar_current: "docs-volcengine-datasource-rocketmq_allow_lists"
description: |-
  Use this data source to query detailed information of rocketmq allow lists
---
# volcengine_rocketmq_allow_lists
Use this data source to query detailed information of rocketmq allow lists
## Example Usage
```hcl
data "volcengine_rocketmq_allow_lists" "foo" {

}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rocketmq_allow_lists` - The collection of query.
    * `allow_list_desc` - The description of the rocketmq allow list.
    * `allow_list_id` - The id of the rocketmq allow list.
    * `allow_list_ip_num` - The number of ip address in the rocketmq allow list.
    * `allow_list_name` - The name of the rocketmq allow list.
    * `allow_list_type` - The type of the rocketmq allow list.
    * `allow_list` - The IP address or a range of IP addresses in CIDR format of the allow list.
    * `associated_instance_num` - The number of the rocketmq instances associated with the allow list.
    * `associated_instances` - The associated instance information of the allow list.
        * `instance_id` - The id of the rocketmq instance.
        * `instance_name` - The name of the rocketmq instance.
        * `vpc` - The vpc id of the rocketmq instance.
    * `id` - The id of the rocketmq allow list.
* `total_count` - The total count of query.


