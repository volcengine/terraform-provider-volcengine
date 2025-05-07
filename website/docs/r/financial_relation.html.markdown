---
subcategory: "FINANCIAL_RELATION"
layout: "volcengine"
page_title: "Volcengine: volcengine_financial_relation"
sidebar_current: "docs-volcengine-resource-financial_relation"
description: |-
  Provides a resource to manage financial relation
---
# volcengine_financial_relation
Provides a resource to manage financial relation
## Example Usage
```hcl
resource "volcengine_financial_relation" "foo" {
  sub_account_id = 2100260000
  relation       = 4
  account_alias  = "acc-test-financial"
  auth_list      = [1, 2, 3]
}
```
## Argument Reference
The following arguments are supported:
* `sub_account_id` - (Required, ForceNew) The sub account id.
* `account_alias` - (Optional, ForceNew) The display name of the sub account.
* `auth_list` - (Optional) The authorization list of financial management. This field is valid and required when the relation is 4. Valid value range is `1-5`.
* `relation` - (Optional, ForceNew) The relation of the financial. Valid values: `1`, `4`. `1` means financial custody, `4` means financial management.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `relation_id` - The id of the financial relation.
* `status` - The status of the financial relation.


## Import
FinancialRelation can be imported using the sub_account_id:relation:relation_id, e.g.
```
$ terraform import volcengine_financial_relation.default resource_id
```

