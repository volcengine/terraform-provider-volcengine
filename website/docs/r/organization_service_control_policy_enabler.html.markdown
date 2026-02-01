---
subcategory: "ORGANIZATION"
layout: "volcengine"
page_title: "Volcengine: volcengine_organization_service_control_policy_enabler"
sidebar_current: "docs-volcengine-resource-organization_service_control_policy_enabler"
description: |-
  Provides a resource to manage organization service control policy enabler
---
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

