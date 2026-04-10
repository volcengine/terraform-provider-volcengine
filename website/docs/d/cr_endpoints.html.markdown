---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_endpoints"
sidebar_current: "docs-volcengine-datasource-cr_endpoints"
description: |-
  Use this data source to query detailed information of cr endpoints
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_cr_endpoints
Use this data source to query detailed information of cr endpoints
## Example Usage
```hcl
data "volcengine_cr_endpoints" "foo" {
  registry = "tf-1"
}
```
## Argument Reference
The following arguments are supported:
* `registry` - (Required) The CR instance name.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `endpoints` - The collection of endpoint query.
    * `acl_policies` - The list of acl policies.
        * `description` - The description of the acl policy.
        * `entry` - The ip of the acl policy.
    * `enabled` - Whether public endpoint is enabled.
    * `registry` - The name of CR instance.
    * `status` - The status of public endpoint.
* `total_count` - The total count of tag query.


