---
subcategory: "IAM"
layout: "volcengine"
page_title: "Volcengine: volcengine_iam_allowed_ip_addresses"
sidebar_current: "docs-volcengine-datasource-iam_allowed_ip_addresses"
description: |-
  Use this data source to query detailed information of iam allowed ip addresses
---
# volcengine_iam_allowed_ip_addresses
Use this data source to query detailed information of iam allowed ip addresses
## Example Usage
```hcl
data "volcengine_iam_allowed_ip_addresses" "default" {

}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `allowed_ip_addresses` - The collection of query.
    * `enable_ip_list` - Whether to enable the IP whitelist.
    * `ip_list` - The IP whitelist list.
        * `description` - The description of the IP address.
        * `ip` - The IP address.
    * `quota` - The quota of the IP whitelist.


