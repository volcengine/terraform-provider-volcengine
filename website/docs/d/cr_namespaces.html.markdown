---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_namespaces"
sidebar_current: "docs-volcengine-datasource-cr_namespaces"
description: |-
  Use this data source to query detailed information of cr namespaces
---
# volcengine_cr_namespaces
Use this data source to query detailed information of cr namespaces
## Example Usage
```hcl
data "volcengine_cr_namespaces" "foo" {
  registry = "tf-1"
  names    = ["namespace-*"]
}
```
## Argument Reference
The following arguments are supported:
* `registry` - (Required) The target cr instance name.
* `names` - (Optional) The list of instance IDs.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `namespaces` - The collection of namespaces query.
    * `create_time` - The time when namespace created.
    * `name` - The name of OCI repository.
* `total_count` - The total count of instance query.


