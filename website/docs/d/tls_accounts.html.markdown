---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_accounts"
sidebar_current: "docs-volcengine-datasource-tls_accounts"
description: |-
  Use this data source to query detailed information of tls accounts
---
# volcengine_tls_accounts
Use this data source to query detailed information of tls accounts
## Example Usage
```hcl
data "volcengine_tls_accounts" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `tls_accounts` - The collection of tls account query.
    * `arch_version` - The version of the log service architecture. Valid values: 2.0 (new architecture), 1.0 (old architecture).
    * `status` - The status of the log service. Valid values: Activated (already activated), NonActivated (not activated).
* `total_count` - The total count of tls account query.


