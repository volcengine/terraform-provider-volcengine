---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_allow_lists"
sidebar_current: "docs-volcengine-datasource-mongodb_allow_lists"
description: |-
  Use this data source to query detailed information of mongodb allow lists
---
# volcengine_mongodb_allow_lists
Use this data source to query detailed information of mongodb allow lists
## Example Usage
```hcl
data "volcengine_mongodb_allow_lists" "default" {
  region_id      = "cn-xxx"
  instance_id    = "mongo-replica-xxx"
  allow_list_ids = ["acl-2ecfc3318fd24bfab6xxx", "acl-ada659ab83e941d6adc2xxxf"]
}
```
## Argument Reference
The following arguments are supported:
* `region_id` - (Required) The region ID.
* `allow_list_ids` - (Optional) The allow list IDs to query.
* `instance_id` - (Optional) The instance ID to query.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `allow_lists` - The collection of mongodb allow list query.
    * `allow_list_desc` - The description of allow list.
    * `allow_list_id` - The ID of allow list.
    * `allow_list_ip_num` - The number of allow list IPs.
    * `allow_list_name` - The allow list name.
    * `allow_list_type` - The IP address type in allow list.
    * `allow_list` - The list of IP address in allow list.
    * `associated_instance_num` - The total number of instances bound under the allow list.
    * `associated_instances` - The list of associated instances.
        * `instance_id` - The instance id that bound to the allow list.
        * `instance_name` - The instance name that bound to the allow list.
        * `vpc` - The VPC ID.


