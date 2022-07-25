---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_clb_rule"
sidebar_current: "docs-volcengine-resource-clb_rule"
description: |-
  Provides a resource to manage clb rule
---
# volcengine_clb_rule
Provides a resource to manage clb rule
## Example Usage
```hcl
resource "volcengine_clb_rule" "foo" {
  listener_id     = "lsn-273ywvnmiu70g7fap8u2xzg9d"
  server_group_id = "rsp-273yxuqfova4g7fap8tyemn6t"
  domain          = "test-volc123.com"
  url             = "/yyyy"
}
```
## Argument Reference
The following arguments are supported:
* `listener_id` - (Required, ForceNew) The ID of listener.
* `server_group_id` - (Required) Server Group Id.
* `description` - (Optional) The description of the Rule.
* `domain` - (Optional, ForceNew) The domain of Rule.
* `url` - (Optional, ForceNew) The Url of Rule.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Rule can be imported using the id, e.g.
Notice: resourceId is ruleId, due to the lack of describeRuleAttributes in openapi, for import resources, please use ruleId:listenerId to import.
we will fix this problem later.
```
$ terraform import volcengine_clb_rule.foo rule-273zb9hzi1gqo7fap8u1k3utb:lsn-273ywvnmiu70g7fap8u2xzg9d
```

