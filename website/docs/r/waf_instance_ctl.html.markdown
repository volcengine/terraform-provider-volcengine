---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_instance_ctl"
sidebar_current: "docs-volcengine-resource-waf_instance_ctl"
description: |-
  Provides a resource to manage waf instance ctl
---
# volcengine_waf_instance_ctl
Provides a resource to manage waf instance ctl
## Example Usage
```hcl
resource "volcengine_waf_instance_ctl" "foo" {
  allow_enable = 0
  block_enable = 1
  project_name = "default"
}
```
## Argument Reference
The following arguments are supported:
* `allow_enable` - (Optional) Whether to enable the allowed access list policy for the instance corresponding to the current region.
* `block_enable` - (Optional) Whether to enable the prohibited access list policy for the instance corresponding to the current region.
* `project_name` - (Optional) The name of the project associated with the current resource.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
WafInstanceCtl can be imported using the id, e.g.
```
$ terraform import volcengine_waf_instance_ctl.default resource_id
```

