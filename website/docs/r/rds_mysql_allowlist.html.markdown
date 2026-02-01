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
  allow_list_name = "acc-test-allowlist"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  user_allow_list = ["192.168.0.0/24", "192.168.1.0/24"]
  //user_allow_list = ["192.168.0.0/24", "192.168.1.0/24"]
  security_group_bind_infos {
    bind_mode         = "IngressDirectionIp"
    security_group_id = "sg-13fd7wyduxekg3n6nu5t9fhj7"
  }
  security_group_bind_infos {
    bind_mode         = "IngressDirectionIp"
    security_group_id = "sg-mjoa9qfyzg1s5smt1a6dmc1l"
  }
  #security_group_ids = ["sg-13fd7wyduxekg3n6nu5t9fhj7", "sg-mjoa9qfyzg1s5smt1a6dmc1l", "sg-mirtbey0outc5smt1bom7lwz"]
}
```
## Argument Reference
The following arguments are supported:
* `allow_list_name` - (Required) The name of the allow list.
* `allow_list_category` - (Optional) White list category. Values:
Ordinary: Ordinary white list.
Default: Default white list.
 Description: When this parameter is used as a request parameter, the default value is Ordinary.
* `allow_list_desc` - (Optional) The description of the allow list.
* `allow_list_type` - (Optional) The type of IP address in the whitelist. Currently only IPv4 addresses are supported.
* `allow_list` - (Optional) Enter an IP address or a range of IP addresses in CIDR format. Please note that if you want to use security group - related parameters, do not use this field. Instead, use the user_allow_list.
* `security_group_bind_infos` - (Optional) Whitelist information for the associated security group.
* `security_group_ids` - (Optional) The security group ids of the allow list.
* `user_allow_list` - (Optional) IP addresses outside the security group that need to be added to the whitelist. IP addresses or IP address segments in CIDR format can be entered. Note: This field cannot be used simultaneously with AllowList.

The `security_group_bind_infos` object supports the following:

* `bind_mode` - (Required) The schema for the associated security group.
 IngressDirectionIp: Incoming Direction IP. 
 AssociateEcsIp: Associate ECSIP. 
explain: In the CreateAllowList interface, SecurityGroupBindInfoObject BindMode and SecurityGroupId fields are required.
* `security_group_id` - (Required) The security group id of the allow list.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `allow_list_id` - The id of the allow list.


## Import
RDS AllowList can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_allowlist.default acl-d1fd76693bd54e658912e7337d5b****
```

