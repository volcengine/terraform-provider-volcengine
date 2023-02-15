---
subcategory: "MONGODB"
layout: "volcengine"
page_title: "Volcengine: volcengine_mongodb_allow_list_associate"
sidebar_current: "docs-volcengine-resource-mongodb_allow_list_associate"
description: |-
  Provides a resource to manage mongodb allow list associate
---
# volcengine_mongodb_allow_list_associate
Provides a resource to manage mongodb allow list associate
## Example Usage
```hcl
resource "volcengine_mongodb_allow_list_associate" "foo" {
  instance_id   = "mongo-replica-b2xxx"
  allow_list_id = "acl-4d66bec945d14fa48xxx"
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_id` - (Required, ForceNew) Id of allow list to associate.
* `instance_id` - (Required, ForceNew) Id of instance to associate.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
mongosdb allow list associate can be imported using the instanceId:allowListId, e.g.
```
$ terraform import volcengine_mongosdb_allow_list_associate.default mongo-replica-e405f8e2****:acl-d1fd76693bd54e658912e7337d5b****
```

