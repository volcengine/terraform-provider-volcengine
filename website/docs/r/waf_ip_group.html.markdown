---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_ip_group"
sidebar_current: "docs-volcengine-resource-waf_ip_group"
description: |-
  Provides a resource to manage waf ip group
---
# volcengine_waf_ip_group
Provides a resource to manage waf ip group
## Example Usage
```hcl
resource "volcengine_waf_ip_group" "foo" {
  add_type = "List"
  ip_list  = ["1.1.1.1", "1.1.1.2", "1.1.1.3"]
  name     = "tf-test"
}
```
## Argument Reference
The following arguments are supported:
* `add_type` - (Required) The way of addition.
* `ip_list` - (Required) The IP address to be added.
* `name` - (Required) The name of ip group.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `ip_count` - The number of IP addresses within the address group.
* `ip_group_id` - The ID of the ip group.
* `related_rules` - The list of associated rules.
    * `host` - The information of the protected domain names associated with the rules.
    * `rule_name` - The name of the rule.
    * `rule_tag` - The ID of the rule.
    * `rule_type` - The type of the rule.
* `update_time` - ip group update time.


## Import
WafIpGroup can be imported using the id, e.g.
```
$ terraform import volcengine_waf_ip_group.default resource_id
```

