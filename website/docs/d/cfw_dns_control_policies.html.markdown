---
subcategory: "CLOUD_FIREWALL"
layout: "volcengine"
page_title: "Volcengine: volcengine_cfw_dns_control_policies"
sidebar_current: "docs-volcengine-datasource-cfw_dns_control_policies"
description: |-
  Use this data source to query detailed information of cfw dns control policies
---
# volcengine_cfw_dns_control_policies
Use this data source to query detailed information of cfw dns control policies
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

data "volcengine_cfw_dns_control_policies" "foo" {
  ids = [volcengine_cfw_dns_control_policy.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `description` - (Optional) The description of the dns control policy. This field support fuzzy query.
* `destination` - (Optional) The destination list of the dns control policy. This field support fuzzy query.
* `ids` - (Optional) The rule id list of the dns control policy. This field support fuzzy query.
* `internet_firewall_id` - (Optional) The internet firewall id of the dns control policy.
* `output_file` - (Optional) File name where to save data source results.
* `source` - (Optional) The source list of the dns control policy. This field support fuzzy query.
* `status` - (Optional) The enable status list of the dns control policy. This field support fuzzy query.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `dns_control_policies` - The collection of query.
    * `account_id` - The account id of the dns control policy.
    * `description` - The description of the dns control policy.
    * `destination_group_list` - The destination group list of the dns control policy.
    * `destination_type` - The destination type of the dns control policy.
    * `destination` - The destination of the dns control policy.
    * `domain_list` - The destination domain list of the dns control policy.
    * `hit_cnt` - The hit count of the dns control policy.
    * `id` - The id of the dns control policy.
    * `last_hit_time` - The last hit time of the dns control policy. Unix timestamp.
    * `rule_id` - The id of the dns control policy.
    * `source` - The source vpc list of the dns control policy.
        * `region` - The region of the source vpc.
        * `vpc_id` - The id of the source vpc.
    * `status` - Whether to enable the dns control policy.
    * `use_count` - The use count of the dns control policy.
* `total_count` - The total count of query.


