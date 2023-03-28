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
  allow_list_name = "tf-test-hh"
  allow_list_desc = "test1"
  allow_list_type = "IPv4"
  allow_list      = "10.1.1.3,10.2.3.0/24,10.1.1.1"
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_name` - (Required) The name of allow list.
* `allow_list` - (Required) IP address or IP address segment in CIDR format.
* `allow_list_desc` - (Optional) The description of allow list.
* `allow_list_type` - (Optional) The IP address type of allow list, valid value contains `IPv4`.
* `modify_mode` - (Optional) The modify mode. Only support Cover mode.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
mongodb allow list can be imported using the allowListId, e.g.
```
$ terraform import volcengine_mongodb_allow_list.default acl-d1fd76693bd54e658912e7337d5b****
```

