---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_rule_applier"
sidebar_current: "docs-volcengine-resource-tls_rule_applier"
description: |-
  Provides a resource to manage tls rule applier
---
# volcengine_tls_rule_applier
Provides a resource to manage tls rule applier
## Example Usage
```hcl
resource "volcengine_tls_rule_applier" "foo" {
  host_group_id = "a2a9e8c5-9835-434f-b866-2c1cfa82887d"
  rule_id       = "25104b0f-28b7-4a5c-8339-7f9c431d77c8"
}
```
## Argument Reference
The following arguments are supported:
* `host_group_id` - (Required, ForceNew) The id of the host group.
* `rule_id` - (Required, ForceNew) The id of the rule.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
tls rule applier can be imported using the rule id and host group id, e.g.
```
$ terraform import volcengine_tls_rule_applier.default fa************:bcb*******
```

