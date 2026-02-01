---
subcategory: "DIRECT_CONNECT"
layout: "volcengine"
page_title: "Volcengine: volcengine_direct_connect_connection"
sidebar_current: "docs-volcengine-resource-direct_connect_connection"
description: |-
  Provides a resource to manage direct connect connection
---
# volcengine_direct_connect_connection
Provides a resource to manage direct connect connection
## Example Usage
```hcl
resource "volcengine_direct_connect_connection" "foo" {
  direct_connect_connection_name = "tf-test-connection"
  description                    = "tf-test"
  direct_connect_access_point_id = "ap-cn-beijing-a"
  line_operator                  = "ChinaOther"
  port_type                      = "10GBase"
  port_spec                      = "10G"
  bandwidth                      = 1000
  peer_location                  = "XX路XX号XX楼XX机房"
  customer_name                  = "tf-a"
  customer_contact_phone         = "12345678911"
  customer_contact_email         = "email@aaa.com"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `bandwidth` - (Required, ForceNew) The line band width,unit:Mbps.
* `customer_contact_email` - (Required) The dedicated line contact email.
* `customer_contact_phone` - (Required) The dedicated line contact phone.
* `customer_name` - (Required) The dedicated line contact name.
* `direct_connect_access_point_id` - (Required, ForceNew) The direct connect access point id.
* `line_operator` - (Required, ForceNew) The physical leased line operator.valid value contains `ChinaTelecom`,`ChinaMobile`,`ChinaUnicom`,`ChinaOther`.
* `peer_location` - (Required, ForceNew) The local IDC address.
* `port_spec` - (Required, ForceNew) The physical leased line port spec.valid value contains `1G`,`10G`.
* `port_type` - (Required, ForceNew) The physical leased line port type and spec.valid value contains `1000Base-T`,`10GBase-T`,`1000Base`,`10GBase`,`40GBase`,`100GBase`.
* `description` - (Optional) The description of direct connect.
* `direct_connect_connection_name` - (Optional) The name of direct connect.
* `tags` - (Optional) The physical leased line tags.

The `tags` object supports the following:

* `key` - (Optional) The tag key.
* `value` - (Optional) The tag value.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
DirectConnectConnection can be imported using the id, e.g.
```
$ terraform import volcengine_direct_connect_connection.default dcc-7qthudw0ll6jmc****
```

