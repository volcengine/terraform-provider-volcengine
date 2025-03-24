---
subcategory: "RDS_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_rds_mysql_allowlists"
sidebar_current: "docs-volcengine-datasource-rds_mysql_allowlists"
description: |-
  Use this data source to query detailed information of rds mysql allowlists
---
# volcengine_rds_mysql_allowlists
Use this data source to query detailed information of rds mysql allowlists
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_rds_mysql_allowlist" "foo" {
  allow_list_name = "acc-test-allowlist-${count.index}"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = ["192.168.0.0/24", "192.168.1.0/24"]
  count           = 3
}

resource "volcengine_rds_mysql_instance" "foo" {
  instance_name          = "acc-test-rds-mysql"
  db_engine_version      = "MySQL_5_7"
  node_spec              = "rds.mysql.1c2g"
  primary_zone_id        = data.volcengine_zones.foo.zones[0].id
  secondary_zone_id      = data.volcengine_zones.foo.zones[0].id
  storage_space          = 80
  subnet_id              = volcengine_subnet.foo.id
  lower_case_table_names = "1"
  charge_info {
    charge_type = "PostPaid"
  }
  parameters {
    parameter_name  = "auto_increment_increment"
    parameter_value = "2"
  }
  parameters {
    parameter_name  = "auto_increment_offset"
    parameter_value = "4"
  }

  allow_list_ids = volcengine_rds_mysql_allowlist.foo[*].id
}

data "volcengine_rds_mysql_allowlists" "foo" {
  instance_id = volcengine_rds_mysql_instance.foo.id
  region_id   = "cn-beijing"
}
```
## Argument Reference
The following arguments are supported:
* `region_id` - (Required) The region of the allow lists.
* `instance_id` - (Optional) Instance ID. When an InstanceId is specified, the DescribeAllowLists interface will return the whitelist bound to the specified instance.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `allow_lists` - The list of allowed list.
    * `allow_list_category` - White list category. Values:
 Ordinary: Ordinary white list. Default: Default white list. Description: When this parameter is used as a request parameter, the default value is Ordinary.
    * `allow_list_desc` - The description of the allow list.
    * `allow_list_id` - The id of the allow list.
    * `allow_list_ip_num` - The total number of IP addresses (or address ranges) in the whitelist.
    * `allow_list_name` - The name of the allow list.
    * `allow_list_type` - The type of the allow list.
    * `allow_list` - The IP address or a range of IP addresses in CIDR format.
    * `associated_instance_num` - The total number of instances bound under the whitelist.
    * `associated_instances` - The list of instances.
        * `instance_id` - The id of the instance.
        * `instance_name` - The name of the instance.
        * `vpc` - The id of the vpc.
    * `security_group_bind_infos` - Whitelist information for the associated security group.
        * `bind_mode` - The schema for the associated security group.
 IngressDirectionIp: Incoming Direction IP. 
 AssociateEcsIp: Associate ECSIP. 
explain: In the CreateAllowList interface, SecurityGroupBindInfoObject BindMode and SecurityGroupId fields are required.
        * `ip_list` - The ip list of the security group.
        * `security_group_id` - The security group id of the allow list.
        * `security_group_name` - The name of the security group.
    * `user_allow_list` - IP addresses outside the security group that need to be added to the whitelist. IP addresses or IP address segments in CIDR format can be entered. Note: This field cannot be used simultaneously with AllowList.
* `total_count` - The total count of Scaling Activity query.


