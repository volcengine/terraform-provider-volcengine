---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_rule_bound_host_group"
sidebar_current: "docs-volcengine-resource-tls_rule_bound_host_group"
description: |-
  Provides a resource to manage tls rule bound host group
---
# volcengine_tls_rule_bound_host_group
Provides a resource to manage tls rule bound host group
## Example Usage
```hcl
resource "volcengine_tls_rule_bound_host_group" "foo" {
  rule_id       = "048dc010-6bb1-4189-858a-281d654d6686"
  host_group_id = "e45643c1-8ab8-4d99-94c6-ddcc7eefbd2b"
}
```
## Argument Reference
The following arguments are supported:
* `host_group_id` - (Required, ForceNew) The ID of the host group.
* `rule_id` - (Required, ForceNew) The ID of the rule.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TlsRuleBoundHostGroup can be imported using the id, e.g.
```
$ terraform import volcengine_tls_rule_bound_host_group.default rule_id:host_group_id
```

