---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_host_group"
sidebar_current: "docs-volcengine-resource-waf_host_group"
description: |-
  Provides a resource to manage waf host group
---
# volcengine_waf_host_group
Provides a resource to manage waf host group
## Example Usage
```hcl
resource "volcengine_waf_host_group" "foo" {
  description = "tf-test"
  host_list   = ["www.tf-test.com"]
  name        = "tf-test"
}
```
## Argument Reference
The following arguments are supported:
* `host_list` - (Required) Domain names that need to be added to this domain name group.
* `name` - (Required) The name of the domain name group.
* `action` - (Optional) Domain name list modification action. Works only on modified scenes.
* `description` - (Optional) Domain name group description.
* `project_name` - (Optional) The project of Domain name group.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `host_count` - The number of domain names contained in the domain name group.
* `host_group_id` - The ID of the domain name group.
* `related_rules` - The list of associated rules.
    * `rule_name` - The name of the rule.
    * `rule_tag` - The ID of the rule.
    * `rule_type` - The type of the rule.
* `update_time` - Domain name group update time.


## Import
WafHostGroup can be imported using the id, e.g.
```
$ terraform import volcengine_waf_host_group.default resource_id
```

