---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_integration_task_enable"
sidebar_current: "docs-volcengine-resource-vmp_integration_task_enable"
description: |-
  Provides a resource to manage vmp integration task enable
---
# volcengine_vmp_integration_task_enable
Provides a resource to manage vmp integration task enable
## Example Usage
```hcl
# Create a VMP integration task enable
resource "volcengine_vmp_integration_task_enable" "default" {
  task_ids = ["3c55cdd4-f240-4fc8-a43b-ca83ad44807a", "a09fdaf5-ce90-4f34-8ab3-4decd5aef8e1"]
}
```
## Argument Reference
The following arguments are supported:
* `task_ids` - (Required, ForceNew) A list of integration task IDs to enable.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
VMP Integration Task Enable can be imported using the task ids, e.g.
```
$ terraform import volcengine_vmp_integration_task_enable.default task-id1,task-id2
```

