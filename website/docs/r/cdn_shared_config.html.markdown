---
subcategory: "CDN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cdn_shared_config"
sidebar_current: "docs-volcengine-resource-cdn_shared_config"
description: |-
  Provides a resource to manage cdn shared config
---
# volcengine_cdn_shared_config
Provides a resource to manage cdn shared config
## Example Usage
```hcl
resource "volcengine_cdn_shared_config" "foo" {
  config_name = "tftest"
  config_type = "allow_referer_access_rule"
  allow_ip_access_rule {
    rules = ["1.1.1.1", "2.2.2.0/24", "3.3.3.3"]
  }
  deny_ip_access_rule {
    rules = ["1.1.1.1", "2.2.2.0/24"]
  }
  common_match_list {
    common_type {
      rules = ["1.1.1.1", "2.2.2.0/24"]
    }
  }
  allow_referer_access_rule {
    common_type {
      rules = ["1.1.1.1", "2.2.2.0/24", "3.3.4.4"]
    }
  }
  deny_referer_access_rule {
    common_type {
      rules = ["1.1.1.1", "2.2.2.0/24"]
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `config_name` - (Required, ForceNew) The name of the shared config. The name cannot be the same as the name of an existing global configuration under the main account.
* `config_type` - (Required, ForceNew) The type of the shared config. The type of global configuration. The parameter can have the following values: `deny_ip_access_rule`: represents IP blacklist. `allow_ip_access_rule`: represents IP whitelist. `deny_referer_access_rule`: represents Referer blacklist. `allow_referer_access_rule`: represents Referer whitelist. `common_match_list`: represents common list.
* `allow_ip_access_rule` - (Optional) The configuration for IP whitelist corresponds to ConfigType allow_ip_access_rule.
* `allow_referer_access_rule` - (Optional) The configuration for the Referer whitelist corresponds to ConfigType allow_referer_access_rule.
* `common_match_list` - (Optional) The configuration for a common list is represented by ConfigType common_match_list.
* `deny_ip_access_rule` - (Optional) The configuration for IP blacklist is denoted by ConfigType deny_ip_access_rule.
* `deny_referer_access_rule` - (Optional) The configuration for the Referer blacklist corresponds to ConfigType deny_referer_access_rule.
* `project_name` - (Optional, ForceNew) The ProjectName of the cdn shared config.

The `allow_ip_access_rule` object supports the following:

* `rules` - (Required) The entries in this list are an array of IP addresses and CIDR network segments. The total number of entries cannot exceed 3,000. The IP addresses and segments can be in IPv4 and IPv6 format. Duplicate entries in the list will be removed and will not count towards the limit.

The `allow_referer_access_rule` object supports the following:

* `common_type` - (Required) The content indicating the Referer whitelist.
* `allow_empty` - (Optional) Indicates whether an empty Referer header, or a request without a Referer header, is not allowed. Default is false.

The `common_match_list` object supports the following:

* `common_type` - (Required) The content indicating the Referer blacklist.

The `common_type` object supports the following:

* `rules` - (Required) The entries in this list are an array of IP addresses and CIDR network segments. The total number of entries cannot exceed 3,000. The IP addresses and segments can be in IPv4 and IPv6 format. Duplicate entries in the list will be removed and will not count towards the limit.
* `ignore_case` - (Optional) This list is case-sensitive when matching requests. Default is true.

The `deny_ip_access_rule` object supports the following:

* `rules` - (Required) The entries in this list are an array of IP addresses and CIDR network segments. The total number of entries cannot exceed 3,000. The IP addresses and segments can be in IPv4 and IPv6 format. Duplicate entries in the list will be removed and will not count towards the limit.

The `deny_referer_access_rule` object supports the following:

* `common_type` - (Required) The content indicating the Referer blacklist.
* `allow_empty` - (Optional) Indicates whether an empty Referer header, or a request without a Referer header, is not allowed. Default is false.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
CdnSharedConfig can be imported using the id, e.g.
```
$ terraform import volcengine_cdn_shared_config.default resource_id
```

