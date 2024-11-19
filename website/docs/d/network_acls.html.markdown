---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_network_acls"
sidebar_current: "docs-volcengine-datasource-network_acls"
description: |-
  Use this data source to query detailed information of network acls
---
# volcengine_network_acls
Use this data source to query detailed information of network acls
## Example Usage
```hcl
data "volcengine_network_acls" "default" {
  #  ids = ["nacl-172leak37mi9s4d1w33pswqkh"]
  #  vpc_id = "vpc-ru0wv9alfoxsu3nuld85rpp"
  #  subnet_id = "subnet-637jxq81u5mon3gd6ivc7rj"
  network_acl_name = "ms-tf-acl"
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of Network Acl IDs.
* `name_regex` - (Optional) A Name Regex of Network Acl.
* `network_acl_name` - (Optional) The name of Network Acl.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of the network acl.
* `subnet_id` - (Optional) The subnet id of Network Acl.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional) The vpc id of Network Acl.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `network_acls` - The collection of Network Acl query.
    * `acl_entry_count` - The count of Network acl entry.
    * `creation_time` - Creation time of Network Acl.
    * `description` - The description of Network Acl.
    * `egress_acl_entries` - The egress entries info of Network Acl.
        * `description` - The description of entry.
        * `destination_cidr_ip` - The DestinationCidrIp of entry.
        * `network_acl_entry_id` - The id of entry.
        * `network_acl_entry_name` - The name of entry.
        * `policy` - The policy of entry.
        * `port` - The port of entry.
        * `priority` - The priority of entry.
        * `protocol` - The protocol of entry.
    * `id` - The ID of Network Acl.
    * `ingress_acl_entries` - The ingress entries info of Network Acl.
        * `description` - The description of entry.
        * `network_acl_entry_id` - The id of entry.
        * `network_acl_entry_name` - The name of entry.
        * `policy` - The policy of entry.
        * `port` - The port of entry.
        * `priority` - The priority of entry.
        * `protocol` - The protocol of entry.
        * `source_cidr_ip` - The SourceCidrIp of entry.
    * `network_acl_id` - The ID of Network Acl.
    * `network_acl_name` - The Name of Network Acl.
    * `project_name` - The project name of the network acl.
    * `resources` - The resources info of Network Acl.
        * `resource_id` - The resource id of Network Acl.
        * `status` - The resource status of Network Acl.
    * `status` - The Status of Network Acl.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `update_time` - Update time of Network Acl.
    * `vpc_id` - The vpc id of Network Acl.
* `total_count` - The total count of Network Acl query.


