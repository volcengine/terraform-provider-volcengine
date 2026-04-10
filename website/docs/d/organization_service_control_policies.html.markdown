---
subcategory: "ORGANIZATION"
layout: "volcengine"
page_title: "Volcengine: volcengine_organization_service_control_policies"
sidebar_current: "docs-volcengine-datasource-organization_service_control_policies"
description: |-
  Use this data source to query detailed information of organization service control policies
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_organization_service_control_policies
Use this data source to query detailed information of organization service control policies
## Example Usage
```hcl
data "volcengine_organization_service_control_policies" "foo" {
  policy_type = "Custom"
  query       = "test"
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `policy_type` - (Optional) The type of policy. The value can be System or Custom.
* `query` - (Optional) Query policies, support policy name or description.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `policies` - The collection of Policy query.
    * `create_date` - The create time of the Policy.
    * `description` - The description of the Policy.
    * `id` - The ID of the Policy.
    * `policy_name` - The name of the Policy.
    * `policy_type` - The type of the Policy.
    * `statement` - The statement of the Policy.
    * `update_date` - The update time of the Policy.
* `total_count` - The total count of Policy query.


