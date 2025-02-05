---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_allow_list"
sidebar_current: "docs-volcengine-resource-mongodb_allow_list"
description: |-
  Provides a resource to manage mongodb allow list
---
# volcengine_mongodb_allow_list
Provides a resource to manage mongodb allow list
## Example Usage
```hcl
resource "volcengine_mongodb_allow_list" "foo" {
  allow_list_name = "acc-test-allow-list"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = "10.1.1.3,10.2.3.0/24,10.1.1.1"
  project_name    = "default"
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_name` - (Required) The name of allow list.
* `allow_list` - (Required) IP address or IP address segment in CIDR format. Duplicate addresses are not allowed, multiple addresses should be separated by commas (,) in English.
* `allow_list_desc` - (Optional) The description of allow list.
* `allow_list_type` - (Optional) The IP address type of allow list, valid value contains `IPv4`.
* `project_name` - (Optional) The project name of the allow list.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `allow_list_ip_num` - The number of allow list IPs.
* `associated_instance_num` - The total number of instances bound under the allow list.
* `associated_instances` - The list of associated instances.
    * `instance_id` - The instance id that bound to the allow list.
    * `instance_name` - The instance name that bound to the allow list.
    * `project_name` - The project name of the instance.
    * `vpc` - The VPC ID.


## Import
mongodb allow list can be imported using the allowListId, e.g.
```
$ terraform import volcengine_mongodb_allow_list.default acl-d1fd76693bd54e658912e7337d5b****
```

