---
subcategory: "PRIVATE_ZONE"
layout: "volcengine"
page_title: "Volcengine: volcengine_private_zone_resolver_rules"
sidebar_current: "docs-volcengine-datasource-private_zone_resolver_rules"
description: |-
  Use this data source to query detailed information of private zone resolver rules
---
# volcengine_private_zone_resolver_rules
Use this data source to query detailed information of private zone resolver rules
## Example Usage
```hcl
data "volcengine_private_zone_resolver_rules" "foo" {}
```
## Argument Reference
The following arguments are supported:
* `endpoint_id` - (Optional) ID of the exit terminal node.
* `name_regex` - (Optional) A Name Regex of Resource.
* `name` - (Optional) The name of the rule.
* `output_file` - (Optional) File name where to save data source results.
* `zone_name` - (Optional) The main domain associated with the forwarding rule. For example, if you set this parameter to example.com, DNS requests for example.com and all subdomains of example.com will be forwarded.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `rules` - The collection of query.
    * `created_at` - The created time of the rule.
    * `endpoint_id` - The endpoint ID of the rule.
    * `forward_ips` - The IP address and port of the DNS server outside of the VPC.
        * `ip` - The IP address of the DNS server outside of the VPC.
        * `port` - The port of the DNS server outside of the VPC.
    * `id` - The id of the rule.
    * `line` - The ISP of the exit IP address of the recursive DNS server.
    * `name` - The name of the rule.
    * `rule_id` - The id of the rule.
    * `type` - The type of the rule.
    * `updated_at` - The updated time of the rule.
    * `zone_name` - The zone name of the rule.
* `total_count` - The total count of query.


