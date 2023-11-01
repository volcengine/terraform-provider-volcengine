---
subcategory: "DIRECT_CONNECT"
layout: "volcengine"
page_title: "Volcengine: volcengine_direct_connect_connections"
sidebar_current: "docs-volcengine-datasource-direct_connect_connections"
description: |-
  Use this data source to query detailed information of direct connect connections
---
# volcengine_direct_connect_connections
Use this data source to query detailed information of direct connect connections
## Example Usage
```hcl
data "volcengine_direct_connect_connections" "foo" {
  direct_connect_connection_name = "tf_test"
}
```
## Argument Reference
The following arguments are supported:
* `connection_type` - (Optional) The connection type of physical leased line,valid value contains `SharedConnection`,`DedicatedConnection`.
* `direct_connect_access_point_id` - (Optional) The ID of the physical leased line access point.
* `direct_connect_connection_name` - (Optional) The name of directi connect connection.
* `ids` - (Optional) A list of IDs.
* `line_operator` - (Optional) The operator of the physical leased line,valid value contains `ChinaTelecom`,`ChinaMobile`,`ChinaUnicom`,`ChinaOther`.
* `name_regex` - (Optional) A Name Regex of Resource.
* `output_file` - (Optional) File name where to save data source results.
* `peer_location` - (Optional) The peer access point of the physical leased line.
* `tag_filters` - (Optional) The filter tag of direct connect.

The `tag_filters` object supports the following:

* `key` - (Optional) The tag key of cloud resource instance.
* `value` - (Optional) The tag value of cloud resource instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `direct_connect_connections` - The collection of query.
    * `account_id` - The account ID which the physical leased line belongs.
    * `bandwidth` - The bandwidth of direct connect.
    * `billing_type` - The dedicated line billing type,only support `1` for yearly and monthly billing currently.
    * `business_status` - The dedicated line billing status.
    * `connection_type` - The connection type of direct connect.
    * `creation_time` - The creation time of direct connect.
    * `customer_contact_email` - The dedicated line contact email.
    * `customer_contact_phone` - The dedicated line contact phone.
    * `customer_name` - The dedicated line contact name.
    * `deleted_time` - The expected resource force collection time.
    * `description` - The description of direct connect connection.
    * `direct_connect_access_point_id` - The access point id of direct connect.
    * `direct_connect_connection_id` - The ID of direct connect connection.
    * `direct_connect_connection_name` - The name of direct connect connection.
    * `expect_bandwidth` - The expect bandwidth of direct connect.
    * `expired_time` - The expired time.
    * `line_operator` - The operator of physical leased line.
    * `parent_connection_account_id` - The account ID of physical leased line to which the shared leased line belongs.If the physical leased line type is an exclusive leased line,this parameter returns empty.
    * `parent_connection_id` - The ID of the physical leased line to which the shared leased line belongs. If the physical leased line type is an exclusive leased line, this parameter returns empty.
    * `peer_location` - The peer access point of the physical leased line.
    * `port_spec` - The dedicated line port spec.
    * `port_type` - The port type of direct connect.
    * `status` - The status of physical leased line.
    * `tags` - All tags that physical leased line added.
        * `key` - The tag key.
        * `value` - The tag value.
    * `update_time` - The update time of direct connect.
    * `vlan_id` - The vlan ID of shared connection,if `connection_type` is `DedicatedConnection`,this parameter returns 0.
* `total_count` - The total count of query.


