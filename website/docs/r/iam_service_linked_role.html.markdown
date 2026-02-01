---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_service_linked_role"
sidebar_current: "docs-volcengine-resource-iam_service_linked_role"
description: |-
  Provides a resource to manage iam service linked role
---
# volcengine_iam_service_linked_role
Provides a resource to manage iam service linked role
## Example Usage
```hcl
resource "volcengine_iam_service_linked_role" "foo" {
  service_name = "ecs"
}
```
## Argument Reference
The following arguments are supported:
* `service_name` - (Required, ForceNew) The name of the service.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `role_id` - The id of the role.
* `role_name` - The name of the role.
* `status` - The status of the role.


## Import
IamServiceLinkedRole can be imported using the id, e.g.
```
$ terraform import volcengine_iam_service_linked_role.default resource_id
```

