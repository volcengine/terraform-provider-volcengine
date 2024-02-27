---
subcategory: "ORGANIZATION"
layout: "volcengine"
page_title: "Volcengine: volcengine_organization_service_control_policy_attachment"
sidebar_current: "docs-volcengine-resource-organization_service_control_policy_attachment"
description: |-
  Provides a resource to manage organization service control policy attachment
---
# volcengine_organization_service_control_policy_attachment
Provides a resource to manage organization service control policy attachment
## Example Usage
```hcl
resource "volcengine_organization_service_control_policy" "foo" {
  policy_name = "tfpolicy11"
  description = "tftest1"
  statement   = "{\"Statement\":[{\"Effect\":\"Deny\",\"Action\":[\"ecs:RunInstances\"],\"Resource\":[\"*\"]}]}"
}

resource "volcengine_organization_service_control_policy_attachment" "foo" {
  policy_id   = volcengine_organization_service_control_policy.foo.id
  target_id   = "21*********94"
  target_type = "Account"
}

resource "volcengine_organization_service_control_policy_attachment" "foo1" {
  policy_id   = volcengine_organization_service_control_policy.foo.id
  target_id   = "73*********9"
  target_type = "OU"
}
```
## Argument Reference
The following arguments are supported:
* `policy_id` - (Required, ForceNew) The id of policy.
* `target_id` - (Required, ForceNew) The id of target.
* `target_type` - (Required, ForceNew) The type of target. Support Account or OU.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Service Control policy attachment can be imported using the id, e.g.
```
$ terraform import volcengine_organization_service_control_policy_attachment.default PolicyID:TargetID
```

