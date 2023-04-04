---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_vpc_endpoints"
sidebar_current: "docs-volcengine-datasource-cr_vpc_endpoints"
description: |-
  Use this data source to query detailed information of cr vpc endpoints
---
# volcengine_cr_vpc_endpoints
Use this data source to query detailed information of cr vpc endpoints
## Example Usage
```hcl
data "volcengine_cr_vpc_endpoints" "default" {
  registry = "enterprise-1"
  statuses = ["Enabled", "Enabling", "Disabling", "Failed"]
}
```
## Argument Reference
The following arguments are supported:
* `registry` - (Required) The CR registry name.
* `output_file` - (Optional) File name where to save data source results.
* `statuses` - (Optional) VPC access entry state array, used to filter out VPC access entries in the specified state. Available values are Enabling, Enabled, Disabling, Failed.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `endpoints` - List of CR vpc endpoints.
    * `registry` - The name of CR registry.
    * `vpcs` - List of vpc information.
        * `account_id` - The id of the account.
        * `create_time` - The creation time.
        * `ip` - The IP address of the mirror repository in the VPC.
        * `region` - The region id.
        * `status` - The status of the vpc endpoint.
        * `subnet_id` - The ID of the subnet.
        * `vpc_id` - The ID of the vpc.
* `total_count` - The total count of CR vpc endpoints query.


