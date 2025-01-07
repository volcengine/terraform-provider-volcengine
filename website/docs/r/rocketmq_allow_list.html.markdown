---
subcategory: "ROCKETMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rocketmq_allow_list"
sidebar_current: "docs-volcengine-resource-rocketmq_allow_list"
description: |-
  Provides a resource to manage rocketmq allow list
---
# volcengine_rocketmq_allow_list
Provides a resource to manage rocketmq allow list
## Example Usage
```hcl
resource "volcengine_rocketmq_allow_list" "foo" {
  allow_list_name = "acc-test-allow-list"
  allow_list_desc = "acc-test"
  allow_list      = ["192.168.0.0/24", "192.168.2.0/24"]
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_name` - (Required) The name of the allow list.
* `allow_list` - (Required) The list of ip addresses. Enter an IP address or a range of IP addresses in CIDR format.
* `allow_list_desc` - (Optional) The description of the allow list.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `allow_list_ip_num` - The number of ip address in the rocketmq allow list.
* `allow_list_type` - The type of the rocketmq allow list.
* `associated_instance_num` - The number of the rocketmq instances associated with the allow list.
* `associated_instances` - The associated instance information of the allow list.
    * `instance_id` - The id of the rocketmq instance.
    * `instance_name` - The name of the rocketmq instance.
    * `vpc` - The vpc id of the rocketmq instance.


## Import
RocketmqAllowList can be imported using the id, e.g.
```
$ terraform import volcengine_rocketmq_allow_list.default resource_id
```

