---
subcategory: "ORGANIZATION"
layout: "volcengine"
page_title: "Volcengine: volcengine_organization_service_control_policy_enabler"
sidebar_current: "docs-volcengine-resource-organization_service_control_policy_enabler"
description: |-
  Provides a resource to manage organization service control policy enabler
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_organization_service_control_policy_enabler
Provides a resource to manage organization service control policy enabler
## Example Usage
```hcl
resource "volcengine_organization_service_control_policy_enabler" "foo" {

}
```
## Argument Reference
The following arguments are supported:


## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ServiceControlPolicy enabler can be imported using the default_id (organization:service_control_policy_enable) , e.g.
```
$ terraform import volcengine_organization_service_control_policy_enabler.default organization:service_control_policy_enable
```

