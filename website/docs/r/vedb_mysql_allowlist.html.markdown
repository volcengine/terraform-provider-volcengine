---
subcategory: "VEDB_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_vedb_mysql_allowlist"
sidebar_current: "docs-volcengine-resource-vedb_mysql_allowlist"
description: |-
  Provides a resource to manage vedb mysql allowlist
---
# volcengine_vedb_mysql_allowlist
Provides a resource to manage vedb mysql allowlist
## Example Usage
```hcl
resource "volcengine_vedb_mysql_allowlist" "foo" {
  allow_list_name = "acc-test-allowlist"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = ["192.168.0.0/24", "192.168.1.0/24", "192.168.2.0/24"]
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_name` - (Required) The name of the allow list.
* `allow_list` - (Required) Enter an IP address or a range of IP addresses in CIDR format.
* `allow_list_desc` - (Optional) The description of the allow list.
* `allow_list_type` - (Optional) The type of IP address in the whitelist. Currently only IPv4 addresses are supported.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `allow_list_id` - The id of the allow list.


## Import
VedbMysqlAllowlist can be imported using the id, e.g.
```
$ terraform import volcengine_vedb_mysql_allowlist.default resource_id
```

