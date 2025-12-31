---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_alerting_rule_enable_disable"
sidebar_current: "docs-volcengine-resource-vmp_alerting_rule_enable_disable"
description: |-
  Provides a resource to manage vmp alerting rule enable disable
---
# volcengine_vmp_alerting_rule_enable_disable
Provides a resource to manage vmp alerting rule enable disable
## Example Usage
```hcl
resource "volcengine_vmp_alerting_rule_enable_disable" "example" {
  ids = [
    "b9b6407d-f602-4f2e-b2e8-3b21286b7efa",
    "1cb9a731-d182-4ccc-b374-d4a06ae84714"
  ]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Required, ForceNew) The ids of alerting rule.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
The VmpAlertingRuleEnableDisable is not support import.

