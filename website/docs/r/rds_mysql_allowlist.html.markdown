---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_allowlist"
sidebar_current: "docs-volcengine-resource-rds_mysql_allowlist"
description: |-
  Provides a resource to manage rds mysql allowlist
---
# volcengine_rds_mysql_allowlist
Provides a resource to manage rds mysql allowlist
## Example Usage
```hcl
resource "volcengine_rds_mysql_allowlist" "foo" {
  allow_list_name = "tf-test-opt"
  allow_list_desc = "terraform test zzm"
  allow_list = [
    "127.0.0.1"
  ]
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
RDS AllowList can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_allowlist.default acl-d1fd76693bd54e658912e7337d5b****
```

