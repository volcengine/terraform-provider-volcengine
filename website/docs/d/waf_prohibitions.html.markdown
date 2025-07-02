---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_prohibitions"
sidebar_current: "docs-volcengine-datasource-waf_prohibitions"
description: |-
  Use this data source to query detailed information of waf prohibitions
---
# volcengine_waf_prohibitions
Use this data source to query detailed information of waf prohibitions
## Example Usage
```hcl
data "volcengine_waf_prohibitions" "foo" {
  start_time = 1749805224
  end_time   = 1749808824
  host       = "www.tf-test.com"
}
```
## Argument Reference
The following arguments are supported:
* `end_time` - (Required) end time.
* `host` - (Required) The domain name of the website that needs to be queried.
* `start_time` - (Required) starting time.
* `letter_order_by` - (Optional) The list shows the order.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `reason` - (Optional) Attack type filtering.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `ip_agg_group` - Details of the attack IP.
    * `drop_count` - The number of attacks on the source IP of this attack.
    * `ip` - Attack source IP.
    * `reason` - Reason for the ban.
        * `black` - The number of visits to the blacklist.
        * `bot` - The number of Bot attacks.
        * `geo_black` - The number of geographical location access control.
        * `http_flood` - The number of CC attacks.
        * `param_abnormal` - The number of API parameter exceptions.
        * `route_abnormal` - The number of API routing exceptions.
        * `sensitive_info` - The number of times sensitive information is leaked.
        * `web_vulnerability` - The number of Web vulnerability attacks.
    * `rule_name` - Name of the ban rule.
    * `rule_tag` - Ban rule ID.
    * `status` - IP banned status.
    * `update_time` - Status update time.
* `total_count` - The total count of query.


