---
subcategory: "CLOUD_FIREWALL"
layout: "volcengine"
page_title: "Volcengine: volcengine_cfw_dns_control_policy"
sidebar_current: "docs-volcengine-resource-cfw_dns_control_policy"
description: |-
  Provides a resource to manage cfw dns control policy
---
# volcengine_cfw_dns_control_policy
Provides a resource to manage cfw dns control policy
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_cfw_dns_control_policy" "foo" {
  description      = "acc-test-dns-control-policy"
  destination_type = "domain"
  destination      = "www.test.com"
  source {
    vpc_id = volcengine_vpc.foo.id
    region = "cn-beijing"
  }
}
```
## Argument Reference
The following arguments are supported:
* `destination_type` - (Required) The destination type of the dns control policy. Valid values: `group`, `domain`.
* `destination` - (Required) The destination of the dns control policy.
* `source` - (Required) The source vpc list of the dns control policy.
* `description` - (Optional) The description of the dns control policy.
* `internet_firewall_id` - (Optional) The internet firewall id of the control policy.
* `status` - (Optional) Whether to enable the dns control policy.

The `source` object supports the following:

* `region` - (Required) The region of the source vpc.
* `vpc_id` - (Required) The id of the source vpc.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account id of the dns control policy.
* `hit_cnt` - The hit count of the dns control policy.
* `last_hit_time` - The last hit time of the dns control policy. Unix timestamp.
* `use_count` - The use count of the dns control policy.


## Import
DnsControlPolicy can be imported using the id, e.g.
```
$ terraform import volcengine_dns_control_policy.default resource_id
```

