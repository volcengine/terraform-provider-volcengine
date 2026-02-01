---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_silence_policy_enable_disable"
sidebar_current: "docs-volcengine-resource-vmp_silence_policy_enable_disable"
description: |-
  Provides a resource to manage vmp silence policy enable disable
---
# volcengine_vmp_silence_policy_enable_disable
Provides a resource to manage vmp silence policy enable disable
## Example Usage
```hcl
resource "volcengine_vmp_silence_policy_enable_disable" "foo" {
  ids = ["4d62547e-a0f4-4bdd-a658-399fc4464ae8", "ea51e747-0ead-4e09-9187-76beba6400b7"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Required, ForceNew) The ids of silence policy.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
The VmpSilencePolicyEnableDisable is not support import.

