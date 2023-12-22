---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_role_attachment"
sidebar_current: "docs-volcengine-resource-iam_role_attachment"
description: |-
  Provides a resource to manage iam role attachment
---
# volcengine_iam_role_attachment
Provides a resource to manage iam role attachment
## Example Usage
```hcl

```
## Argument Reference
The following arguments are supported:
* `iam_role_name` - (Required, ForceNew) The name of the iam role.
* `instance_id` - (Required, ForceNew) The id of the ecs instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
IamRoleAttachment can be imported using the iam_role_name:instance_id, e.g.
```
$ terraform import volcengine_iam_role_attachment.default role_name:instance_id
```

