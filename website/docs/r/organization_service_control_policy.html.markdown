---
subcategory: "ORGANIZATION"
layout: "volcengine"
page_title: "Volcengine: volcengine_organization_service_control_policy"
sidebar_current: "docs-volcengine-resource-organization_service_control_policy"
description: |-
  Provides a resource to manage organization service control policy
---
# volcengine_organization_service_control_policy
Provides a resource to manage organization service control policy
## Example Usage
```hcl
resource "volcengine_organization_service_control_policy" "foo" {
  policy_name = "tfpolicy11"
  description = "tftest1"
  statement   = "{\"Statement\":[{\"Effect\":\"Deny\",\"Action\":[\"ecs:RunInstances\"],\"Resource\":[\"*\"]}]}"
}

resource "volcengine_organization_service_control_policy" "foo2" {
  policy_name = "tfpolicy21"
  statement   = "{\"Statement\":[{\"Effect\":\"Deny\",\"Action\":[\"ecs:DeleteInstance\"],\"Resource\":[\"*\"]}]}"
}
```
## Argument Reference
The following arguments are supported:
* `policy_name` - (Required) The name of the Policy.
* `statement` - (Required) The statement of the Policy.
* `description` - (Optional) The description of the Policy.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_date` - The create time of the Policy.
* `policy_type` - The type of the Policy.
* `update_date` - The update time of the Policy.


## Import
Service Control Policy can be imported using the id, e.g.
```
$ terraform import volcengine_organization_service_control_policy.default 1000001
```

