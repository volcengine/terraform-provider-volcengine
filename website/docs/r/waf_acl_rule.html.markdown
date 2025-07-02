---
subcategory: "WAF"
layout: "volcengine"
page_title: "Volcengine: volcengine_waf_acl_rule"
sidebar_current: "docs-volcengine-resource-waf_acl_rule"
description: |-
  Provides a resource to manage waf acl rule
---
# volcengine_waf_acl_rule
Provides a resource to manage waf acl rule
## Example Usage
```hcl
resource "volcengine_waf_acl_rule" "foo" {
  action        = "block"
  description   = "tf-test"
  name          = "tf-test-1"
  url           = "/"
  ip_add_type   = 3
  host_add_type = 3
  enable        = 1
  acl_type      = "Allow"
  host_list     = ["www.tf-test.com"]
  ip_list       = ["1.2.2.2", "1.2.3.30"]
  accurate_group {
    accurate_rules {
      http_obj     = "request.uri"
      obj_type     = 1
      opretar      = 2
      property     = 0
      value_string = "GET"
    }
    logic = 1
  }
  advanced     = 1
  project_name = "default"
}
```
## Argument Reference
The following arguments are supported:
* `acl_type` - (Required, ForceNew) The type of access control rules.
* `enable` - (Required, ForceNew) Whether to enable the rule.
* `host_add_type` - (Required, ForceNew) Type of domain name addition.
* `ip_add_type` - (Required, ForceNew) Type of IP address addition.
* `name` - (Required, ForceNew) Rule name.
* `url` - (Required, ForceNew) The path of Matching.
* `accurate_group` - (Optional, ForceNew) Advanced conditions.
* `action` - (Optional, ForceNew) Action to be taken on requests that match the rule.
* `advanced` - (Optional, ForceNew) Whether to set advanced conditions.
* `description` - (Optional, ForceNew) Rule description.
* `host_group_id` - (Optional, ForceNew) The ID of the domain group.
* `host_list` - (Optional, ForceNew) Required if HostAddType = 3. Single or multiple domain names are supported.
* `ip_group_id` - (Optional, ForceNew) Required if IpAddType = 2.
* `ip_list` - (Optional, ForceNew) Required if IpAddType = 3. Single or multiple IP addresses are supported.
* `ip_location_country` - (Optional, ForceNew) Country or region code.
* `ip_location_subregion` - (Optional, ForceNew) Domestic region code.
* `project_name` - (Optional, ForceNew) The name of the project to which your domain names belong.

The `accurate_group` object supports the following:

* `accurate_rules` - (Required, ForceNew) Details of advanced conditions.
* `logic` - (Required, ForceNew) The logical relationship of advanced conditions.

The `accurate_rules` object supports the following:

* `http_obj` - (Required, ForceNew) The HTTP object to be added to the advanced conditions.
* `obj_type` - (Required, ForceNew) The matching field for HTTP objects.
* `opretar` - (Required, ForceNew) The logical operator for the condition.
* `property` - (Required, ForceNew) Operate the properties of the http object.
* `value_string` - (Required, ForceNew) The value to be matched.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `client_ip` - IP address.
* `host_groups` - The list of domain name groups.
    * `host_group_id` - The ID of host group.
    * `name` - The Name of host group.
* `ip_groups` - The list of domain name groups.
    * `ip_group_id` - The ID of the IP address group.
    * `name` - The Name of the IP address group.
* `rule_tag` - Rule unique identifier.
* `update_time` - Update time of the rule.


## Import
WafAclRule can be imported using the id, e.g.
```
$ terraform import volcengine_waf_acl_rule.default resource_id:AclType
```

