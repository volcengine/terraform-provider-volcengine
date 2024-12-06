---
subcategory: "VEDB_MYSQL"
layout: "volcengine"
page_title: "Volcengine: volcengine_vedb_mysql_allowlists"
sidebar_current: "docs-volcengine-datasource-vedb_mysql_allowlists"
description: |-
  Use this data source to query detailed information of vedb mysql allowlists
---
# volcengine_vedb_mysql_allowlists
Use this data source to query detailed information of vedb mysql allowlists
## Example Usage
```hcl
resource "volcengine_vedb_mysql_allowlist" "foo" {
  allow_list_name = "acc-test-allowlist"
  allow_list_desc = "acc-test"
  allow_list_type = "IPv4"
  allow_list      = ["192.168.0.0/24", "192.168.1.0/24", "192.168.2.0/24"]
}

data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[2].id
  vpc_id      = volcengine_vpc.foo.id
}


resource "volcengine_vedb_mysql_instance" "foo" {
  charge_type         = "PostPaid"
  storage_charge_type = "PostPaid"
  db_engine_version   = "MySQL_8_0"
  db_minor_version    = "3.0"
  node_number         = 2
  node_spec           = "vedb.mysql.x4.large"
  subnet_id           = volcengine_subnet.foo.id
  instance_name       = "tf-test"
  project_name        = "testA"
  tags {
    key   = "tftest"
    value = "tftest"
  }
  tags {
    key   = "tftest2"
    value = "tftest2"
  }
}

resource "volcengine_vedb_mysql_allowlist_associate" "foo" {
  allow_list_id = volcengine_vedb_mysql_allowlist.foo.id
  instance_id   = volcengine_vedb_mysql_instance.foo.id
}

data "volcengine_vedb_mysql_allowlists" "foo" {
  instance_id = volcengine_vedb_mysql_instance.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `region_id` - (Required) The region of the allow lists.
* `instance_id` - (Optional) Instance ID. When an InstanceId is specified, the DescribeAllowLists interface will return the whitelist bound to the specified instance.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `allow_lists` - The collection of query.
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
* `total_count` - The total count of query.


