---
subcategory: "CDN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cdn_shared_configs"
sidebar_current: "docs-volcengine-datasource-cdn_shared_configs"
description: |-
  Use this data source to query detailed information of cdn shared configs
---
# volcengine_cdn_shared_configs
Use this data source to query detailed information of cdn shared configs
## Example Usage
```hcl
data "volcengine_cdn_shared_configs" "foo" {
  config_name  = "tf-test"
  config_type  = "allow_ip_access_rule"
  project_name = "default"
}
```
## Argument Reference
The following arguments are supported:
* `config_name` - (Optional) The name of the shared config.
* `config_type_list` - (Optional) The config type list. The parameter value can be a combination of available values for ConfigType. ConfigType and ConfigTypeList cannot be specified at the same time.
* `config_type` - (Optional) The type of the shared config.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The name of the project.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `config_data` - The collection of query.
    * `allow_ip_access_rule` - The configuration for IP whitelist corresponds to ConfigType allow_ip_access_rule.
        * `rules` - The entries in this list are an array of IP addresses and CIDR network segments. The total number of entries cannot exceed 3,000. The IP addresses and segments can be in IPv4 and IPv6 format. Duplicate entries in the list will be removed and will not count towards the limit.
    * `allow_referer_access_rule` - The configuration for the Referer whitelist corresponds to ConfigType allow_referer_access_rule.
        * `allow_empty` - Indicates whether an empty Referer header, or a request without a Referer header, is not allowed. Default is false.
        * `common_type` - The content indicating the Referer whitelist.
            * `ignore_case` - This list is case-sensitive when matching requests. Default is true.
            * `rules` - The entries in this list are an array of IP addresses and CIDR network segments. The total number of entries cannot exceed 3,000. The IP addresses and segments can be in IPv4 and IPv6 format. Duplicate entries in the list will be removed and will not count towards the limit.
    * `common_match_list` - The configuration for a common list is represented by ConfigType common_match_list.
        * `common_type` - The content indicating the Referer blacklist.
            * `ignore_case` - This list is case-sensitive when matching requests. Default is true.
            * `rules` - The entries in this list are an array of IP addresses and CIDR network segments. The total number of entries cannot exceed 3,000. The IP addresses and segments can be in IPv4 and IPv6 format. Duplicate entries in the list will be removed and will not count towards the limit.
    * `config_name` - The name of the config.
    * `config_type` - The type of the config.
    * `deny_ip_access_rule` - The configuration for IP blacklist is denoted by ConfigType deny_ip_access_rule.
        * `rules` - The entries in this list are an array of IP addresses and CIDR network segments. The total number of entries cannot exceed 3,000. The IP addresses and segments can be in IPv4 and IPv6 format. Duplicate entries in the list will be removed and will not count towards the limit.
    * `deny_referer_access_rule` - The configuration for the Referer blacklist corresponds to ConfigType deny_referer_access_rule.
        * `allow_empty` - Indicates whether an empty Referer header, or a request without a Referer header, is not allowed. Default is false.
        * `common_type` - The content indicating the Referer blacklist.
            * `ignore_case` - This list is case-sensitive when matching requests. Default is true.
            * `rules` - The entries in this list are an array of IP addresses and CIDR network segments. The total number of entries cannot exceed 3,000. The IP addresses and segments can be in IPv4 and IPv6 format. Duplicate entries in the list will be removed and will not count towards the limit.
    * `domain_count` - The number of domains.
    * `project_name` - The name of the project.
    * `update_time` - The update time of the shared config.
* `total_count` - The total count of query.


