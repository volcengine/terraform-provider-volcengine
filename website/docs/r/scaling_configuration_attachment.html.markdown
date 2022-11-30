---
subcategory: "AUTOSCALING"
layout: "volcengine"
page_title: "Volcengine: volcengine_scaling_configuration_attachment"
sidebar_current: "docs-volcengine-resource-scaling_configuration_attachment"
description: |-
  Provides a resource to manage scaling configuration attachment
---
# volcengine_scaling_configuration_attachment
Provides a resource to manage scaling configuration attachment
## Example Usage
```hcl
resource "volcengine_scaling_configuration_attachment" "foo1" {
  scaling_configuration_id = "scc-ybrurj4uw6gh9zecj327"
}
```
## Argument Reference
The following arguments are supported:
* `scaling_configuration_id` - (Required, ForceNew) The id of the scaling configuration.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Scaling Configuration attachment can be imported using the scaling_configuration_id e.g.
```
$ terraform import volcengine_scaling_configuration_attachment.default enable:scc-ybrurj4uw6gh9zecj327
```

